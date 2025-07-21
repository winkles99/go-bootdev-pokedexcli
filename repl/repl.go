package repl

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/winkles99/go-bootdev-pokedexcli/internal/pokeapi"
	"github.com/winkles99/go-bootdev-pokedexcli/internal/pokecache"
	"github.com/winkles99/go-bootdev-pokedexcli/pokedex"
	"github.com/winkles99/go-bootdev-pokedexcli/utils"
)

type replConfig struct {
	PokeapiClient  pokeapi.Client
	Cache          pokecache.Cache[any]
	Pokedex        map[string]pokedex.Pokemon
	NextOffset     int
	PreviousOffset int
	Limit          int
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *replConfig, opts []string) error
}

func StartRepl() {
	cfg := replConfig{pokeapi.NewClient(), pokecache.NewCache(5 * time.Minute), make(map[string]pokedex.Pokemon), 0, 0, 20}
	scanner := bufio.NewScanner(os.Stdin)

	// handle error
	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}
	for {
		fmt.Print("Pokedex > ")
		// reads user input until \n by default
		scanner.Scan()
		// Holds the string that was scanned
		prompt := scanner.Text()

		cmd, err := utils.GetCmdFromPrompt(prompt)
		if err != nil {
			continue
		}
		command, ok := getCommandsMap()[cmd[0]]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}
		err = command.callback(&cfg, cmd[1:])
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
	}
}

func getCommandsMap() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Each subsequent call displays the names of next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <areaName>",
			description: "Displays a list of Pokémon that can be encountered in this area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemonName>",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemonName>",
			description: "Show an already caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show a pokedex contents",
			callback:    commandPokedex,
		},
	}
}