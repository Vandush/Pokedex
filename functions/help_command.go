package functions

import (
	"fmt"
	"slices"
	"github.com/Vandush/pokedexcli/config"
	"github.com/Vandush/pokedexcli/pokecache"
)

func CommandHelp(_ *config.Config, _ *pokecache.Cache, _ []string) error {
	m := GetCommands()
	var order []string
	for key, _ := range m {
		order = append(order, key)
	}
	slices.Sort(order)
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for i := 0; i < len(order); i++ {
		fmt.Printf("%s: %s\n", m[order[i]].name, m[order[i]].description)
	}
//	for _, value := range m {
//		fmt.Printf("%s: %s\n", value.name, value.description)
//	}
	return nil
}
