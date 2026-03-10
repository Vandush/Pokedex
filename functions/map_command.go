package functions

import (
	"fmt"
	"github.com/Vandush/pokedexcli/config"
	"github.com/Vandush/pokedexcli/pokecache"
)

func CommandMap(cfg *config.Config, cache *pokecache.Cache, _ []string) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.NextUrl != nil {
		url = string(*cfg.NextUrl)
	}
	locations, err := config.CallEndpointAPI(url, cache)
	if err != nil {
		return err
	}
	for _, location := range locations.Results {
		fmt.Printf("%s\n",location.Name)
	}
	cfg.NextUrl = locations.Next
	cfg.PrevUrl = locations.Previous
	return nil
}

func CommandMapBack(cfg *config.Config, cache *pokecache.Cache, _ []string) error {
	if cfg.PrevUrl == nil {
		fmt.Println("You're on the first page.")
		return nil
	}
	url := string(*cfg.PrevUrl)
	locations, err := config.CallEndpointAPI(url, cache)
	if err != nil {
		return err
	}
	for _, location := range locations.Results {
		fmt.Printf("%s\n",location.Name)
	}
	cfg.NextUrl = locations.Next
	cfg.PrevUrl = locations.Previous
	return nil
}


