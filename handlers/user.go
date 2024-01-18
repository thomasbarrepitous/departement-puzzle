package handlers

import (
	"context"
	"database/sql"
	"departement/models"
	"departement/utils"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// UserHandler represents the controller for user-related operations
type UserHandler struct {
	DB  *sql.DB
	Ctx context.Context
}

// GetAllUsers handles the request to get all users
func (uh *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := uh.DB.Query("SELECT * FROM users")
	if err != nil {
		utils.JSONRespond(w, http.StatusInternalServerError, err)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
		if err != nil {
			utils.JSONRespond(w, http.StatusInternalServerError, err)
			return
		}

		users = append(users, user)
	}

	utils.JSONRespond(w, http.StatusOK, users)
}

// getUserByIDInDB queries the database to get a user by its ID
func (uh *UserHandler) getUserByIDInDB(userID string) (models.User, error) {
	row := uh.DB.QueryRow("SELECT * FROM users WHERE id = $1", userID)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return user, err
	}

	return user, nil
}

// GetUserByID handles the request to get a user by its ID
func (uh *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the URL
	vars := mux.Vars(r)
	userID := vars["id"]

	// Query the database
	user, err := uh.getUserByIDInDB(userID)
	if err != nil {
		utils.JSONRespond(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSONRespond(w, http.StatusOK, user)
}

// createUserInDB creates a new user in the database
func (uh *UserHandler) createUserInDB(user models.User) (models.User, error) {
	_, err := uh.DB.Exec("INSERT INTO users (username, password, email) VALUES ($1, $2, $3)", user.Username, user.Password, user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

// CreateUser handles the request to create a new user
func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		utils.JSONRespond(w, http.StatusBadRequest, readErr)
		return
	}

	// Parse the request body to a user struct
	var user models.User
	parseErr := json.Unmarshal(body, &user)
	if parseErr != nil {
		utils.JSONRespond(w, http.StatusBadRequest, parseErr)
		return
	}

	// Hash the password
	user.SetPassword(user.Password)

	// Create the user in the database
	user, dbErr := uh.createUserInDB(user)
	if dbErr != nil {
		utils.JSONRespond(w, http.StatusInternalServerError, dbErr)
		return
	}

	utils.JSONRespond(w, http.StatusOK, user)
}
