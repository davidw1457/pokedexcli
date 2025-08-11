package pokeapi

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
    url := baseURL + "/pokemon/" + name
    
    dat, ok := c.cache.Get(url)
    if !ok {
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            return Pokemon{}, err
        }

        resp, err := c.httpClient.Do(req)
        if err != nil {
            return Pokemon{}, err
        }
        defer resp.Body.Close()

        dat, err = io.ReadAll(resp.Body)
        if err != nil {
            return Pokemon{}, err
        }

        if string(dat) == "Not Found" {
            return Pokemon{}, fmt.Errorf("%s is not a valid pokemon", name)
        }
    }

    pokemon := Pokemon{}
    err := json.Unmarshal(dat, &pokemon)
    if err != nil {
        return Pokemon{}, err
    }

    if !ok {
        c.cache.Add(url, dat)
    }

    return pokemon, nil
}
