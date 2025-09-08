package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

func (c *Client) GetShallowLocation(pageURL *string) (ShallowLocation, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	return getResource[ShallowLocation](c, url)
}

func (c *Client) GetLocationArea(location string) (LocationArea, error) {
	url := baseURL + "/location-area/" + location
	return getResource[LocationArea](c, url)
}

func (c *Client) GetPokemonDetails(pokemonName string) (DetailedPokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName
	return getResource[DetailedPokemon](c, url)
}

func getResource[T any](c *Client, url string) (T, error) {
	var zero T

	cacheData, exist := c.cache.Get(url)
	if exist {
		var t T
		if err := json.Unmarshal(cacheData, &t); err != nil {
			return zero, fmt.Errorf("error decoding cache data: %w", err)
		}
		return t, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return zero, fmt.Errorf("error constructing request: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return zero, fmt.Errorf("error with response: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return zero, fmt.Errorf("error unexpected status: %s", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return zero, fmt.Errorf("error reading body: %w", err)
	}

	c.cache.Add(url, data)

	var t T
	if err := json.Unmarshal(data, &t); err != nil {
		return zero, fmt.Errorf("decode json: %w", err)
	}

	return t, nil
}
