package db

import "os"

type AuthConfig struct {
	apiKey            string
	authDomain        string
	databaseURL       string
	projectId         string
	storageBucket     string
	messagingSenderId string
	appId             string
}

func NewAuthConfig() *AuthConfig {
	firebaseConfig := &AuthConfig{
		apiKey:            os.Getenv("FIREBASE_API_KEY"),
		authDomain:        os.Getenv("FIREBASE_AUTH_DOMAIN"),
		databaseURL:       os.Getenv("FIREBASE_DATABASE_URL"),
		projectId:         os.Getenv("FIREBASE_PROJECT_ID"),
		storageBucket:     os.Getenv("FIREBASE_STORAGE_BUCKET"),
		messagingSenderId: os.Getenv("FIREBASE_MESSAGING_SENDER_ID"),
		appId:             os.Getenv("FIREBASE_APP_ID"),
	}
	return firebaseConfig
}
