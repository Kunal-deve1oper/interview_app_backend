package utils

import (
	"fmt"
	"os"

	"github.com/Kunal-deve1oper/interview_app_backend/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

// validateJWT parses and validates the token
func ValidateJWT(tokenString string) (*models.UserClaims, error) {
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	claims := &models.UserClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Check required fields
	if claims.UserID == nil ||
		claims.UserID["id"] == "" ||
		claims.UserID["name"] == "" ||
		claims.UserID["email"] == "" ||
		claims.UserID["orgId"] == "" ||
		claims.UserID["orgName"] == "" {
		return nil, fmt.Errorf("missing required claims")
	}

	return claims, nil
}
