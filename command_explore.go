package main

import "fmt"

func commandExplore(c *config, args ...string) error {
    location := args[0]
    if location == "" {
        return fmt.Errorf("no location specified")
    }

    pokemons, err := c.pokeapiClient.LocationPokemon(location)
    if err != nil {
        return err
    }

    fmt.Printf("Exploring %s\n", location)
    fmt.Println("Found Pokemon:")
    for _, p := range pokemons {
        fmt.Printf(" - %s\n", p)
    }

    return nil
}
