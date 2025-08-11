package main

import (
	"bufio"
	"fmt"
	"github.com/davidw1457/pokedexcli/internal/pokeapi"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		moreData := scanner.Scan()
		if moreData == false {
			fmt.Println(scanner.Err())
			os.Exit(1)
		}

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		} else if len(input) == 1 {
            input = append(input, "")
        }

		if command, ok := getCommands()[input[0]]; ok {
            err := command.callback(cfg, input[1:]...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	words := strings.ToLower(text)
	return strings.Fields(words)
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config, args ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "List next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous 20 locations",
			callback:    commandMapb,
		},
        "explore": {
            name: "explore <LOCATION>",
            description: "Show pokemon present in <LOCATION>",
            callback: commandExplore,
        },
        "catch": {
            name: "catch <POKEMON>",
            description: "Attempt to catch <POKEMON>",
            callback: commandCatch,
        },
        "inspect": {
            name: "inspect <POKEMON>",
            description: "Show details of <POKEMON> if caught",
            callback: commandInspect,
        },
        "pokedex": {
            name: "pokedex",
            description: "Show names of all captured pokemon",
            callback: commandPokedex,
        },
	}
}

type config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
    pokedex map[string]pokeapi.Pokemon
}
