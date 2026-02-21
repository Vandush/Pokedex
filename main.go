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
			fmt.Println("Your command was: ")
		} else {
			fmt.Printf("Your command was: %s\n", cleaned[0])
		}
	}
}
