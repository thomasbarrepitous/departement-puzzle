package handlers

import (
	"departement/components"
	"departement/models"
	"departement/storage"
	"net/http"
	"sort"
)

type ProfileHandler struct {
	ProfileStore storage.ProfileStorage
	RankingStore storage.RankingStorage
}

func (ph *ProfileHandler) RenderProfilePage(w http.ResponseWriter, r *http.Request) {
	userID := int(r.Context().Value("user_id").(float64))

	// Get the profile from the database
	profile, err := ph.ProfileStore.GetProfileByUserID(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the best rankings from the user
	rankings, err := ph.RankingStore.GetAllRankingsByUserID(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rankings = []models.Ranking{
		{
			ID:          1,
			UserID:      1,
			PointsScore: 101,
			TimeScore:   102,
			TotalScore:  203,
		},
		{
			ID:          2,
			UserID:      1,
			PointsScore: 100,
			TimeScore:   100,
			TotalScore:  200,
		},
		{
			ID:          3,
			UserID:      1,
			PointsScore: 1020,
			TimeScore:   1020,
			TotalScore:  2040,
		},
		{
			ID:          8,
			UserID:      1,
			PointsScore: 1000,
			TimeScore:   1000,
			TotalScore:  2000,
		},
		{
			ID:          51,
			UserID:      1,
			PointsScore: 100,
			TimeScore:   200,
			TotalScore:  300,
		},
		{
			ID:          6,
			UserID:      1,
			PointsScore: 3920,
			TimeScore:   50,
			TotalScore:  3970,
		},
	}

	// Sort array by total score
	sort.Slice(rankings, func(i, j int) bool {
		return rankings[i].TotalScore > rankings[j].TotalScore
	})

	component := components.ProfilePageComponent(r, &profile, &rankings)
	component.Render(r.Context(), w)
}
