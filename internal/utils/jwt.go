package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewToken(email string, lifetime time.Duration, secret string) (string, error) {
	payload := jwt.MapClaims{
		"sub": email,
		"exp": time.Now().Add(lifetime).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(secret))
}
