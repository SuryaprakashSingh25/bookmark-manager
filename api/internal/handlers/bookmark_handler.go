package handlers

import (
	"bookmark-api/internal/db"
	"bookmark-api/internal/models"
	"bookmark-api/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateBookmarkRequest struct {
	URL string `json:"url" binding:"required"`
}

var bookmarkService = services.NewBookmarkService()

func CreateBookmark(c *gin.Context) {
	var req CreateBookmarkRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookmark, err := bookmarkService.CreateBookmark(req.URL)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create bookmark",
		})
		return
	}

	c.JSON(http.StatusOK, bookmark)
}

func GetBookmarks(c *gin.Context) {
	query := `
		SELECT id, url, COALESCE(title, ''), COALESCE(description, ''), created_at
		FROM bookmarks
		ORDER BY created_at DESC
	`

	rows, err := db.Conn.Query(query)
	if err != nil {
		log.Printf("GetBookmarks query error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch bookmarks"})
		return
	}
	defer rows.Close()

	var bookmarks []models.Bookmark

	for rows.Next() {
		var b models.Bookmark

		err := rows.Scan(
			&b.ID,
			&b.URL,
			&b.Title,
			&b.Description,
			&b.CreatedAt,
		)
		if err != nil {
			continue
		}
		bookmarks = append(bookmarks, b)
	}
	c.JSON(http.StatusOK, bookmarks)
}
