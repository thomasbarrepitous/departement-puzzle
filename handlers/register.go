package handlers

import (
	"departement/components"
	"departement/models"
	"departement/storage"
	"departement/utils"
	"errors"
	"log"
	"net/http"
)

type RegisterHandler struct {
	Store storage.UserStorage
}

func (rh *RegisterHandler) RegisterHandle(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Decode the request body into the user struct
	decodeErr := utils.DecodeJSONBody(w, r, &user)
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

	// Create the user in the database
	// We reinitialize the user struct to hash the password
	// As the New() function is not called when unmarshalling the JSON
	user, dbErr := rh.Store.CreateUser(models.NewUser(user.Username, user.Password, user.Email))
	if dbErr != nil {
		log.Print(dbErr.Error())
		utils.JSONRespond(w, http.StatusInternalServerError, dbErr)
		return
	}

	userResponse := models.UserResponse{
		Username: user.Username,
		Email:    user.Email,
	}

	w.Header().Add("HX-Redirect", "/login")
	utils.JSONRespond(w, http.StatusOK, userResponse)
}

// RenderRegisterPage renders the login page
func (rh *RegisterHandler) RenderRegisterPage(w http.ResponseWriter, r *http.Request) {
	component := components.RegisterPageComponent()
	component.Render(r.Context(), w)
}
