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

func (lh *LoginHandler) RenderLoginPage(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		Username: "test",
		Password: "test",
		Email:    "oui@oui.com",
	}

	component := components.LoginPageComponent(user)
	component.Render(r.Context(), w)
}
