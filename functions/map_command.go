package functions

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
	"github.com/Vandush/pokedexcli/config"

)

func CommandMap(c *config.Config) error {
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
	locations := config.PokedexAPI{}
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

func CommandMapBack(c *config.Config) error {
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
	locations := config.PokedexAPI{}
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


