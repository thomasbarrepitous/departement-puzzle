package handlers

import (
	"departement/components"
	"departement/storage"
	"net/http"
)

type ProfileHandler struct {
	ProfileStore storage.ProfileStorage
}

func (ph *ProfileHandler) RenderProfilePage(w http.ResponseWriter, r *http.Request) {
	// Get the profile from the database
	profile, err := ph.ProfileStore.GetProfileByUserID(r.Context(), 1)
	if err != nil {
	}
	component := components.ProfilePageComponent(r, &profile, nil)
	component.Render(r.Context(), w)
}
