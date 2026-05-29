package repositories

import (
	"bookmark-api/internal/db"
	"bookmark-api/internal/models"
	"context"
)

type BookmarkRepository struct{}

func (r *BookmarkRepository) Create(
	ctx context.Context,
	userID int64,
	url string,
	title string,
	description string,
) (*models.Bookmark, error) {
	query := `
	INSERT INTO bookmarks (user_id,url,title,description)
	VALUES ($1,$2,$3,$4)
	RETURNING id,
			  url,
			  COALESCE(title,''),
			  COALESCE(description,''),
			  created_at
	`
	var bookmark models.Bookmark

	err := db.Conn.QueryRow(
		query,
		userID,
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

func (r *BookmarkRepository) Delete(
	ctx context.Context,
	id int,
	userID int64,
) error {
	query := `
		DELETE FROM bookmarks
		WHERE id = $1
		AND user_id = $2
	`
	_, err := db.Conn.Exec(
		query,
		id,
		userID,
	)
	return err
}
