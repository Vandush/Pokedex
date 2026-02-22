package main

import (
	"strings"
	"os"
	"fmt"
)

func CommandHelp() error {
	m := GetCommands()
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for _, value := range m {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}

func CommandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CleanInput(text string) []string {
	var list []string
	list = strings.Fields(text)
	for i, text := range list {
		list[i] = strings.ToLower(text)
	}
	return list
}
