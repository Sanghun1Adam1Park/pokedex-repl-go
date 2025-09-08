package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("unexpected argument. usage: inspect <pokemon_name>")
	}

	pokemon := args[0]
	if detailedPokemon, caught := cfg.pokedex[pokemon]; caught {
		fmt.Printf("Name: %s\n", detailedPokemon.Name)
		fmt.Printf("Height: %d\n", detailedPokemon.Height)
		fmt.Printf("Weight: %d\n", detailedPokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range detailedPokemon.Stats {
			fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range detailedPokemon.Types {
			fmt.Printf(" -%s\n", t.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}
