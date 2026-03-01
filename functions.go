package main

import (
	"strings"
	"os"
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)

type PokedexAPI struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		Url string `json:"url"`
	} `json:"results"`
}

type config struct {
	NextUrl *string
	PrevUrl *string
}

func CommandMap(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if c.NextUrl != nil {
		url = string(*c.NextUrl)
	}
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Response: %v\n", err)
		return nil
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	locations := PokedexAPI{}
	if err := json.Unmarshal(data, &locations); err != nil {
		fmt.Printf("Unmarshal: %v\n", err)
		return nil
	}
	for _, location := range locations.Results {
		fmt.Printf("%s\n",location.Name)
	}
	c.NextUrl = locations.Next
	c.PrevUrl = locations.Previous
	return nil
}

func CommandMapBack(c *config) error {
	if c.PrevUrl == nil {
		fmt.Println("You're on the first page.")
		return nil
	}

	url := string(*c.PrevUrl)

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Response: %v\n", err)
		return nil
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	locations := PokedexAPI{}
	if err := json.Unmarshal(data, &locations); err != nil {
		fmt.Printf("Unmarshal: %v\n", err)
		return nil
	}
	for _, location := range locations.Results {
		fmt.Printf("%s\n",location.Name)
	}
	c.NextUrl = locations.Next
	c.PrevUrl = locations.Previous
	return nil
}

func CommandHelp(c *config) error {
	m := GetCommands()
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for _, value := range m {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}

func CommandExit(c *config) error {
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
