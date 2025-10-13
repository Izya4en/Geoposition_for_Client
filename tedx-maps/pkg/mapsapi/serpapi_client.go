package mapsapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SerpAPIClient struct {
	apiKey string
}

func NewSerpAPIClient(apiKey string) *SerpAPIClient {
	return &SerpAPIClient{apiKey: apiKey}
}

func (c *SerpAPIClient) SearchPlace(ctx context.Context, query string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://serpapi.com/search.json?q=%s&api_key=%s", query, c.apiKey)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var data map[string]interface{}
	json.Unmarshal(body, &data)
	return data, nil
}
