package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID, username string, roleID int) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", fmt.Errorf("JWT_SECRET is not set")
	}

	durationMinutes := 60 // default
	if v := os.Getenv("JWT_EXPIRE_MINUTES"); v != "" {
	}

	claims := jwt.MapClaims{
		"sub":      userID,
		"username": username,
		"role_id":  roleID,
		"exp":      time.Now().Add(time.Minute * time.Duration(durationMinutes)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
