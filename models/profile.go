package models

import "github.com/google/uuid"

type Profile struct {
	ID          int       `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	Username    string    `json:"last_name"`
	Email       string    `json:"email"`
	Picture     string    `json:"picture"`
	Description string    `json:"description"`
	Country     string    `json:"country"`
}

func NewProfile(userID uuid.UUID, username string, email string, picture string, description string, country string) Profile {
	return Profile{
		UserID:      userID,
		Username:    username,
		Email:       email,
		Picture:     picture,
		Description: description,
		Country:     country,
	}
}
