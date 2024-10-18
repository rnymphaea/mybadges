package models

import (
	"time"

	"github.com/google/uuid"
)

type Badge struct {
	ID           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	ReleaseDate  time.Time `json:"release_date"`
	ImageURL     string    `json:"image_url"`
	Price        int       `json:"price"`
	CategoryID   uuid.UUID `json:"category_id"`
	UserID       uuid.UUID `json:"user_id"`
	CollectionID uuid.UUID `json:"collection_id"`
	Material     string    `json:"material"`
	Weight       int       `json:"weight"`
	Height       int       `json:"height"`
	Width        int       `json:"width"`
	Thickness    int       `json:"thickness"`
	Coverage     string    `json:"coverage"`
	Fastening    string    `json:"fastening"`
}
