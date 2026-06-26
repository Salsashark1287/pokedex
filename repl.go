package main

import ( 
	"strings"
	"fmt"
	"os"
	"net/http"
	"io"
	"encoding/json"
)
var cfg = &config{}

func cleanInput(text string) []string {
	lower_text := strings.ToLower(text)
	string_list := strings.Split(lower_text, " ")
	return string_list
}

func commandExit(*config) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(*config) error {
	command_map := getCommands()
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\n")
	for _, data := range command_map {
		fmt.Printf("%s: %s\n", data.name, data.description)
	} 
	return nil
}

func commandMap(cfg *config) error {

	url := "https://pokeapi.co/api/v2/location-area/"

	if cfg.Next != nil {
		url = *cfg.Next
	}
	
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var data ResponseStruct
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}
	cfg.Next = data.Next
	cfg.Previous = data.Previous

	for _, area := range data.Results {
		fmt.Println(area.Name)
	}
	

	return nil 
}

func commandMapb(cfg *config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
	} else {
		cfg.Next = cfg.Previous
		return commandMap(cfg)
	}
	return nil
}