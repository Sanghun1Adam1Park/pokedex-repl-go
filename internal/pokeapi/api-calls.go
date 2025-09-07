package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

func (c *Client) GetShallowLocation(pageUrl *string) (ShallowLocation, error) {
	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ShallowLocation{}, fmt.Errorf("error with request: %w", err)
	}

	client := c.httpClient
	res, err := client.Do(req)
	if err != nil {
		return ShallowLocation{}, fmt.Errorf("error with response: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return ShallowLocation{}, fmt.Errorf("error unexpected status: %s", res.Status)
	}

	var shallowLocation ShallowLocation
	if err := json.NewDecoder(res.Body).Decode(&shallowLocation); err != nil {
		return ShallowLocation{}, fmt.Errorf("error decoding json: %w", err)
	}

	return shallowLocation, nil
}
