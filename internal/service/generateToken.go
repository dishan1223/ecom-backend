package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken function takes a user id and generates a jwt token that will be sent
// to the client
func GenerateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24 * 365).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(JwtSecret))
}
