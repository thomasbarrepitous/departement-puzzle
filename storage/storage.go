package storage

import "departement/models"

type UserStorage interface {
	// GetAllUsers retrieves all users from the database
	GetAllUsers() ([]models.User, error)

	// GetUserByEmail retrieves a user from the database by email
	GetUserByEmail(email string) (models.User, error)

	// GetUserByUsername retrieves a user from the database by username
	GetUserByUsername(username string) (models.User, error)

	// CreateUser creates a new user in the database
	CreateUser(user models.User) (models.User, error)
}

type RankingStorage interface {
	// GetAllRankings retrieves all rankings from the database
	GetAllRankings() ([]models.Ranking, error)

	// CreateRanking creates a new ranking in the database
	CreateRanking(ranking models.Ranking) (models.Ranking, error)

	// GetRankingByUserID retrieves a ranking from the database by ID
	GetRankingByUserID(userID int) (models.Ranking, error)
}
