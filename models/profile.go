package models

type Profile struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Username    string `json:"last_name"`
	Email       string `json:"email"`
	Picture     string `json:"picture"`
	Description string `json:"description"`
	Country     string `json:"country"`
}

func NewProfile(userID int, username string, email string) Profile {
	return Profile{
		UserID:   userID,
		Username: username,
		Email:    email,
	}
}
