package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hellotheremike/go-tasker/internal/middleware"
)

func GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(2 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(middleware.JWT_SECRET))
}
