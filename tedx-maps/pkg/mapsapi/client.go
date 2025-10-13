package mapsapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	apiKey string
	client *http.Client
}

func NewClient(apiKey string) Client {
	return Client{
		apiKey: apiKey,
		client: &http.Client{},
	}
}

func (c Client) FetchMapData(ctx context.Context, query string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json", query)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "tedx-maps-app")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data []map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("no results found")
	}

	return data[0], nil
}
