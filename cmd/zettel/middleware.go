package main

import (
	"log"

	"github.com/urfave/cli"
)

// MustHaveConfig acts like a middleware to load config
// in Hub for the commands which require it.
func (hub *Hub) MustHaveConfig(fn cli.ActionFunc) cli.ActionFunc {
	return func(c *cli.Context) error {
		var err error
		// Initialize config.
		hub.Config, err = initConfig(c)
		if err != nil {
			log.Fatalf("error while reading config: %v", err)
		}
		return fn(c)
	}
}
