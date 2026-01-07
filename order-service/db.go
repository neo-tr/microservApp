package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func initDB(dbPath string) *sql.DB {
	if err := os.MkdirAll("./data", 0755); err != nil {
		log.Fatalf("failed to create data directory: %v", err)
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY,
		user_id INTEGER NOT NULL,
		product TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("failed to create orders table: %v", err)
	}

	log.Printf("Orders DB initialized at %s\n", dbPath)
	return db
}
