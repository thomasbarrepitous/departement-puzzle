package handlers

import (
	"departement/components"
	"log"
	"net/http"
)

// GameHandler is the handler for the game page
type GameHandler struct{}

func (gh *GameHandler) RenderGamePage(w http.ResponseWriter, r *http.Request) {
	log.Print(r.Context().Value("user"))
	log.Print(r.Context().Value("authorized"))

	component := components.GamePageComponent(r)
	component.Render(r.Context(), w)
}
