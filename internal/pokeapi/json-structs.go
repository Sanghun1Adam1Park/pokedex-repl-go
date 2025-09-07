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
