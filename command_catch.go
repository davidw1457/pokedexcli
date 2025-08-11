package main

import (
    "fmt"
    "math/rand"
    "time"
    "github.com/davidw1457/pokedexcli/internal/pokeapi"
)

func commandCatch(c* config, args ...string) error {
    name := args[0]
    if name == "" {
        return fmt.Errorf("No pokemon provided")
    }

    pokemon, err := c.pokeapiClient.GetPokemon(name)
    if err != nil {
        return err
    }

    fmt.Printf("Throwing a Pokeball at %s...\n", name)
    if isPokemonCaught(pokemon) {
        fmt.Printf("%s was caught!\n", name)
        c.pokedex[name] = pokemon
    } else {
        fmt.Printf("%s escaped!\n", name)
    }

    return nil
}

func isPokemonCaught(p pokeapi.Pokemon) bool {
    max := 615
    basexp := p.BaseExperience
    
    generator := rand.New(rand.NewSource(time.Now().Unix()))

    return generator.Intn(max) > basexp
}
