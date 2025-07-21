package pokeapi

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const endpoint = "https://pokeapi.co/api/v2/"

type Client struct {
	httpClient http.Client
}

func NewClient() Client {
	return Client{httpClient: http.Client{Timeout: time.Minute}}
}

func (c *Client) BuildUrl(paths ...string) string {
	path := strings.Join(paths, "/")
	return endpoint + path
}

func (c *Client) BuildListUrl(path string, offset, limit int) string {
	u, _ := url.Parse(endpoint + path)
	query := url.Values{}
	if limit != 0 {
		query.Set("limit", strconv.Itoa(limit))
	}
	if offset != 0 {
		query.Set("offset", strconv.Itoa(offset))
	}

	u.RawQuery = query.Encode()

	return u.String()
}