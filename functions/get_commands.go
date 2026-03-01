package functions

import (
	"github.com/Vandush/pokedexcli/config"
)

type cliCommand struct {
	name string
	description string
	Callback func(*config.Config) error
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
