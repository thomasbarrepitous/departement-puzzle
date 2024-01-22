package handlers

import (
	"departement/models"
	"departement/storage"
	"departement/utils"
	"encoding/json"
	"io"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// RankingHandler represents the controller for ranking-related operations
type RankingHandler struct {
	Store storage.RankingStorage
}

// GetAllRankings handles the request to get all rankings
func (rh *RankingHandler) GetAllRankings(w http.ResponseWriter, r *http.Request) {
	rankings, err := rh.Store.GetAllRankings()
	if err != nil {
		utils.JSONRespond(w, http.StatusInternalServerError, err)
		return
	}
	utils.JSONRespond(w, http.StatusOK, rankings)
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

	// Parse the request body into a ranking struct
	var ranking models.Ranking
	parseErr := json.Unmarshal(body, &ranking)
	if parseErr != nil {
		log.Fatal(parseErr)
		utils.JSONRespond(w, http.StatusBadRequest, parseErr)
		return
	}

	// Insert the new ranking into the database
	ranking, dbErr := rh.Store.CreateRanking(ranking)
	if dbErr != nil {
		utils.JSONRespond(w, http.StatusInternalServerError, dbErr)
		return
	}

	utils.JSONRespond(w, http.StatusCreated, ranking)
}
