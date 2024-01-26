package handlers

import (
	"departement/components"
	"departement/storage"
	"net/http"
)

type ProfileHandler struct {
	Store storage.ProfileStorage
}

func (ph *ProfileHandler) RenderProfilePage(w http.ResponseWriter, r *http.Request) {
	component := components.ProfilePageComponent(r)
	component.Render(r.Context(), w)
}
