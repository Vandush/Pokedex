package functions

import (
	"fmt"
	"os"
	"github.com/Vandush/pokedexcli/config"
	"github.com/Vandush/pokedexcli/pokecache"
)

func CommandExit(_ *config.Config, _ *pokecache.Cache, _ []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}


