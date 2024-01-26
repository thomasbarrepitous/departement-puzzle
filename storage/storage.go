package storage

import (
	"context"
	"departement/models"
)

// Storage is an interface that wraps all the storage methods
type Storage interface {
	UserStorage
	RankingStorage
	ProfileStorage
}

type UserStorage interface {
	// GetAllUsers retrieves all users from the database
	GetAllUsers(ctx context.Context) ([]models.User, error)

	// GetUserByID retrieves a user from the database by ID
	GetUserByID(ctx context.Context, id int) (models.User, error)

	// GetUserByEmail retrieves a user from the database by email
	GetUserByEmail(ctx context.Context, email string) (models.User, error)

	// GetUserByUsername retrieves a user from the database by username
	GetUserByUsername(ctx context.Context, username string) (models.User, error)

	// CreateUser creates a new user in the database
	CreateUser(ctx context.Context, user models.User) (models.User, error)

	// UpdateUser updates a user in the database
	UpdateUser(ctx context.Context, user models.User) (models.User, error)

	// DeleteUser deletes a user from the database
	DeleteUser(ctx context.Context, user models.User) error
}

type ProfileStorage interface {
	// GetAllProfiles retrieves all profiles from the database
	GetAllProfiles() ([]models.Profile, error)

	// GetProfileByUserID retrieves a profile from the database by ID
	GetProfileByUserID(userID int) (models.Profile, error)

	// CreateProfile creates a new profile in the database
	CreateProfile(profile models.Profile) (models.Profile, error)

	// UpdateProfile updates a profile in the database
	UpdateProfile(profile models.Profile) (models.Profile, error)

	// DeleteProfile deletes a profile from the database
	DeleteProfile(profile models.Profile) error
}

type RankingStorage interface {
	// GetAllRankings retrieves all rankings from the database
	GetAllRankings() ([]models.Ranking, error)

	// GetAllRankingsByUserID retrieves all rankings from the database by user ID
	GetAllRankingsByUserID(userID int) ([]models.Ranking, error)

	// GetRankingByUserID retrieves a ranking from the database by ID
	GetRankingByUserID(userID int) (models.Ranking, error)

	// CreateRanking creates a new ranking in the database
	CreateRanking(ranking models.Ranking) (models.Ranking, error)
}
