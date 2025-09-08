package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("unexpected argument. usage: exit")
	}

	fmt.Println("Closing the Pokedex... Goodbye!")
	cfg.pokeapiClient.Close()
	os.Exit(0)
	return nil
}
