package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "user-service/data/users.db"
	}

	db := initDB(dbPath)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("User Service started on port %s\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
