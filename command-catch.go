package main

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	baseExpMax = 635.0
)

func logDistributedInt(max int) int {
	u := rand.Float64()
	// transform uniform into log skew
	val := math.Log1p(u*float64(max)) / math.Log1p(float64(max))
	return int(val * float64(max))
}

func tryCatch(baseExp int) bool {
	draw := logDistributedInt(int(baseExpMax))
	return float64(draw) >= float64(baseExp)
}

func commandCactch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("unepxected argument. usage: catch <pokemon_name>")
	}

	pokemon := args[0]
	detailedPokemon, err := cfg.pokeapiClient.GetPokemonDetails(pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", detailedPokemon.Name)
	if catched := tryCatch(detailedPokemon.BaseExperience); catched {
		fmt.Printf("%s was caught!\n", detailedPokemon.Name)
		cfg.pokedex[pokemon] = detailedPokemon
	} else {
		fmt.Printf("%s escaped!\n", detailedPokemon.Name)
	}

	return nil
}
