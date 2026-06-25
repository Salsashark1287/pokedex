package main

import ( 
	"strings"
	"fmt"
	"os"
)

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

func commandMap(*config) error {

}