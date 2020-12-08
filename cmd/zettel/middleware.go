package main

import (
	"log"

	"github.com/urfave/cli/v2"
)

// MustHaveConfig acts like a middleware to load config
// in Hub for the commands which require it.
func (hub *Hub) MustHaveConfig(fn cli.ActionFunc) cli.ActionFunc {
	return func(c *cli.Context) error {
		var err error

		// Initialize config.
		hub.Config, err = initConfig()
		if err != nil {
			log.Fatalf("error while reading config: %v", err)
		}

		return fn(c)
	}
}

func (hub *Hub) MustInitFileSystem(fn cli.ActionFunc) cli.ActionFunc {
	return func(c *cli.Context) error {
		fs := c.String("filesystem")
		if fs == "builtin" {
			return fn(c)
		}

		disk, err := initDiskFileSystem(fs)
		if err != nil {
			log.Fatalf("error while initializing disk filesystem %q: %v", fs, err)
		}

		hub.Fs = disk

		return fn(c)
	}
}
