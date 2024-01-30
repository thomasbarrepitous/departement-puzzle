package models

type Ranking struct {
	ID          int `json:"id"`
	UserID      int `json:"user_id"`
	PointsScore int `json:"points"`
	TimeScore   int `json:"time"`
	TotalScore  int `json:"total"`
}
