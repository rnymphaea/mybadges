package utils

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"mybadges/internal/utils/errors"
)

func NewToken(email string, lifetime time.Duration, secret string) (string, error) {
	payload := jwt.MapClaims{
		"sub": email,
		"exp": time.Now().Add(lifetime).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(secret))
}

func GetTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.ErrNoAuthHeader
	}

	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	if tokenString == "" {
		return "", errors.ErrNoBearerToken
	}

	return tokenString, nil
}

func GetEmailFromToken(tokenString, secret string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid claims")
	}

	email, ok := claims["sub"].(string)
	if !ok {
		return "", fmt.Errorf("email not found in token claims")
	}

	return email, nil
}
