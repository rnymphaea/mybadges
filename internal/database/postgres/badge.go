package postgres

import (
	"context"

	"mybadges/internal/database/models"
)

func (s *Storage) AddBadge(badge models.Badge) error {
	_, err := s.pool.Exec(context.Background(), "insert into badges(id, user_id, title, description, "+
		"image_url, category_id, collection_id, material, weight, height, width, thickness, coverage, fastening)"+
		"values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)",
		badge.ID,
		badge.UserID,
		badge.Title,
		badge.Description,
		badge.ImageURL,
		badge.CategoryID,
		badge.CollectionID,
		badge.Material,
		badge.Weight,
		badge.Height,
		badge.Width,
		badge.Thickness,
		badge.Coverage,
		badge.Fastening)
	if err != nil {
		return err
	}
	return nil
}
