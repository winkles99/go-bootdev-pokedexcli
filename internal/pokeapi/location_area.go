package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// Struct for a single location area
type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Struct for the API response
type LocationAreaResponse struct {
	Results  []LocationArea `json:"results"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
}

// GetLocationAreas fetches paginated location areas
func (c *Client) GetLocationAreas(pageURL string) (LocationAreaResponse, error) {
	if pageURL == "" {
		pageURL = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	}

	resp, err := c.httpClient.Get(pageURL)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("failed to GET location areas: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var locationResp LocationAreaResponse
	err = json.Unmarshal(body, &locationResp)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	// ✅ Fix: enforce HTTPS only after locationResp is populated
	locationResp.Next = enforceHTTPS(locationResp.Next)
	locationResp.Previous = enforceHTTPS(locationResp.Previous)

	return locationResp, nil
}

// Small helper function
func enforceHTTPS(url string) string {
	if strings.HasPrefix(url, "http://") {
		return strings.Replace(url, "http://", "https://", 1)
	}
	return url
}
