package main

type cliCommand struct {
	name string
	description string
	callback func(*config) error
}
type config struct {
	Next *string
	Previous *string
}
type ResponseStruct struct {
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"results"`
}
func getCommands() map[string]cliCommand{
	return map[string]cliCommand{
	"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
		},
	"help": {
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
		},
	"map": {
		name: "map",
		description: "Displays next area locations",
		callback: commandMap,
	},
	"mapb": {
		name: "mapb",
		description: "Displays previous area locations",
		callback: commandMapb,
	},
	}
}