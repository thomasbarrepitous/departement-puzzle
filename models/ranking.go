package models

type Ranking struct {
	ID          int64 `json:"id"`
	UserID      int64 `json:"user_id"`
	PointsScore int64 `json:"score"`
	TimeScore   int64 `json:"time"`
}
