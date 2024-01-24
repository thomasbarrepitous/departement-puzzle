package utils

import (
	"context"
	"fmt"
	"net/http"
	"os"
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

func JWTVerifyMiddleware(next http.Handler) http.Handler {
	// Verify the token and call the next handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the Authorization cookie is present
		authCookie, err := r.Cookie("Authorization")
		if err != nil {
			// Cookie is not present, return 403
			// http.Error(w, "Cookie is not present", http.StatusForbidden)
			http.Redirect(w, r, "/404", http.StatusSeeOther)
			return
		}

		// Check if the token is valid
		token, err := jwt.Parse(authCookie.Value, func(token *jwt.Token) (interface{}, error) {
			// Check if the signing method is valid
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				// Signing method is not valid, return 403
				return nil, fmt.Errorf("Signing method is not valid: %v", token.Header["alg"])
			}
			return JwtSecretKey, nil
		})
		if err != nil {
			// Token is invalid, return 403
			// http.Error(w, "Invalid token", http.StatusForbidden)
			http.Redirect(w, r, "/404", http.StatusSeeOther)
			return
		}

		// Check if the token claims are valid
		if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
			// Token is not valid, return 403
			// http.Error(w, "Token is not valid", http.StatusForbidden)
			http.Redirect(w, r, "/404", http.StatusSeeOther)
			return
		}

		// Add the user ID to the context
		ctx := context.WithValue(r.Context(), "user_id", token.Claims.(jwt.MapClaims)["user_id"])
		ctx = context.WithValue(ctx, "authorized", token.Claims.(jwt.MapClaims)["authorized"])

		// Call the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
