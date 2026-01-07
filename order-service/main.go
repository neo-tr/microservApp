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
		w.Write([]byte("OK"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "5001"
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./data/users.db"
	}

	db := initDB(dbPath)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Order Service started on port %s\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
