package repl

import (
	"fmt"
	"strings"
)

func commandInspect(cfg *replConfig, opts []string) error {
	pokemonName := opts[0]

	if pokemonName == "" {
		return fmt.Errorf("Pokemon name name cannot be empty")
	}

	p, ok := cfg.Pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	var stats, types strings.Builder
	for _, stat := range p.Stats {
		fmt.Fprintf(&stats, "  - %s: %2d\n", stat.Stat.Name, stat.BaseStat)
	}
	for _, typ := range p.Types {
		fmt.Fprintf(&types, "  - %s\n", typ.Type.Name)
	}
	fmt.Printf("Name: %s\nHeight: %2d\nWeight: %2d\nStats:\n%sTypes:\n%s", p.Name, p.Height, p.Weight, stats.String(), types.String())

	return nil
}