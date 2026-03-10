package functions

import (
	"fmt"
	"github.com/Vandush/pokedexcli/config"
	"github.com/Vandush/pokedexcli/pokecache"
)

func CommandExplore(cfg *config.Config, cache *pokecache.Cache, input []string) error {
	if len(input) == 0 {
		fmt.Println("No location given. Use 'map/mapb' to list locations")
		return nil
	} 
	if len(input) > 1 {
		fmt.Println("Locations are a continuous string.")
		fmt.Println("Example: 'explore canalave-city-area'")
		return nil
	}
	url := "https://pokeapi.co/api/v2/location-area/" + input[0]
	res, err := config.CallLocationAPI(url, cache)
	if err != nil {
		return err
	}

	for _, encounter := range res.PokemonEncounters {
		fmt.Printf("%s\n", encounter.Pokemon.Name)
	}
	return nil
}
