package models

import "time"

type Bookmark struct {
	ID          int64    `json:"id"`
	URL         string   `json:"url"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
