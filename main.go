package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := CleanInput(input)
		if len(cleaned) == 0 {
			fmt.Println("No input")
		} else {
			commands := GetCommands()
			key, ok := commands[cleaned[0]]
			if ok == false {
				fmt.Println("Unknown command")
			} else {
				key.callback()
			}
		}
	}
}
