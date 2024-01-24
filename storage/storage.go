package storage

import (
	"context"
	"departement/models"
)

// Storage is an interface that wraps all the storage methods
type Storage interface {
	UserStorage
	RankingStorage
}

type UserStorage interface {
	// GetAllUsers retrieves all users from the database
	GetAllUsers(ctx context.Context) ([]models.User, error)

	// GetUserByEmail retrieves a user from the database by email
	GetUserByEmail(ctx context.Context, email string) (models.User, error)

	// GetUserByUsername retrieves a user from the database by username
	GetUserByUsername(ctx context.Context, username string) (models.User, error)

	// CreateUser creates a new user in the database
	CreateUser(ctx context.Context, user models.User) (models.User, error)
}

type RankingStorage interface {
	// GetAllRankings retrieves all rankings from the database
	GetAllRankings() ([]models.Ranking, error)

	// CreateRanking creates a new ranking in the database
	CreateRanking(ranking models.Ranking) (models.Ranking, error)

	// GetRankingByUserID retrieves a ranking from the database by ID
	GetRankingByUserID(userID int) (models.Ranking, error)
}
