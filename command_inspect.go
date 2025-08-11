package main

import "fmt"

func commandInspect(c *config, args ...string) error {
    name := args[0]
    if name == "" {
        return fmt.Errorf("No pokemon specified")
    }

    pokemon, ok := c.pokedex[name]
    if !ok {
        fmt.Println("you have not caught that pokemon")
        return nil
    }

    fmt.Printf("Name: %s\n", name)
    fmt.Printf("Height: %d\n", pokemon.Height)
    fmt.Printf("Weight: %d\n", pokemon.Weight)
    fmt.Println("Stats:")
    for _, s := range pokemon.Stats {
        fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
    }
    fmt.Println("Types:")
    for _, t := range pokemon.Types {
        fmt.Printf("  - %s\n", t.Type.Name)
    }
    return nil
}
