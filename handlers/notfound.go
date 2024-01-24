package handlers

import (
	"departement/components"
	"net/http"
)

// NotFoundHandler is the handler for the not found page
type NotFoundHandler struct{}

func (nfh *NotFoundHandler) RenderNotFoundPage(w http.ResponseWriter, r *http.Request) {
	component := components.NotFoundPageComponent(r)
	component.Render(r.Context(), w)
}
