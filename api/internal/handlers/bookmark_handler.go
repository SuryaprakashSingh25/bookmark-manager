package handlers

import (
	"bookmark-api/internal/db"
	"bookmark-api/internal/models"
	"bookmark-api/internal/services"
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

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

	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		3*time.Second,
	)
	defer cancel()

	userID := c.MustGet("user_id").(int64)

	bookmark, err := bookmarkService.CreateBookmark(
		ctx,
		userID,
		req.URL,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create bookmark",
		})
		return
	}

	c.JSON(http.StatusOK, bookmark)
}

func GetBookmarks(c *gin.Context) {
	userID := c.MustGet("user_id").(int64)
	query := `
		SELECT id, url, COALESCE(title, ''), COALESCE(description, ''), created_at
		FROM bookmarks
		WHERE user_id=$1
		ORDER BY created_at DESC
	`

	rows, err := db.Conn.Query(query, userID)
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

func DeleteBookmark(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	userID := c.MustGet("user_id").(int64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}
	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		3*time.Second,
	)
	defer cancel()
	err = bookmarkService.DeleteBookmark(
		ctx, id, userID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete bookmark",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "bookmark deleted",
	})
}
