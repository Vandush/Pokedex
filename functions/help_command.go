package functions

import (
	"fmt"
	"github.com/Vandush/pokedexcli/config"

)

func CommandHelp(c *config.Config) error {
	m := GetCommands()
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for _, value := range m {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}
