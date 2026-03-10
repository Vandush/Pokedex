package main

import (
	"fmt"
	"bufio"
	"os"
	"time"
	"github.com/Vandush/pokedexcli/config"
	"github.com/Vandush/pokedexcli/functions"
	"github.com/Vandush/pokedexcli/pokecache"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	cfg := &config.Config{
		NextUrl: nil,
		PrevUrl: nil,
	}
	
	interval := 20 * time.Second
	cache := pokecache.NewCache(interval)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := functions.CleanInput(input)
		if len(cleanedInput) == 0 {
			fmt.Println("No input")
		} else {
			commands := functions.GetCommands()
			key, ok := commands[cleanedInput[0]]
			if ok == false {
				fmt.Println("Unknown command")
			} else {
				cleanedInput = cleanedInput[1:]
				key.Callback(cfg, cache, cleanedInput)
			}
		}
	}
}
