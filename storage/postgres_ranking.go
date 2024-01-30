package storage

import (
	"database/sql"
	"departement/db"
	"departement/models"
	"log"

	_ "github.com/lib/pq"
)

// TODO : Split this into multiple files
type PostgresRankingStorage struct {
	// DB is the database connection
	DB *sql.DB
}

// TODO : Transform this into a singleton ?
func NewPostgresRankingStorage() *PostgresRankingStorage {
	// Initialize the postgres database
	postgresDB, err := db.ConnectDB(*db.NewPostgresConfig())
	if err != nil {
		log.Fatal(err)
	}

	return &PostgresRankingStorage{
		DB: postgresDB,
	}
}

// Rankings

// GetAllRankings retrieves all rankings from the database
func (prs *PostgresRankingStorage) GetAllRankings() ([]models.Ranking, error) {
	rows, err := prs.DB.Query("SELECT * FROM rankings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rankings []models.Ranking
	for rows.Next() {
		var ranking models.Ranking

		err := rows.Scan(&ranking.ID, &ranking.UserID, &ranking.PointsScore, &ranking.TimeScore, &ranking.TotalScore)
		if err != nil {
			return nil, err
		}

		rankings = append(rankings, ranking)
	}

	return rankings, nil
}

// GetAllRankingsByUserID retrieves all rankings from the database by user ID
func (prs *PostgresRankingStorage) GetAllRankingsByUserID(userID int) ([]models.Ranking, error) {
	rows, err := prs.DB.Query("SELECT * FROM rankings WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rankings []models.Ranking
	for rows.Next() {
		var ranking models.Ranking

		err := rows.Scan(&ranking.ID, &ranking.UserID, &ranking.PointsScore, &ranking.TimeScore, &ranking.TotalScore)
		if err != nil {
			return nil, err
		}

		rankings = append(rankings, ranking)
	}

	return rankings, nil
}

// CreateRanking creates a new ranking in the database
func (prs *PostgresRankingStorage) CreateRanking(ranking models.Ranking) (models.Ranking, error) {
	query := "INSERT INTO rankings (id, user_id, points_score, time_score, total_score) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	row := prs.DB.QueryRow(query, ranking.ID, ranking.UserID, ranking.PointsScore, ranking.TimeScore, ranking.TotalScore)

	err := row.Scan(&ranking.ID)
	return ranking, err
}

// GetRankingByID retrieves a ranking from the database by ID
func (prs *PostgresRankingStorage) GetRankingByUserID(userID int) (models.Ranking, error) {
	query := "SELECT id, user_id, points_score, time_score, total_score FROM rankings WHERE user_id = $1"
	row := prs.DB.QueryRow(query, userID)

	var ranking models.Ranking
	err := row.Scan(&ranking.ID, &ranking.UserID, &ranking.PointsScore, &ranking.TimeScore, &ranking.TotalScore)
	return ranking, err
}
