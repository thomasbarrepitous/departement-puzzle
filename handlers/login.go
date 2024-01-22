package handlers

import (
	"database/sql"
	"departement/components"
	"departement/models"
	"departement/utils"
	"errors"
	"log"
	"net/http"
)

// LoginHandler is the handler for the login page
type LoginHandler struct {
	DB *sql.DB
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// getUserByUsername retrieves a user from the database by username
func (lh *LoginHandler) getUserByUsername(username string) (models.User, error) {
	query := "SELECT id, username, password FROM users WHERE username = $1"
	row := lh.DB.QueryRow(query, username)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	return user, err
}

// Handle the classic login submission
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

	user, wrongUsernameErr := lh.getUserByUsername(loginRequest.Username)
	if wrongUsernameErr != nil {
		log.Print(wrongUsernameErr.Error())
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	wrongPasswordErr := user.CheckPassword(loginRequest.Password)
	if wrongPasswordErr != nil {
		log.Print(wrongPasswordErr.Error())
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	// Create a JWT token
	token, tokenErr := utils.CreateToken(user.ID)
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

	w.Header().Add("HX-Redirect", "/")
	utils.JSONRespond(w, http.StatusOK, map[string]string{})
}

// RenderLoginPage renders the login page
func (lh *LoginHandler) RenderLoginPage(w http.ResponseWriter, r *http.Request) {
	component := components.LoginPageComponent()
	component.Render(r.Context(), w)
}
