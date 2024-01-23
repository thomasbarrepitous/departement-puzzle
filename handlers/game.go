package handlers

import (
	"departement/components"
	"net/http"
)

// GameHandler is the handler for the game page
type GameHandler struct{}

func (gh *GameHandler) RenderGamePage(w http.ResponseWriter, r *http.Request) {
	component := components.GamePageComponent(r)
	component.Render(r.Context(), w)
}
