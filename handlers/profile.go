package handlers

import (
	"departement/components"
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

	// Sort array by total score
	sort.Slice(rankings, func(i, j int) bool {
		return rankings[i].TotalScore > rankings[j].TotalScore
	})

	component := components.ProfilePageComponent(r, &profile, &rankings)
	component.Render(r.Context(), w)
}
