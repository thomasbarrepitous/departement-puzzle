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
			ID:          1,
			Name:        "Jeu 1",
			Description: "Description du jeu 1",
			URL:         "/departement",
			Picture:     "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png",
		},
		{
			ID:          2,
			Name:        "Jeu 2",
			Description: "Description du jeu 2",
			URL:         "https://www.google.com",
			Picture:     "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png",
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
