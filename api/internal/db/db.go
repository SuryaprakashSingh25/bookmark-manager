package db

import (
	"log"

	"github.com/jackc/pgx"
)

var Conn *pgx.Conn

func InitDB() {
	var err error
	connStr := "postgres://postgres:postgres@localhost:5432/bookmarkdb"

	cfg, err := pgx.ParseURI(connStr)
	if err != nil {
		log.Fatal("Unable to parse DB URI:", err)
	}

	Conn, err = pgx.Connect(cfg)
	if err != nil {
		log.Fatal("Unable to connect to DB:", err)
	}
	log.Println("Connected to PostgreSQL")
}
