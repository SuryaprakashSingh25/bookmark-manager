package main

import (
	"bookmark-api/internal/config"
	"bookmark-api/internal/db"
	"bookmark-api/internal/grpcclient"
	"bookmark-api/internal/handlers"
	"bookmark-api/internal/logger"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	logger.InitLogger()
	defer logger.Log.Sync()
	db.InitDB()
	grpcclient.InitGRPC()
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.POST("/bookmarks", handlers.CreateBookmark)
	r.GET("/bookmarks", handlers.GetBookmarks)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(
		quit,
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	<-quit

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()
	srv.Shutdown(ctx)
}
