package main

type cliCommand struct {
	name string
	description string
	callback func(*config) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name: "map",
			description: "List twenty locations.",
			callback: CommandMap,
		},
		"mapb": {
			name: "mapb",
			description: "List the previous locations.",
			callback: CommandMapBack,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: CommandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: CommandHelp,
		},
	}
}
