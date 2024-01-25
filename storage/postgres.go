package storage

import (
	"context"
	"database/sql"
	"departement/db"
	"departement/models"
	"log"

	_ "github.com/lib/pq"
)

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

		err := rows.Scan(&ranking.ID, &ranking.UserID, &ranking.Score)
		if err != nil {
			return nil, err
		}

		rankings = append(rankings, ranking)
	}

	return rankings, nil
}

// CreateRanking creates a new ranking in the database
func (ps *PostgresStorage) CreateRanking(ranking models.Ranking) (models.Ranking, error) {
	query := "INSERT INTO rankings (id, user_id, score) VALUES ($1, $2, $3) RETURNING id"
	row := ps.DB.QueryRow(query, ranking.ID, ranking.UserID, ranking.Score)

	err := row.Scan(&ranking.ID)
	return ranking, err
}

// GetRankingByID retrieves a ranking from the database by ID
func (ps *PostgresStorage) GetRankingByUserID(userID int) (models.Ranking, error) {
	query := "SELECT id, id, user_id, score FROM rankings WHERE id = $1"
	row := ps.DB.QueryRow(query, userID)

	var ranking models.Ranking
	err := row.Scan(&ranking.ID, &ranking.UserID, &ranking.Score)
	return ranking, err
}
