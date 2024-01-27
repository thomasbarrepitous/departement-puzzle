package storage

import (
	"context"
	"database/sql"
	"departement/db"
	"departement/models"
	"log"

	_ "github.com/lib/pq"
)

// TODO : Split this into multiple files
type PostgresStorage struct {
	// DB is the database connection
	DB *sql.DB
}

// TODO : Transform this into a singleton
func NewPostgresStorage() *PostgresStorage {
	// Initialize the postgres database
	postgresDB, err := db.ConnectDB(*db.NewGameConfig())
	if err != nil {
		log.Fatal(err)
	}

	return &PostgresStorage{
		DB: postgresDB,
	}
}

// GetAllUsers retrieves all users from the database
func (ps *PostgresStorage) GetAllUsers(ctx context.Context) ([]models.User, error) {
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
func (ps *PostgresStorage) GetUserByID(ctx context.Context, id int) (models.User, error) {
	query := "SELECT id, username, email, password FROM users WHERE id = $1"
	row := ps.DB.QueryRow(query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	return user, err
}

// GetUserByEmail retrieves a user from the database by email
func (ps *PostgresStorage) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	query := "SELECT id, username, email, password FROM users WHERE email = $1"
	row := ps.DB.QueryRow(query, email)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	return user, err
}

// GetUserByUsername retrieves a user from the database by username
func (ps *PostgresStorage) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	query := "SELECT id, username, email, password FROM users WHERE username = $1"
	row := ps.DB.QueryRow(query, username)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	return user, err
}

// CreateUser creates a new user in the database
func (ps *PostgresStorage) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	query := "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id"
	row := ps.DB.QueryRow(query, user.Username, user.Email, user.Password)

	err := row.Scan(&user.ID)
	return user, err
}

// UpdateUser updates a user in the database
func (ps *PostgresStorage) UpdateUser(ctx context.Context, id int, user models.User) (models.User, error) {
	query := "UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4 RETURNING id"
	row := ps.DB.QueryRow(query, user.Username, user.Email, user.Password, id)

	err := row.Scan(&user.ID)
	return user, err
}

// DeleteUser deletes a user from the database
func (ps *PostgresStorage) DeleteUser(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := ps.DB.Exec(query, id)

	return err
}

// Profile

// GetAllProfiles retrieves all profiles from the database
func (ps *PostgresStorage) GetAllProfiles() ([]models.Profile, error) {
	rows, err := ps.DB.Query("SELECT * FROM profiles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []models.Profile
	for rows.Next() {
		var profile models.Profile

		err := rows.Scan(&profile.ID, &profile.UserID, &profile.Username, &profile.Email, &profile.Picture, &profile.Description)
		if err != nil {
			return nil, err
		}

		profiles = append(profiles, profile)
	}

	return profiles, nil
}

// CreateProfile creates a new profile in the database
func (ps *PostgresStorage) CreateProfile(profile models.Profile) (models.Profile, error) {
	query := "INSERT INTO profiles (id, user_id, username, email, picture, description) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	row := ps.DB.QueryRow(query, profile.ID, profile.UserID, profile.Username, profile.Email, profile.Picture, profile.Description)

	err := row.Scan(&profile.ID)
	return profile, err
}

// GetProfileByID retrieves a profile from the database by ID
func (ps *PostgresStorage) GetProfileByUserID(userID int) (models.Profile, error) {
	query := "SELECT id, user_id, username, email, picture, description FROM profiles WHERE user_id = $1"
	row := ps.DB.QueryRow(query, userID)

	var profile models.Profile
	err := row.Scan(&profile.ID, &profile.UserID, &profile.Username, &profile.Email, &profile.Picture, &profile.Description)
	return profile, err
}

// UpdateProfile updates a profile in the database
func (ps *PostgresStorage) UpdateProfile(id int, profile models.Profile) (models.Profile, error) {
	query := "UPDATE profiles SET username = $1, email = $2, picture = $3, description = $4 WHERE id = $5 RETURNING id"
	row := ps.DB.QueryRow(query, profile.Username, profile.Email, profile.Picture, profile.Description, profile.ID)

	err := row.Scan(&profile.ID)
	return profile, err
}

// DeleteProfile deletes a profile from the database
func (ps *PostgresStorage) DeleteProfile(id int) error {
	query := "DELETE FROM profiles WHERE id = $1"
	_, err := ps.DB.Exec(query, id)

	return err
}

// Rankings

// GetAllRankings retrieves all rankings from the database
func (ps *PostgresStorage) GetAllRankings() ([]models.Ranking, error) {
	rows, err := ps.DB.Query("SELECT * FROM rankings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rankings []models.Ranking
	for rows.Next() {
		var ranking models.Ranking

		err := rows.Scan(&ranking.ID, &ranking.UserID, &ranking.PointsScore, &ranking.TimeScore)
		if err != nil {
			return nil, err
		}

		rankings = append(rankings, ranking)
	}

	return rankings, nil
}

// GetAllRankingsByUserID retrieves all rankings from the database by user ID
func (ps *PostgresStorage) GetAllRankingsByUserID(userID int) ([]models.Ranking, error) {
	rows, err := ps.DB.Query("SELECT * FROM rankings WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rankings []models.Ranking
	for rows.Next() {
		var ranking models.Ranking

		err := rows.Scan(&ranking.ID, &ranking.UserID, &ranking.PointsScore, &ranking.TimeScore)
		if err != nil {
			return nil, err
		}

		rankings = append(rankings, ranking)
	}

	return rankings, nil
}

// CreateRanking creates a new ranking in the database
func (ps *PostgresStorage) CreateRanking(ranking models.Ranking) (models.Ranking, error) {
	query := "INSERT INTO rankings (id, user_id, points_score, time_score) VALUES ($1, $2, $3, $4) RETURNING id"
	row := ps.DB.QueryRow(query, ranking.ID, ranking.UserID, ranking.PointsScore, ranking.TimeScore)

	err := row.Scan(&ranking.ID)
	return ranking, err
}

// GetRankingByID retrieves a ranking from the database by ID
func (ps *PostgresStorage) GetRankingByUserID(userID int) (models.Ranking, error) {
	query := "SELECT id, id, user_id, points_score, time_score FROM rankings WHERE id = $1"
	row := ps.DB.QueryRow(query, userID)

	var ranking models.Ranking
	err := row.Scan(&ranking.ID, &ranking.UserID, &ranking.PointsScore, &ranking.TimeScore)
	return ranking, err
}
