package handlers

import (
	"database/sql"
	"departement/components"
	"departement/models"
	"departement/storage"
	"departement/utils"
	"errors"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

// LoginHandler is the handler for the login page
type LoginHandler struct {
	Store storage.UserStorage
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RenderLoginPage renders the login page
func (lh *LoginHandler) RenderLoginPage(w http.ResponseWriter, r *http.Request) {
	component := components.LoginPageComponent()
	component.Render(r.Context(), w)
}

// Handle the login submission
func (lh *LoginHandler) JWTLoginHandle(w http.ResponseWriter, r *http.Request) {
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

	user, wrongUsernameErr := lh.Store.GetUserByUsername(loginRequest.Username)
	if wrongUsernameErr != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	wrongPasswordErr := user.CheckPassword(loginRequest.Password)
	if wrongPasswordErr != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	lh.setAuthCookieAndRedirect(w, r, user, "/")
	utils.JSONRespond(w, http.StatusOK, map[string]string{})
}

// Handle the Google login submission, redirect to Google's OAuth 2.0 server
func (lh *LoginHandler) GoogleLoginHandle(w http.ResponseWriter, r *http.Request) {
	// Create oauth2 config
	config := utils.CreateGoogleOAuth2Config()

	// Generate the URL to request an authorization code
	url := config.AuthCodeURL("state", oauth2.AccessTypeOffline)

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

	// Get the user's profile with the access token
	profile, err := utils.GetGoogleProfile(r.Context(), config, accessToken)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Check if the user is already in the database
	user, err := lh.Store.GetUserByEmail(profile.Email)
	if err != nil {
		// If the user is not in the database, create it
		if err == sql.ErrNoRows {
			// Create the user
			log.Print("User not found, creating it")
			user, err = lh.Store.CreateUser(models.User{
				Username: profile.Name,
				Email:    profile.Email,
				// TODO: Generate a random password
				Password: "",
			})
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

	// Set the JWT token in the HTTPOnly cookie and redirect to the home page
	lh.setAuthCookieAndRedirect(w, r, user, "")
	http.Redirect(w, r, "/", http.StatusFound)
	utils.JSONRespond(w, http.StatusOK, map[string]string{})
}

// Handle the logout : delete the cookie and redirect to the login page
func (lh *LoginHandler) LogoutHandle(w http.ResponseWriter, r *http.Request) {
	// Delete the cookie
	http.SetCookie(w, &http.Cookie{
		Name:   "Authorization",
		MaxAge: -1,
	})

	w.Header().Add("HX-Redirect", "/login")
	utils.JSONRespond(w, http.StatusOK, map[string]string{})
}

// Set the JWT token in the HTTPOnly cookie and redirect to the callback URL
func (lh *LoginHandler) setAuthCookieAndRedirect(w http.ResponseWriter, r *http.Request, user models.User, redirectURL string) {
	// Create a JWT token
	token, tokenErr := utils.CreateJWT(user.ID)
	if tokenErr != nil {
		log.Print(tokenErr.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Set the token in the HTTPOnly cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "Authorization",
		Value:    token,
		MaxAge:   3600 * 24 * 7,
		SameSite: http.SameSiteLaxMode,
		// Don't set Secure to true in development
		Secure:   false,
		HttpOnly: true,
	})

	if redirectURL != "" {
		w.Header().Add("HX-Redirect", redirectURL)
	}
}
