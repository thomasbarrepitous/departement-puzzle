package storage

import (
	"context"
	"departement/models"

	"github.com/google/uuid"
)

// IStorage is an interface that wraps all the storage methods
type IStorage interface {
	UserStorage
	RankingStorage
	ProfileStorage
}

// Storage is a struct that wraps all the storage methods
type Storage struct {
	Users    UserStorage
	Rankings RankingStorage
	Profiles ProfileStorage
}

func NewStorage(env string) *Storage {
	switch env {
	case "dev":
		return &Storage{
			Users:    NewPostgresUserStorage(),
			Rankings: NewPostgresRankingStorage(),
			Profiles: NewPostgresProfileStorage(),
		}
	default:
		return &Storage{
			Users:    NewPostgresUserStorage(),
			Rankings: NewPostgresRankingStorage(),
			Profiles: NewPostgresProfileStorage(),
		}
	}
}

// UserStorage is an interface that wraps all the user storage methods
type UserStorage interface {
	// GetAllUsers retrieves all users from the database
	GetAllUsers(ctx context.Context) ([]models.User, error)

	// GetUserByID retrieves a user from the database by ID
	GetUserByID(ctx context.Context, id uuid.UUID) (models.User, error)

	// GetUserByEmail retrieves a user from the database by email
	GetUserByEmail(ctx context.Context, email string) (models.User, error)

	// GetUserByUsername retrieves a user from the database by username
	GetUserByUsername(ctx context.Context, username string) (models.User, error)

	// CreateUser creates a new user in the database
	CreateUser(ctx context.Context, user models.User) (models.User, error)

	// UpdateUser updates a user in the database
	UpdateUser(ctx context.Context, id uuid.UUID, user models.User) (models.User, error)

	// DeleteUser deletes a user from the database
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

// ProfileStorage is an interface that wraps all the profile storage methods
type ProfileStorage interface {
	// GetAllProfiles retrieves all profiles from the database
	GetAllProfiles(ctx context.Context) ([]models.Profile, error)

	// GetProfileByUserID retrieves a profile from the database by ID
	GetProfileByUserID(ctx context.Context, userID uuid.UUID) (models.Profile, error)

	// CreateProfile creates a new profile in the database
	CreateProfile(ctx context.Context, profile models.Profile) (models.Profile, error)

	// UpdateProfile updates a profile in the database
	UpdateProfile(ctx context.Context, id uuid.UUID, profile models.Profile) (models.Profile, error)

	// DeleteProfile deletes a profile from the database
	DeleteProfile(ctx context.Context, id uuid.UUID) error
}

// RankingStorage is an interface that wraps all the ranking storage methods
type RankingStorage interface {
	// GetAllRankings retrieves all rankings from the database
	// If the user has no rankings, an empty slice is returned
	GetAllRankings(ctx context.Context) ([]models.Ranking, error)

	// GetAllRankingsByUserID retrieves all rankings from the database by user ID
	// If the user has no rankings, an empty slice is returned
	GetAllRankingsByUserID(ctx context.Context, userID uuid.UUID) ([]models.Ranking, error)

	// CreateRanking creates a new ranking in the database
	CreateRanking(ctx context.Context, ranking models.Ranking) (models.Ranking, error)
}
