package main

import "fmt"

func commandHelp(cfg *config) error {
	cliCommands := getCLICommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	for _, cmd := range cliCommands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}
