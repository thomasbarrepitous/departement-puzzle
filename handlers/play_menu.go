package handlers

import (
	"departement/components"
	"net/http"
)

type PlayMenuHandler struct{}

func (pmh *PlayMenuHandler) RenderPlayMenuPage(w http.ResponseWriter, r *http.Request) {
	component := components.PlayMenuPageComponent(r)
	component.Render(r.Context(), w)
}
