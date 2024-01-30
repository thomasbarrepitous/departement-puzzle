package storage

import (
	"context"
	"database/sql"
	"departement/db"
	"departement/models"
	"log"

	_ "github.com/lib/pq"
)

type PostgresUserStorage struct {
	DB *sql.DB
}

// TODO : Transform this into a singleton ?
func NewPostgresUserStorage() *PostgresUserStorage {
	// Initialize the postgres database
	postgresDB, err := db.ConnectDB(*db.NewPostgresConfig())
	if err != nil {
		log.Fatal(err)
	}

	return &PostgresUserStorage{
		DB: postgresDB,
	}
}

// GetAllUsers retrieves all users from the database
func (ps *PostgresUserStorage) GetAllUsers(ctx context.Context) ([]models.User, error) {
	rows, err := ps.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// GetUserByID retrieves a user from the database by ID
func (ps *PostgresUserStorage) GetUserByID(ctx context.Context, id int) (models.User, error) {
	query := "SELECT id, username, email, password FROM users WHERE id = $1"
	row := ps.DB.QueryRow(query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	return user, err
}

// GetUserByEmail retrieves a user from the database by email
func (ps *PostgresUserStorage) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	query := "SELECT id, username, email, password FROM users WHERE email = $1"
	row := ps.DB.QueryRow(query, email)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	return user, err
}

// GetUserByUsername retrieves a user from the database by username
func (ps *PostgresUserStorage) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	query := "SELECT id, username, email, password FROM users WHERE username = $1"
	row := ps.DB.QueryRow(query, username)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	return user, err
}

// CreateUser creates a new user in the database
func (ps *PostgresUserStorage) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id"
	row := ps.DB.QueryRow(query, user.Username, user.Email, user.Password)

	err := row.Scan(&user.ID)
	return user, err
}

// UpdateUser updates a user in the database
func (ps *PostgresUserStorage) UpdateUser(ctx context.Context, id int, user models.User) (models.User, error) {
	query := "UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4 RETURNING id"
	row := ps.DB.QueryRow(query, user.Username, user.Email, user.Password, id)

	err := row.Scan(&user.ID)
	return user, err
}

// DeleteUser deletes a user from the database
func (ps *PostgresUserStorage) DeleteUser(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := ps.DB.Exec(query, id)

	return err
}
