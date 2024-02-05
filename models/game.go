package models

type Game struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Picture     string `json:"picture"`
	Available   bool   `json:"available"`
}

func NewGame(name string, description string, url string, picture string, available bool) Game {
	return Game{
		Name:        name,
		Description: description,
		URL:         url,
		Picture:     picture,
		Available:   available,
	}
}
