package main

import "fmt"

func commandPokedex(c *config, _ ...string) error {
    fmt.Println("Your Pokedex:")
    for _, p := range c.pokedex {
        fmt.Printf(" - %s\n", p.Name)
    }
    return nil
}
