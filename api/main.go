package main

import (
	"bookmark-api/internal/db"
	"bookmark-api/internal/grpcclient"
	"bookmark-api/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	grpcclient.InitGRPC()
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.POST("/bookmarks", handlers.CreateBookmark)
	r.GET("/bookmarks", handlers.GetBookmarks)
	r.Run(":8080")
}
