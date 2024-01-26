package models

type Profile struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id"`
	Username    string `json:"last_name"`
	Email       string `json:"email"`
	Picture     string `json:"picture"`
	Description string `json:"description"`
}
