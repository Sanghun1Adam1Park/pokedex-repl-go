package main

import (
	"time"

	"github.com/Sanghun1Adam1Park/pokedex-repl-go/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	repl(cfg)
}
