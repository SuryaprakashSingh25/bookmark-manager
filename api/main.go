package main

import (
	"bookmark-api/internal/db"
	"bookmark-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.POST("/bookmarks", handlers.CreateBookmark)
	r.GET("/bookmarks", handlers.GetBookmarks)
	r.Run(":8080")
}
