package handlers

import (
	"database/sql"
	"departement/components"
	"departement/models"
	"net/http"
)

// LoginHandler is the handler for the login page
type LoginHandler struct {
	DB *sql.DB
}

// getUserByUsername retrieves a user from the database by username
func (lh *LoginHandler) getUserByUsername(username string) (models.User, error) {
	query := "SELECT id, username, password FROM users WHERE username = $1"
	row := lh.DB.QueryRow(query, username)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	return user, err
}

// userIsAuthenticated checks if the user is authenticated
func userIsAuthenticated(r *http.Request) bool {
	return true // Dummy example: Always assume the user is authenticated
}

// Handle the classic login submission and redirect to the dashboard if successful
func (lh *LoginHandler) ClassicHandle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	user, wrongUsernameErr := lh.getUserByUsername(username)
	if wrongUsernameErr != nil {
		lh.RenderLoginPage(w, r)
		return
	}

	wrongPasswordErr := user.CheckPassword(password)
	if wrongPasswordErr != nil {
		lh.RenderLoginPage(w, r)
		return
	}

	// lh.RenderProfilePage(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// RenderLoginPage renders the login page
func (lh *LoginHandler) RenderLoginPage(w http.ResponseWriter, r *http.Request) {
	component := components.LoginPageComponent()
	component.Render(r.Context(), w)
}