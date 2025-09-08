package pokeapi

type ShallowLocation struct {
	Count   int        `json:"count"`
	Next    *string    `json:"next"`
	Prev    *string    `json:"previous"`
	Results []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationArea struct {
	Name              string             `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
}

type DetailedPokemon struct {
	Name           string  `json:"name"`
	BaseExperience int     `json:"base_experience"`
	Height         int     `json:height`
	Weight         int     `json:weight`
	Stats          []Stats `json:"stats"`
	Types          []Types `json:"types"`
}

type Stats struct {
	BaseStat int  `json:"base_stat"`
	Stat     Stat `json:stat`
}

type Stat struct {
	Name string `json:"name"`
}

type Types struct {
	Type Type `json:"type"`
}

type Type struct {
	Name string `json:"name"`
}
