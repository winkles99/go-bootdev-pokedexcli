package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/winkles99/go-bootdev-pokedexcli/pokedex"
)

func (c *Client) GetPokemon(url string) (pokedex.Pokemon, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokedex.Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return pokedex.Pokemon{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return pokedex.Pokemon{}, fmt.Errorf("Response failed with status code: %d and\n", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return pokedex.Pokemon{}, err
	}

	pokemon := pokedex.Pokemon{}
	err = json.Unmarshal(body, &pokemon)

	if err != nil {
		return pokedex.Pokemon{}, fmt.Errorf("bad response JSON")
	}
	return pokemon, nil
}