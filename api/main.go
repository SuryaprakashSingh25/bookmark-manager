package main

import (
	"bookmark-api/internal/config"
	"bookmark-api/internal/db"
	"bookmark-api/internal/grpcclient"
	"bookmark-api/internal/handlers"
	"bookmark-api/internal/logger"
	"bookmark-api/internal/middleware"
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
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"https://bookmark-frontend-rho.vercel.app",
			"http://localhost:4200",
		},
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.Login)

	authRoutes := r.Group("/")
	authRoutes.Use(middleware.AuthMiddleware())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	authRoutes.POST("/bookmarks", handlers.CreateBookmark)
	authRoutes.GET("/bookmarks", handlers.GetBookmarks)
	authRoutes.DELETE(
		"/bookmarks/:id",
		handlers.DeleteBookmark,
	)

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
