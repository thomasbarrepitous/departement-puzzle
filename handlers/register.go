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

type RegisterHandler struct {
	DB *sql.DB
}

// createUserInDB creates a new user in the database
func (rh *RegisterHandler) createUserInDB(user models.User) (models.User, error) {
	_, err := rh.DB.Exec("INSERT INTO users (username, password, email) VALUES ($1, $2, $3)", user.Username, user.Password, user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

// CreateUser handles the request to create a new user
func (rh *RegisterHandler) CreateUser(user models.User) (models.User, error) {
	// Hash the password
	user.SetPassword(user.Password)

	// Create the user in the database
	user, dbErr := rh.createUserInDB(user)
	if dbErr != nil {
		log.Print(dbErr.Error())
		return user, dbErr
	}

	return user, nil
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
	user, creationError := rh.CreateUser(user)
	if creationError != nil {
		utils.JSONRespond(w, http.StatusInternalServerError, creationError)
		return
	}

	userResponse := models.UserResponse{
		Username: user.Username,
		Email:    user.Email,
	}

	// Redirect to the login page if the user was created successfully
	w.Header().Add("HX-Redirect", "/login")
	utils.JSONRespond(w, http.StatusOK, userResponse)
}

// RenderRegisterPage renders the login page
func (rh *RegisterHandler) RenderRegisterPage(w http.ResponseWriter, r *http.Request) {
	component := components.RegisterPageComponent()
	component.Render(r.Context(), w)
}
