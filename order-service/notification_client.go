package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type NotificationClient struct {
	baseURL string
	client  *http.Client
}

func NewNotificationClient(baseURL string) *NotificationClient {
	return &NotificationClient{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 2 * time.Second,
		},
	}
}

func (c *NotificationClient) Send(userID int, message string) {
	payload := map[string]interface{}{
		"user_id": userID,
		"message": message,
	}

	body, _ := json.Marshal(payload)

	req, err := http.NewRequest(
		http.MethodPost,
		c.baseURL+"/notify",
		bytes.NewBuffer(body),
	)
	if err != nil {
		log.Printf("notification request build error: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		log.Printf("notification service unavailable: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		log.Printf("notification service returned status %d", resp.StatusCode)
	}
}
