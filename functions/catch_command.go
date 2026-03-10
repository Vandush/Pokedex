package functions

import (
	"fmt"
	"math/rand/v2"
	"encoding/json"
	"github.com/Vandush/pokedexcli/config"
	"github.com/Vandush/pokedexcli/pokecache"
)

func CommandCatch(_ *config.Config, cache *pokecache.Cache, input []string) error {
	if len(input) == 0 {
		fmt.Println("No pokemon was specified.")
		return nil
	} 
	if len(input) > 1 {
		fmt.Println("Too many inputs.")
		fmt.Println("Example: 'catch vaporeon'")
		return nil
	}

	url := "https://pokeapi.co/api/v2/pokemon/" + input[0]
	res, err := config.CallPokemonAPI(url, cache)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", res.Name)

	roll := rand.IntN(res.BaseExperience)

	// Like WoW death rolling.
	requirement := 100
	if roll <= requirement {
		data, err := json.Marshal(res)
		if err != nil {
			return err
		}
		cache.CatchPokemon(res.Name, data)
		fmt.Printf("%s was caught!\n", res.Name)
	} else {
		fmt.Printf("%s escaped!\n", res.Name)
	}

	return nil
}
