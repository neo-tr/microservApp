package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserClient_Exists_OK(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewUserClient(server.URL)

	exists, err := client.Exists(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !exists {
		t.Fatal("expected user to exist")
	}
}

func TestUserClient_Exists_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := NewUserClient(server.URL)

	exists, err := client.Exists(999)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if exists {
		t.Fatal("expected user to not exist")
	}
}

func TestUserClient_Exists_ServiceDown(t *testing.T) {
	client := NewUserClient("http://localhost:59999") // несуществующий порт

	_, err := client.Exists(1)
	if err == nil {
		t.Fatal("expected error when service is down")
	}
}
