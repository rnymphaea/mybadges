package database

import (
	"github.com/google/uuid"

	"mybadges/internal/database/models"
)

type UserRepository interface {
	CreateUser(user models.User) error
}

func GenerateUUID() uuid.UUID {
	return uuid.New()
}
