package repositories

import (
	"bookmark-api/internal/db"
	"bookmark-api/internal/models"
	"context"
)

type BookmarkRepository struct{}

func (r *BookmarkRepository) Create(
	ctx context.Context,
	url string,
	title string,
	description string,
) (*models.Bookmark, error) {
	query := `
	INSERT INTO bookmarks (url,title,description)
	VALUES ($1,$2,$3)
	RETURNING id,
			  url,
			  COALESCE(title,''),
			  COALESCE(description,''),
			  created_at
	`
	var bookmark models.Bookmark

	err := db.Conn.QueryRow(
		query,
		url,
		title,
		description,
	).Scan(
		&bookmark.ID,
		&bookmark.URL,
		&bookmark.Title,
		&bookmark.Description,
		&bookmark.CreatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &bookmark, nil
}
