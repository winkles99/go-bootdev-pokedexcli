package repl

import (
	"fmt"

	"github.com/winkles99/go-bootdev-pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *replConfig, opts []string) error {
	var area pokeapi.LocationResp
	var err error
	areaName := opts[0]

	if areaName == "" {
		return fmt.Errorf("location name cannot be empty")
	}

	url := cfg.PokeapiClient.BuildUrl("location-area", areaName)

	if cache, ok := cfg.Cache.Get(url); ok {
		area = cache.(pokeapi.LocationResp)
	} else {
		area, err = cfg.PokeapiClient.GetLocationArea(url)
		if err != nil {
			return err
		}
		cfg.Cache.Add(url, area)
	}

	fmt.Printf("Pokemons in %s:\n", areaName)
	for _, encounter := range area.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}