package transport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"user-simulator/internal/model"
)

type HTTPClient struct {
	BaseURL string
	Client  *http.Client
}

func NewHTTPClient(baseURL string) *HTTPClient {
	return &HTTPClient{
		BaseURL: baseURL,
		Client:  &http.Client{},
	}
}

func (h *HTTPClient) SendStatus(u *model.User) error {
	data, err := json.Marshal(u)
	if err != nil {
		return err
	}

	resp, err := h.Client.Post(fmt.Sprintf("%s/api/user/status", h.BaseURL),
		"application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf("неожиданный статус: %d", resp.StatusCode)
	}
	return nil
}
