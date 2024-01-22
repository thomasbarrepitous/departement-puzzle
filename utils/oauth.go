package utils

import (
	"context"
	"encoding/json"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleProfile struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func CreateGoogleOAuth2Config() *oauth2.Config {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:3000/api/auth/google/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}

func GetGoogleProfile(ctx context.Context, config *oauth2.Config, accessToken *oauth2.Token) (*GoogleProfile, error) {
	var profile GoogleProfile

	// Call the Google OAuth2 API to fetch the user's profile information
	client := config.Client(ctx, accessToken)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	// Decode the JSON response into the GoogleProfile struct
	json.NewDecoder(response.Body).Decode(&profile)

	return &profile, nil
}
