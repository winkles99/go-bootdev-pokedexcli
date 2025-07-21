package repl

import (
	"fmt"
	"math/rand"

	"github.com/winkles99/go-bootdev-pokedexcli/pokedex"
)

func commandCatch(cfg *replConfig, opts []string) error {
	var pokemon pokedex.Pokemon
	var err error
	pokemonName := opts[0]

	if pokemonName == "" {
		return fmt.Errorf("Pokemon name name cannot be empty")
	}

	if _, ok := cfg.Pokedex[pokemonName]; ok {
		return fmt.Errorf("Pokemon %s already caught!", pokemonName)
	}

	url := cfg.PokeapiClient.BuildUrl("pokemon", pokemonName)

	if cache, ok := cfg.Cache.Get(url); ok {
		pokemon = cache.(pokedex.Pokemon)
	} else {
		pokemon, err = cfg.PokeapiClient.GetPokemon(url)
		if err != nil {
			return err
		}
		cfg.Cache.Add(url, pokemon)
	}

	fmt.Println("Throwing a Pokeball at pikachu...")
	chance := rand.Intn(30)
	if chance < 10 {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}

	cfg.Pokedex[pokemonName] = pokemon
	fmt.Println("pikachu was caught!")
	fmt.Println("You may now inspect it with the inspect command.")
	return nil
}