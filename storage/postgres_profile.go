package storage

import (
	"context"
	"database/sql"
	"departement/db"
	"departement/models"
	"log"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

// TODO : Split this into multiple files
type PostgresProfileStorage struct {
	// DB is the database connection
	DB *sql.DB
}

// TODO : Transform this into a singleton ?
func NewPostgresProfileStorage() *PostgresProfileStorage {
	// Initialize the postgres database
	postgresDB, err := db.ConnectDB(*db.NewPostgresConfig())
	if err != nil {
		log.Fatal(err)
	}

	return &PostgresProfileStorage{
		DB: postgresDB,
	}
}

// GetAllProfiles retrieves all profiles from the database
func (pfs *PostgresProfileStorage) GetAllProfiles(ctx context.Context) ([]models.Profile, error) {
	rows, err := pfs.DB.Query("SELECT * FROM profiles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []models.Profile
	for rows.Next() {
		var profile models.Profile

		err := rows.Scan(&profile.ID, &profile.UserID, &profile.Username, &profile.Email, &profile.Picture, &profile.Description, &profile.Country)
		if err != nil {
			return nil, err
		}

		profiles = append(profiles, profile)
	}

	return profiles, nil
}

// CreateProfile creates a new profile in the database
func (pfs *PostgresProfileStorage) CreateProfile(ctx context.Context, profile models.Profile) (models.Profile, error) {
	query := "INSERT INTO profiles (user_id, username, email, picture, description, country) VALUES ($1, $2, $3, $4, $5, $6) RETURNING user_id"
	row := pfs.DB.QueryRow(query, profile.UserID, profile.Username, profile.Email, profile.Picture, profile.Description, profile.Country)

	err := row.Scan(&profile.UserID)
	return profile, err
}

// GetProfileByID retrieves a profile from the database by ID
func (pfs *PostgresProfileStorage) GetProfileByUserID(ctx context.Context, userID uuid.UUID) (models.Profile, error) {
	query := "SELECT id, user_id, username, email, picture, description, country FROM profiles WHERE user_id = $1"
	row := pfs.DB.QueryRow(query, userID)

	var profile models.Profile
	err := row.Scan(&profile.ID, &profile.UserID, &profile.Username, &profile.Email, &profile.Picture, &profile.Description, &profile.Country)
	return profile, err
}

// UpdateProfile updates a profile in the database
func (pfs *PostgresProfileStorage) UpdateProfile(ctx context.Context, id uuid.UUID, profile models.Profile) (models.Profile, error) {
	query := "UPDATE profiles SET username = $1, email = $2, picture = $3, description = $4, country = $5 WHERE id = $6 RETURNING id"
	row := pfs.DB.QueryRow(query, profile.Username, profile.Email, profile.Picture, profile.Description, profile.ID)

	err := row.Scan(&profile.ID)
	return profile, err
}

// DeleteProfile deletes a profile from the database
func (pfs *PostgresProfileStorage) DeleteProfile(ctx context.Context, id uuid.UUID) error {
	query := "DELETE FROM profiles WHERE id = $1"
	_, err := pfs.DB.Exec(query, id)

	return err
}
