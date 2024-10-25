package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/oavdonin/pokedex-cli/api"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commandSet = map[string]cliCommand{
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "List locations",
		callback:    pokedexMap,
	},
	"mapb": {
		name:        "mapb",
		description: "",
		callback:    pokedexMapb,
	},
	"exit": {
		name:        "exit",
		description: "exit",
		callback:    func() error { return nil },
	},
}

func printPrompt() error {
	fmt.Print("pokedex > ")
	return nil
}

func commandVersion() error {
	fmt.Println("version 0.1")
	return nil
}

func commandHelp() error {
	fmt.Println("available commands: [version, help, exit]")
	return nil
}

func pokedexMap() error {
	api.InitializeLocations()
	err := api.GetLocations()
	if err != nil {
		return err
	}
	return nil
}

func pokedexMapb() error {
	api.InitializeLocations()
	err := api.GetLocations("down")
	if err != nil {
		return err
	}
	return nil
}

func main() {

	reader := bufio.NewScanner(os.Stdin)
	printPrompt()
replloop:
	for reader.Scan() {
		command := reader.Text()
		switch commandSet[command].name {
		case "exit":
			break replloop
		case "help":
			commandHelp()
		case "map":
			pokedexMap()
		case "mapb":
			pokedexMapb()
		default:
			commandHelp()
		}
		printPrompt()
	}
}
