package main

import (
	"fmt"
	"bufio"
	"os"
	"github.com/Vandush/pokedexcli/config"
	"github.com/Vandush/pokedexcli/functions"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	cfg := &config.Config{
		NextUrl: nil,
		PrevUrl: nil,
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := functions.CleanInput(input)
		if len(cleaned) == 0 {
			fmt.Println("No input")
		} else {
			commands := functions.GetCommands()
			key, ok := commands[cleaned[0]]
			if ok == false {
				fmt.Println("Unknown command")
			} else {
				key.Callback(cfg)
			}
		}
	}
}
