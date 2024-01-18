package handlers

import (
	"database/sql"
	"departement/components"
	"net/http"
)

// LoginHandler is the handler for the login page
type LoginHandler struct {
	DB *sql.DB
}

func (lh *LoginHandler) RenderLoginPage(w http.ResponseWriter, r *http.Request) {
	component := components.LoginPageComponent()
	component.Render(r.Context(), w)
}
