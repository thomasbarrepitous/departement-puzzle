package handlers

import (
	"departement/components"
	"net/http"
)

// HomeHandler is the handler for the home page
type HomeHandler struct{}

// RenderHomePage renders the home page
func (hh *HomeHandler) RenderHomePage(w http.ResponseWriter, r *http.Request) {
	component := components.HomePageComponent(r)
	component.Render(r.Context(), w)
}
