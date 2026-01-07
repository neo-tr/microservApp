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
		dbPath = "./data/orders.db"
	}

	db := initDB(dbPath)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	repo := NewOrderRepository(db)

	userServiceURL := os.Getenv("USER_SERVICE_URL")
	if userServiceURL == "" {
		userServiceURL = "http://localhost:5050"
	}

	userClient := NewUserClient(userServiceURL)

	notificationURL := os.Getenv("NOTIFICATION_SERVICE_URL")
	if notificationURL == "" {
		notificationURL = "http://localhost:5052"
	}

	notificationClient := NewNotificationClient(notificationURL)

	handler := NewOrderHandler(repo, userClient, notificationClient)

	mux.HandleFunc("/orders", handler.CreateOrder)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Order Service started on port %s\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
