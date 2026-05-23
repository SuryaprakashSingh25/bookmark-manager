package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port            string
	DBURL           string
	GRPCPreviewAddr string
	JWTSecret       string
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()

	if err != nil {
		log.Println(".env not found, using system env")
	}

	AppConfig = Config{
		Port:            os.Getenv("PORT"),
		DBURL:           os.Getenv("DB_URL"),
		GRPCPreviewAddr: os.Getenv("GRPC_PREVIEW_ADDR"),
		JWTSecret:       os.Getenv("JWT_SECRET"),
	}
}
