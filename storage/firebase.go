package storage

import (
	"context"
	"departement/models"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

type FirebaseStorage struct {
	Auth *auth.Client
}

func NewFirebaseStorage(ctx context.Context) *FirebaseStorage {
	opt := option.WithCredentialsFile("firebase.json")
	// Initialize the firebase app
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	// Initialize the auth client
	authClient, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	return &FirebaseStorage{
		Auth: authClient,
	}
}

func (s *FirebaseStorage) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	user, err := s.Auth.GetUserByEmail(ctx, email)
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		Username: user.DisplayName,
		Email:    user.Email,
	}, nil
}

func (s *FirebaseStorage) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	user, err := s.Auth.GetUserByEmail(ctx, username)
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		Username: user.DisplayName,
		Email:    user.Email,
	}, nil
}

func (s *FirebaseStorage) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	params := (&auth.UserToCreate{}).
		Email(user.Email).
		EmailVerified(false).
		Password(user.Password).
		DisplayName(user.Username).
		Disabled(false)

	u, err := s.Auth.CreateUser(ctx, params)
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		Username: u.DisplayName,
		Email:    u.Email,
	}, nil
}

func (s *FirebaseStorage) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return []models.User{}, nil
}

func (s *FirebaseStorage) GetAllRankings() ([]models.Ranking, error) {
	return []models.Ranking{}, nil
}

func (s *FirebaseStorage) CreateRanking(ranking models.Ranking) (models.Ranking, error) {
	return models.Ranking{}, nil
}

func (s *FirebaseStorage) GetRankingByUserID(userID int) (models.Ranking, error) {
	return models.Ranking{}, nil
}
