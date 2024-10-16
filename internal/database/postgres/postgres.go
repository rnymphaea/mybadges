package postgres

import (
	"context"
	stderrors "errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"

	"mybadges/internal/database/models"
	"mybadges/internal/utils/errors"
)

type Storage struct {
	pool *pgxpool.Pool
}

func New(path string) (*Storage, error) {
	pool, err := pgxpool.New(context.Background(), path)
	if err != nil {
		return nil, err
	}
	return &Storage{pool: pool}, nil
}

func (s *Storage) CreateUser(user models.User) error {
	if exists, _ := s.userExists(user.Email); exists {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password: %v", err)
	}

	_, err = s.pool.Exec(context.Background(),
		"insert into users(id, email, password, created_at) values ($1, $2, $3, $4)",
		user.ID,
		user.Email,
		hashedPassword,
		user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) userExists(email string) (bool, error) {
	var id uuid.UUID
	err := s.pool.QueryRow(context.Background(), "select id from users where email = $1", email).Scan(&id)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return true, nil
}

func (s *Storage) CheckCredentials(email, password string) error {
	var hashedPassword string
	err := s.pool.QueryRow(context.Background(), "select password from users where email = $1", email).Scan(&hashedPassword)
	if err != nil {
		if stderrors.Is(err, pgx.ErrNoRows) {
			return errors.ErrUserNotFound
		} else {
			return errors.ErrCheckingPassword
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.ErrInvalidCredentials
	} else {
		return nil
	}
}
