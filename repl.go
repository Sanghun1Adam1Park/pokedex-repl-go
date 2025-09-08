package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Sanghun1Adam1Park/pokedex-repl-go/internal/pokeapi"
)

func cleanInput(text string) []string {
	strings := strings.Fields(strings.ToLower(text))

	return strings
}

type config struct {
	pokeapiClient    *pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func repl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	cliCommands := getCLICommands()

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}

		text := scanner.Text()
		cleanText := cleanInput(text)
		cmd, ok := cliCommands[cleanText[0]]

		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := cmd.callback(cfg, cleanText[1:]...); err != nil {
			fmt.Println(err)
			continue
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCLICommands() map[string]cliCommand {
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
			description: "Shows next 20 locations",
			callback:    commandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows previous 20 locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Shows list of pokemone at give location, usage: explore <location_name>",
			callback:    commandExplore,
		},
	}
}
