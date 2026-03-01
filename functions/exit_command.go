package functions

import (
	"fmt"
	"os"
	"github.com/Vandush/pokedexcli/config"
)

func CommandExit(c *config.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}


