package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

var (
	defaultconfigName    = "zettel.toml"
	defaultindexFileName = "index.md"
)

// InitProject initializes git repo and copies a sample config
func (hub *Hub) InitProject(config Config) cli.Command {
	return cli.Command{
		Name:    "init",
		Aliases: []string{"i"},
		Usage:   "Initializes an empty git repo with a kubekutr config file.",
		Action:  hub.init,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "default, d",
				Usage: "Use the default config file",
			},
			cli.StringFlag{
				Name:  "output, o",
				Usage: "Config file name.",
			},
		},
	}
}

func (hub *Hub) init(cliCtx *cli.Context) error {
	// get a default config
	configFile, err := gatherDefaultConfig()
	if err != nil {
		return fmt.Errorf("error creating default config file template: %v", err)
	}
	// persist the default config file.
	err = createFile(configFile, defaultconfigName)
	if err != nil {
		return fmt.Errorf("error creating default config: %v", err)
	}
	// load a default `index.md`
	indexFile, err := hub.Fs.Read("templates/index.md")
	if err != nil {
		return fmt.Errorf("error reading default config file template: %v", err)
	}
	err = createFile(indexFile, defaultindexFileName)
	if err != nil {
		return fmt.Errorf("error creating default config: %v", err)
	}

	hub.Logger.Infof("Congrats! Your zettel site is created at %s", defaultconfigName)
	return nil
}
