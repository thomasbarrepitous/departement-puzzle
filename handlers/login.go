package handlers

import (
	"context"
	"database/sql"
	"departement/components"
	"departement/models"
	"departement/storage"
	"departement/utils"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

// LoginHandler is the handler for the login page
type LoginHandler struct {
	UserStore    storage.UserStorage
	ProfileStore storage.ProfileStorage
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RenderLoginPage renders the login page
func (lh *LoginHandler) RenderLoginPage(w http.ResponseWriter, r *http.Request) {
	component := components.LoginPageComponent(r)
	component.Render(r.Context(), w)
}

// Handle the login submission
func (lh *LoginHandler) EmailLoginHandle(w http.ResponseWriter, r *http.Request) {
	loginRequest := LoginRequest{}
	// Decode the request body into the user struct
	decodeErr := utils.DecodeJSONBody(w, r, &loginRequest)
	if decodeErr != nil {
		var mr *utils.MalformedRequest
		if errors.As(decodeErr, &mr) {
			http.Error(w, mr.Message, mr.Status)
		} else {
			log.Print(decodeErr.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	// Check if the user is in the database
	user, wrongEmailPassword := lh.UserStore.GetUserByEmail(r.Context(), loginRequest.Email)
	if wrongEmailPassword != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	wrongPasswordErr := user.CheckPassword(loginRequest.Password)
	if wrongPasswordErr != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	lh.setAuthCookie(w, r, user)
	w.Header().Add("HX-Redirect", fmt.Sprintf("/profile/%s", user.Username))
	utils.JSONRespond(w, http.StatusOK, map[string]string{})
}

// Handle the Google login submission, redirect to Google's OAuth 2.0 server
func (lh *LoginHandler) GoogleLoginHandle(w http.ResponseWriter, r *http.Request) {
	// Create oauth2 config
	config := utils.CreateGoogleOAuth2Config()

	// Generate UUID for OAuth2 state
	state := "state"

	// Generate the URL to request an authorization code
	url := config.AuthCodeURL(state, oauth2.AccessTypeOffline)

	// Redirect to Google's OAuth 2.0 server
	w.Header().Add("HX-Redirect", url)
	utils.JSONRespond(w, http.StatusTemporaryRedirect, map[string]string{})
}

// Handle Google OAuth2 callback (redirected from Google's OAuth 2.0 server)
func (lh *LoginHandler) GoogleCallbackHandle(w http.ResponseWriter, r *http.Request) {
	// Create oauth2 config
	config := utils.CreateGoogleOAuth2Config()

	// Get the authorization authCode from the URL query
	authCode := r.FormValue("code")

	// Exchange the authorization code for an access accessToken
	accessToken, err := config.Exchange(r.Context(), authCode)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Get the user's googleProfile with the access token
	googleProfile, err := utils.GetGoogleProfile(r.Context(), config, accessToken)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Check if the user is already in the database
	user, err := lh.UserStore.GetUserByEmail(r.Context(), googleProfile.Email)
	if err != nil {
		// If the user is not in the database, create it
		if err == sql.ErrNoRows {
			// Create the user
			log.Print("User not found, creating it")
			user, err = lh.GoogleCreateUser(r.Context(), googleProfile)
			log.Print(user)
			if err != nil {
				log.Print(err.Error())

				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			if err != nil {
				log.Print(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		} else {
			log.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	// Set the JWT token in the HTTPOnly cookie and redirect to the user's profile page
	lh.setAuthCookie(w, r, user)
	http.Redirect(w, r, fmt.Sprintf("/profile/%s", user.Username), http.StatusFound)
	utils.JSONRespond(w, http.StatusOK, map[string]string{})
}

// Handle the logout : delete the cookie and redirect to the login page
func (lh *LoginHandler) LogoutHandle(w http.ResponseWriter, r *http.Request) {
	// Delete the cookie
	http.SetCookie(w,
		&http.Cookie{
			Name:     "Authorization",
			Value:    "",
			MaxAge:   -1,
			SameSite: http.SameSiteLaxMode,
			Path:     "/",
			Expires:  time.Now().Add(-1 * time.Hour),
			// Don't set Secure to true in development
			Secure:   false,
			HttpOnly: true,
		},
	)

	// Redirect to the login page
	http.Redirect(w, r, "/login", http.StatusFound)
	utils.JSONRespond(w, http.StatusOK, map[string]string{})
}

// Set the JWT token in the HTTPOnly cookie and redirect to the callback URL
func (lh *LoginHandler) setAuthCookie(w http.ResponseWriter, r *http.Request, user models.User) {
	// Create a JWT token
	token, tokenErr := utils.CreateJWT(user.ID)
	if tokenErr != nil {
		log.Print(tokenErr.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Set the JWT token in the HTTPOnly cookie
	http.SetCookie(w,
		&http.Cookie{
			Name:     "Authorization",
			Value:    token,
			MaxAge:   3600 * 24 * 7,
			SameSite: http.SameSiteLaxMode,
			Path:     "/",
			// Don't set Secure to true in development
			Secure:   false,
			HttpOnly: true,
		},
	)
}

// Handle user creation after Google login
// This feels so wrong
func (lh *LoginHandler) GoogleCreateUser(ctx context.Context, profile *utils.GoogleProfile) (models.User, error) {
	// Create the user in the database
	user, err := lh.UserStore.CreateUser(ctx,
		models.NewUser(
			profile.Name,
			// TODO: Generate a random password
			"",
			profile.Email,
		),
	)
	if err != nil {
		return user, err
	}

	// Create the profile in the database
	_, err = lh.ProfileStore.CreateProfile(ctx,
		models.NewProfile(
			user.ID,
			profile.Name,
			profile.Email,
			profile.Picture,
			"",
			"",
		),
	)
	if err != nil {
		// If the  profile creation failed, delete the user
		if err := lh.UserStore.DeleteUser(ctx, user.ID); err != nil {
			log.Print(err.Error())
			return user, err
		}
		return user, err
	}

	return user, nil
}
