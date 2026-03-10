package functions

import (
	"fmt"
	"encoding/json"
	"github.com/Vandush/pokedexcli/config"
	"github.com/Vandush/pokedexcli/pokecache"
)

func ListPokemon(_ *config.Config, cache *pokecache.Cache, _ []string) error {
	fmt.Println("Your Pokedex:")
	for key, _ := range cache.PokemonCaught {
		fmt.Printf(" - %s\n", key)
	}
	return nil
}

func InspectPokemon(_ *config.Config, cache *pokecache.Cache, input []string) error {
	if len(input) == 0 {
		fmt.Println("No pokemon was specified.")
		return nil
	} 
	if len(input) > 1 {
		fmt.Println("Too many inputs.")
		fmt.Println("Example: 'inspect vaporeon'")
		return nil
	}

	results := config.PokemonAPI{}
	data, exists := cache.GetPokemon(input[0])
	if exists {
		if err := json.Unmarshal(data, &results); err != nil {
			fmt.Printf("Unmarshal: %v\n", err)
		}
		fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n", results.Name, results.Height, results.Weight)
		for _, stat := range results.Stats {
			fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, poketype := range results.Types {
			fmt.Printf(" - %s\n", poketype.Type.Name)
		}
	} else {
		fmt.Printf("You have not caught a %s\n", input[0])
	}

	return nil
}
