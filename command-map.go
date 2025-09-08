package main

import "fmt"

func commandMapF(cfg *config, args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("unexpected argument")
	}

	locations, err := cfg.pokeapiClient.GetShallowLocation(cfg.nextLocationsURL)
	if err != nil {
		return fmt.Errorf("error with getting shallow locaitons: %w", err)
	}

	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Prev

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapB(cfg *config, args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("unexpected argument")
	}

	locations, err := cfg.pokeapiClient.GetShallowLocation(cfg.prevLocationsURL)
	if err != nil {
		return fmt.Errorf("error with getting shallow locaitons: %w", err)
	}

	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Prev

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
