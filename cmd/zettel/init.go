package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml"
	"github.com/urfave/cli"
)

var (
	defaultconfigName    = "zettel.toml"
	defaultindexFileName = "index.md"
)

// InitProject initializes a new zettel site copies a sample config.
func (hub *Hub) InitProject(config Config) cli.Command {
	return cli.Command{
		Name:      "init",
		Aliases:   []string{"i"},
		Usage:     "Initializes a new zettel site with default config.",
		Action:    hub.init,
		ArgsUsage: "[SITENAME]",
		Before: func(c *cli.Context) error {
			if c.Args().First() == "" {
				return errors.New("Site name is missing")
			}
			return nil
		},
	}
}

func (hub *Hub) init(cliCtx *cli.Context) error {
	// get a default config
	siteDir := cliCtx.Args().First()
	//Create a folder/directory at a full qualified path
	err := os.Mkdir(siteDir, 0755)
	if err != nil {
		return err
	}
	cfg, err := gatherDefaultConfig()
	if err != nil {
		return fmt.Errorf("error creating default config file template: %v", err)
	}
	configFile, err := toml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("error unmarshalling config: %s", err)
	}
	// persist the default config file.
	err = createFile(configFile, filepath.Join(siteDir, defaultconfigName))
	if err != nil {
		return fmt.Errorf("error creating default config: %v", err)
	}
	// load a default `index.md`
	indexFile, err := hub.Fs.Read("templates/index.md")
	if err != nil {
		return fmt.Errorf("error reading default config file template: %v", err)
	}
	// persist the default index
	err = createFile(indexFile, filepath.Join(siteDir, defaultindexFileName))
	if err != nil {
		return fmt.Errorf("error creating default config: %v", err)
	}

	hub.Logger.Infof("Congrats! Your zettel site is created at %s", siteDir)
	return nil
}
