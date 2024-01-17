package models

type Ranking struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	Score     int64  `json:"points"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
