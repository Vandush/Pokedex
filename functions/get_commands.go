package functions

import (
	"github.com/Vandush/pokedexcli/config"
	"github.com/Vandush/pokedexcli/pokecache"
)

type cliCommand struct {
	name string
	description string
	Callback func(*config.Config, *pokecache.Cache, []string) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name: "map",
			description: "List twenty locations.",
			Callback: CommandMap,
		},
		"mapb": {
			name: "mapb",
			description: "List the previous locations.",
			Callback: CommandMapBack,
		},
		"explore": {
			name: "explore",
			description: "Discover the pokemon of a specific location.",
			Callback: CommandExplore,
		},
		"catch": {
			name: "catch",
			description: "Attempt to capture a pokemon!",
			Callback: CommandCatch,
		},
		"pokedex": {
			name: "pokedex",
			description: "List the pokemon you have caught!",
			Callback: ListPokemon,
		},
		"inspect": {
			name: "inspect",
			description: "Provides in-depth information of a particular pokemon you have caught.",
			Callback: InspectPokemon,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			Callback: CommandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			Callback: CommandHelp,
		},
	}
}
