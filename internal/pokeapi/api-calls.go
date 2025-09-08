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

func (c *Client) GetShallowLocation(pageUrl *string) (ShallowLocation, error) {
	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	cacheData, exist := c.cache.Get(url)
	if exist {
		var shallowLocation ShallowLocation
		if err := json.Unmarshal(cacheData, &shallowLocation); err != nil {
			return ShallowLocation{}, fmt.Errorf("decode json: %w", err)
		}

		return shallowLocation, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ShallowLocation{}, fmt.Errorf("error with request: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ShallowLocation{}, fmt.Errorf("error with response: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return ShallowLocation{}, fmt.Errorf("error unexpected status: %s", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return ShallowLocation{}, fmt.Errorf("error reading body: %w", err)
	}

	c.cache.Add(url, data)

	var shallowLocation ShallowLocation
	if err := json.Unmarshal(data, &shallowLocation); err != nil {
		return ShallowLocation{}, fmt.Errorf("decode json: %w", err)
	}

	return shallowLocation, nil
}

func (c *Client) GetLocationArea(location string) (LocationArea, error) {
	url := baseURL + "/location-area/" + location

	cacheData, exist := c.cache.Get(url)
	if exist {
		var locationArea LocationArea
		if err := json.Unmarshal(cacheData, &locationArea); err != nil {
			return LocationArea{}, fmt.Errorf("error decoding cache data: %w", err)
		}
		return locationArea, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error constructing request: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error with response: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return LocationArea{}, fmt.Errorf("error unexpected status: %s", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error reading body: %w", err)
	}

	c.cache.Add(url, data)

	var locationArea LocationArea
	if err := json.Unmarshal(data, &locationArea); err != nil {
		return LocationArea{}, fmt.Errorf("decode json: %w", err)
	}

	return locationArea, nil
}
