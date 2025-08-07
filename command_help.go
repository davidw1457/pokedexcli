package main

import "fmt"

func commandHelp() error {
    fmt.Print(`
Welcome to the Pokedex!
Usage:

`)
    for _, cmd := range getCommands() {
        fmt.Printf("%s: %s\n", cmd.name, cmd.description)
    }
    fmt.Println()
    return nil
}
