package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	if len(args) > 0 {
		return fmt.Errorf("unexpected argument. usage: help")
	}

	cliCommands := getCLICommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range cliCommands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}
