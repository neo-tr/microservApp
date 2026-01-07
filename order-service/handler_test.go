package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

type mockNotificationClient struct{}

func (m *mockNotificationClient) Send(userID int, message string) {
	// ничего не делаем
}

func setupOrderHandlerTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(`
	CREATE TABLE orders (
		id INTEGER PRIMARY KEY,
		user_id INTEGER NOT NULL,
		product TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`)
	if err != nil {
		t.Fatal(err)
	}

	return db
}

func TestCreateOrder_OK(t *testing.T) {
	userServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer userServer.Close()

	db := setupOrderHandlerTestDB(t)
	repo := NewOrderRepository(db)
	userClient := NewUserClient(userServer.URL)
	notificationClient := &mockNotificationClient{}

	handler := NewOrderHandler(repo, userClient, notificationClient)

	body := map[string]interface{}{
		"user_id": 1,
		"product": "Book",
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewReader(jsonBody))
	rec := httptest.NewRecorder()

	handler.CreateOrder(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", rec.Code)
	}
}

func TestCreateOrder_UserNotFound(t *testing.T) {
	userServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer userServer.Close()

	db := setupOrderHandlerTestDB(t)
	repo := NewOrderRepository(db)
	userClient := NewUserClient(userServer.URL)
	notificationClient := &mockNotificationClient{}

	handler := NewOrderHandler(repo, userClient, notificationClient)

	body := map[string]interface{}{
		"user_id": 999,
		"product": "Book",
	}
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewReader(jsonBody))
	rec := httptest.NewRecorder()

	handler.CreateOrder(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", rec.Code)
	}
}
