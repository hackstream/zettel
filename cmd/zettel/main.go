package main

import (
	cli "github.com/urfave/cli/v2" // imports as package "cli"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "creates a zettel project",
				Action:  initProject,
			},
		},
		Name:        "zettel",
		Description: "A notes organizer",
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
