package pokeapi

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

func (c *Client) LocationPokemon(location string) ([]string, error) {
    url := baseURL + "/location-area/" + location
    
    dat,  ok := c.cache.Get(url)
    if !ok {
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            return nil, err
        }

        resp, err := c.httpClient.Do(req)
        if err != nil {
            return nil, err
        }
        defer resp.Body.Close()

        dat, err = io.ReadAll(resp.Body)
        if err != nil {
            return nil, err
        }

        if string(dat) == "Not Found" {
            return nil, fmt.Errorf("%s is not a valid location", location)
        }
    }
    
    locationDetailsResp := RespDetailsLocation{}
    err := json.Unmarshal(dat, &locationDetailsResp)
    if err != nil {
        return nil, err
    }

    pokemons := make([]string, len(locationDetailsResp.PokemonEncounters))
    for i, p := range locationDetailsResp.PokemonEncounters {
        pokemons[i] = p.Pokemon.Name
    }

    if !ok {
        c.cache.Add(url, dat)
    }

    return pokemons, nil
}
