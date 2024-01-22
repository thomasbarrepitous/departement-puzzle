package handlers

import (
	"departement/storage"
	"departement/utils"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// UserHandler represents the controller for user-related operations
type UserHandler struct {
	Store storage.UserStorage
}

// GetAllUsers handles the request to get all users
func (uh *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uh.Store.GetAllUsers()
	if err != nil {
		utils.JSONRespond(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSONRespond(w, http.StatusOK, users)
}

// GetUserByID handles the request to get a user by its ID
func (uh *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the URL
	vars := mux.Vars(r)
	userID := vars["id"]

	// Get the user from the database
	user, err := uh.Store.GetUserByUsername(userID)
	if err != nil {
		utils.JSONRespond(w, http.StatusInternalServerError, err)
		return
	}

	// Respond with the user
	utils.JSONRespond(w, http.StatusOK, user)
}
