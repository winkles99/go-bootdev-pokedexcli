package repl

import (
	"fmt"

	"github.com/winkles99/go-bootdev-pokedexcli/internal/pokeapi"
	"github.com/winkles99/go-bootdev-pokedexcli/utils"
)

type area struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func getAreas(cfg *replConfig, offset int) error {
	var areas pokeapi.LocationAreasResp
	var err error
	url := cfg.PokeapiClient.BuildListUrl("location-area", offset, cfg.Limit)

	if cache, ok := cfg.Cache.Get(url); ok {
		areas = cache.(pokeapi.LocationAreasResp)
	} else {
		areas, err = cfg.PokeapiClient.GetLocationAreas(url)
		if err != nil {
			return err
		}
		cfg.Cache.Add(url, areas)
	}

	cfg.PreviousOffset = utils.GetOffsetFromUrl(areas.Previous)
	cfg.NextOffset = utils.GetOffsetFromUrl(areas.Next)

	fmt.Println("Location areas:")
	for _, area := range areas.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	return nil
}

func commandMap(cfg *replConfig, opts []string) error {
	err := getAreas(cfg, cfg.NextOffset)
	if err != nil {
		return err
	}
	return nil
}

func commandMapb(cfg *replConfig, opts []string) error {
	if cfg.PreviousOffset == 0 && cfg.NextOffset == cfg.Limit {
		return fmt.Errorf("no previous page")
	}
	err := getAreas(cfg, cfg.PreviousOffset)
	if err != nil {
		return err
	}
	return nil
}