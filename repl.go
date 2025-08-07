package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		moreData := scanner.Scan()
		if moreData == false {
			fmt.Println(scanner.Err())
			os.Exit(1)
		}
		input := cleanInput(scanner.Text())
        
        if command, ok := getCommands()[input[0]]; ok {
            err := command.callback()
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
    name string
    description string
    callback func() error
}

func getCommands() map[string]cliCommand {
    return map[string]cliCommand{
        "help": {
            name: "help",
            description: "Displays a help message",
            callback: commandHelp,
        },
        "exit": {
            name: "exit",
            description: "Exit the Pokedex",
            callback: commandExit,
        },
    }
}
