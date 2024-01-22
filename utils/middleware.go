package utils

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

// Package variable for JWT Secret Key
var JwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type AuthMiddleware struct {
	Handler http.Handler
}

// CreateJWT creates a new JWT token with custom claims
func CreateJWT(userID int) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	// Custom claims
	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userID,
		"exp":        expirationTime.Unix(),
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString(JwtSecretKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func JWTVerifyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	// Verify the token and call the next handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		// Check if the token is present
		if authHeader == "" {
			// Token is missing, return 403
			http.Error(w, "Authorization header is missing", http.StatusForbidden)
			return
		}

		// Bearer token check
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			// Token is malformed, return 403
			http.Error(w, "Malformed token", http.StatusForbidden)
			return
		}

		// Check if the token is valid
		tokenPart := parts[1]
		token, err := jwt.Parse(tokenPart, func(token *jwt.Token) (interface{}, error) {
			// Check if the signing method is valid
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				// Signing method is not valid, return 403
				return nil, fmt.Errorf("Signing method is not valid: %v", token.Header["alg"])
			}
			return JwtSecretKey, nil
		})
		if err != nil {
			// Token is invalid, return 403
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		// Check if the token claims are valid
		if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
			// Token is not valid, return 403
			http.Error(w, "Token is not valid", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
