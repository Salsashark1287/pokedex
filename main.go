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
	if scanner.Scan() {
		input_text := scanner.Text()
		input_list := cleanInput(input_text)
		input_word := input_list[0]
		command_map := getCommands()
		if command, ok := command_map[input_word]; ok {
			command.callback(cfg)
		} else {
			fmt.Println("Unknown command")
		}
	}
}
}
