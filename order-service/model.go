package main

import "time"

type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Product   string    `json:"product"`
	CreatedAt time.Time `json:"created_at"`
}
