package db

import (
	"log"
	"time"

	"github.com/jackc/pgx"
)

var Conn *pgx.Conn

func InitDB() {
	var err error
	connStr := "postgres://postgres:postgres@postgres:5432/bookmarkdb"

	cfg, err := pgx.ParseURI(connStr)
	if err != nil {
		log.Fatal("Unable to parse DB URI:", err)
	}

	// Retry connection up to 5 times with 2 second delay
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		Conn, err = pgx.Connect(cfg)
		if err == nil {
			break
		}
		if i < maxRetries-1 {
			log.Printf("Unable to connect to DB (attempt %d/%d): %v. Retrying in 2s...", i+1, maxRetries, err)
			time.Sleep(2 * time.Second)
		}
	}
	if err != nil {
		log.Fatal("Unable to connect to DB after retries:", err)
	}
	log.Println("Connected to PostgreSQL")

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS bookmarks (
	id SERIAL PRIMARY KEY,
	url TEXT NOT NULL,
	title TEXT,
	description TEXT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err = Conn.Exec(createTableQuery)

	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

}
