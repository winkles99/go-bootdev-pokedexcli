package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type namedUrl struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationAreasResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []namedUrl  `json:"results"`
}

type LocationResp struct {
		PokemonEncounters []struct {
			Pokemon namedUrl `json:"pokemon"`
		} `json:"pokemon_encounters"`
}

func (c *Client) GetLocationAreas(url string) (LocationAreasResp, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return LocationAreasResp{}, fmt.Errorf("Response failed with status code: %d and\n", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	areas := LocationAreasResp{}
	err = json.Unmarshal(body, &areas)

	if err != nil {
		return LocationAreasResp{}, fmt.Errorf("bad response JSON")
	}
	return areas, nil
}

func (c *Client) GetLocationArea(url string) (LocationResp, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationResp{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationResp{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return LocationResp{}, fmt.Errorf("Response failed with status code: %d and\n", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationResp{}, err
	}

	area := LocationResp{}
	err = json.Unmarshal(body, &area)

	if err != nil {
		return LocationResp{}, fmt.Errorf("bad response JSON")
	}
	return area, nil
}