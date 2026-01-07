package main

import (
	"encoding/json"
	"net/http"
)

type OrderHandler struct {
	repo       *OrderRepository
	userClient *UserClient
}

func NewOrderHandler(repo *OrderRepository, userClient *UserClient) *OrderHandler {
	return &OrderHandler{
		repo:       repo,
		userClient: userClient,
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		UserID  int    `json:"user_id"`
		Product string `json:"product"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if input.UserID <= 0 || input.Product == "" {
		http.Error(w, "user_id and product are required", http.StatusBadRequest)
		return
	}

	exists, err := h.userClient.Exists(input.UserID)
	if err != nil {
		http.Error(w, "user service unavailable", http.StatusServiceUnavailable)
		return
	}
	if !exists {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	order, err := h.repo.Create(input.UserID, input.Product)
	if err != nil {
		http.Error(w, "failed to create order", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
