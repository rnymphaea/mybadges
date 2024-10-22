package database

import (
	"io"

	"github.com/google/uuid"

	"mybadges/internal/database/models"
)

type UserRepository interface {
	CreateUser(user models.User) error
	CheckCredentials(email, password string) error
	GetUserIDByEmail(email string) (uuid.UUID, error)
}

type BadgeRepository interface {
	AddBadge(badge models.Badge) error
}

type ImageRepository interface {
	UploadFile(file io.Reader, key string) (string, error)
}

func GenerateUUID() uuid.UUID {
	return uuid.New()
}
