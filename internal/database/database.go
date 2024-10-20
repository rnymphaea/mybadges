package database

import (
	"github.com/google/uuid"

	"mybadges/internal/database/models"
)

type UserRepository interface {
	CreateUser(user models.User) error
	CheckCredentials(email, password string) error
}

type ImageRepository interface {
	UploadFile(filename string) (string, error)
}

func GenerateUUID() uuid.UUID {
	return uuid.New()
}
