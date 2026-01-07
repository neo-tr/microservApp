package main

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupOrderTestDB(t *testing.T) *sql.DB {
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

func TestOrderRepository_CreateAndList(t *testing.T) {
	db := setupOrderTestDB(t)
	repo := NewOrderRepository(db)

	_, err := repo.Create(1, "Book")
	if err != nil {
		t.Fatalf("failed to create order: %v", err)
	}

	orders, err := repo.List()
	if err != nil {
		t.Fatalf("failed to list orders: %v", err)
	}

	if len(orders) != 1 {
		t.Fatalf("expected 1 order, got %d", len(orders))
	}

	if orders[0].Product != "Book" {
		t.Fatalf("unexpected product: %s", orders[0].Product)
	}
}
