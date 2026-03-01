package config

type PokedexAPI struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		Url string `json:"url"`
	} `json:"results"`
}

type Config struct {
	NextUrl *string
	PrevUrl *string
}


