package main

import (
	"fmt"
)

func commandMap(c *config) error {
	if c.Next == nil && c.Previous != nil {
		return fmt.Errorf("you're on the last page")
	}
	return Map(c, c.Next, "last")
}

func commandMapb(c *config) error {
	if c.Previous == nil {
		return fmt.Errorf("you're on the first page")
	}
	return Map(c, c.Previous, "first")
}

func Map(c *config, url *string, end string) error {
	res, err := c.pokeapiClient.ListLocations(url)
	if err != nil {
		return err
	}

	c.Next = res.Next
	c.Previous = res.Previous

	for _, l := range res.Results {
		fmt.Println(l.Name)
	}

	return nil
}
