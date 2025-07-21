package repl

import (
	"fmt"
)

func commandPokedex(cfg *replConfig, opts []string) error {
	if len(cfg.Pokedex) == 0 {
		return fmt.Errorf("Pokemdex is empty. Catch a pokemon first")
	}

	fmt.Println("Your Pokedex:")

	for pokemonName := range cfg.Pokedex {
		fmt.Printf("  - %s\n", pokemonName)
	}

	return nil
}