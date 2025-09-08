package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("unexpected argument. usage: explore <location_name>")
	}

	location := args[0]
	locationArea, err := cfg.pokeapiClient.GetLocationArea(location)
	if err != nil {
		return fmt.Errorf("error making api calls: %w", err)
	}

	fmt.Printf("Exploring %s\n", locationArea.Name)
	fmt.Println("Found Pokemon:")
	for _, pokemonEncounter := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemonEncounter.Pokemon.Name)
	}

	return nil
}
