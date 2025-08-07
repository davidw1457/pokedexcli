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
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		fmt.Printf("Your command was: %s\n", cleanedInput[0])
	}
}

func cleanInput(text string) []string {
    words := strings.ToLower(text)
    return strings.Fields(words)
}
