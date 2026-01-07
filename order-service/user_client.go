package main

import (
	"fmt"
	"net/http"
	"time"
)

type UserClient struct {
	baseURL string
	client  *http.Client
}

func NewUserClient(baseURL string) *UserClient {
	return &UserClient{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 3 * time.Second,
		},
	}
}

// Exists проверяет, существует ли пользователь
func (c *UserClient) Exists(userID int) (bool, error) {
	url := fmt.Sprintf("%s/users/%d", c.baseURL, userID)

	resp, err := c.client.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}
	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	}

	return false, fmt.Errorf("unexpected status from user service: %d", resp.StatusCode)
}
