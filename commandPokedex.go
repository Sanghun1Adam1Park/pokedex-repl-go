package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("unepxected argument. usage: pokdex")
	}

	fmt.Println("Your Pokdex:")
	for pokemon, _ := range cfg.pokedex {
		fmt.Printf(" - %s\n", pokemon)
	}

	return nil
}
