package handlers

import (
	"departement/components"
	"departement/models"
	"net/http"
	"strings"
)

type PlayMenuHandler struct{}

func (pmh *PlayMenuHandler) GetGames(w http.ResponseWriter, r *http.Request) {
	searchQuery := r.URL.Query().Get("search")
	games := []models.Game{
		{
			ID:          "1",
			Name:        "Guess the French departement",
			Description: "Guess the french departement location by its number and its name!.",
			URL:         "/departement",
			Picture:     "/static/assets/fr_dep_logo.png",
			Available:   true,
		},
		{
			ID:          "2",
			Name:        "Guess the Japanese prefecture",
			Description: "Guess the japanese prefecture location by its logo and its name!.",
			URL:         "",
			Picture:     "/static/assets/jp_pref_logo.jpeg",
			Available:   false,
		},
	}
	var filteredGames []models.Game
	for _, game := range games {
		// Check if the game name contains the search query (case-insensitive)
		if strings.Contains(strings.ToLower(game.Name), strings.ToLower(searchQuery)) {
			filteredGames = append(filteredGames, game)
		}
	}
	component := components.GameCardComponent(r, &filteredGames)
	component.Render(r.Context(), w)
}

func (pmh *PlayMenuHandler) RenderPlayMenuPage(w http.ResponseWriter, r *http.Request) {
	component := components.PlayMenuPageComponent(r)
	component.Render(r.Context(), w)
}
