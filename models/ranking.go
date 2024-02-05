package models

import "github.com/google/uuid"

type Ranking struct {
	ID          uuid.UUID `json:"ranking_id"`
	UserID      uuid.UUID `json:"user_id"`
	PointsScore int       `json:"points"`
	TimeScore   int       `json:"time"`
	TotalScore  int       `json:"total"`
}
