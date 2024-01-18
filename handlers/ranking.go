package handlers

import (
	"database/sql"
	"departement/models"
	"departement/utils"
	"encoding/json"
	"io"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// RankingHandler represents the controller for ranking-related operations
type RankingHandler struct {
	DB *sql.DB
}

// getBestRankingsInDB queries the database to get the best rankings
func (rh *RankingHandler) getBestRankingsInDB() ([]models.Ranking, error) {
	rows, err := rh.DB.Query("SELECT * FROM rankings ORDER BY score DESC LIMIT 10")
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

// GetAllRankingsInDB queries the database to get all rankings
func (rh *RankingHandler) GetAllRankingsInDB() ([]models.Ranking, error) {
	rows, err := rh.DB.Query("SELECT * FROM rankings")
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

// GetAllRankings handles the request to get all rankings
func (rh *RankingHandler) GetAllRankings(w http.ResponseWriter, r *http.Request) {
	rankings, err := rh.GetAllRankingsInDB()
	if err != nil {
		utils.JSONRespond(w, http.StatusInternalServerError, err)
		return
	}
	utils.JSONRespond(w, http.StatusOK, rankings)
}

// createRankingInDB creates a new ranking in the database
func (rh *RankingHandler) createRankingInDB(ranking models.Ranking) (models.Ranking, error) {
	_, err := rh.DB.Exec("INSERT INTO rankings (user_id, score) VALUES ($1, $2)", ranking.UserID, ranking.Score)
	if err != nil {
		return ranking, err
	}
	return ranking, nil
}

// CreateRanking handles the request to create a new ranking
func (rh *RankingHandler) CreateRanking(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		log.Fatal(readErr)
		utils.JSONRespond(w, http.StatusBadRequest, readErr)
		return
	}

	// Parse the request body
	var ranking models.Ranking
	parseErr := json.Unmarshal(body, &ranking)
	if parseErr != nil {
		log.Fatal(parseErr)
		utils.JSONRespond(w, http.StatusBadRequest, parseErr)
		return
	}

	// Insert the new ranking into the database
	ranking, dbErr := rh.createRankingInDB(ranking)
	if dbErr != nil {
		utils.JSONRespond(w, http.StatusInternalServerError, dbErr)
		return
	}

	utils.JSONRespond(w, http.StatusCreated, ranking)
}
