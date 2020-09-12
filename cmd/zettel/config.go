package main

import (
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/urfave/cli/v2"
)

// initConfig initializes the app's configuration manager.
func initConfig(c *cli.Context) (Config, error) {
	var cfg = Config{}
	var ko = koanf.New(".")

	log.Printf("reading config: %s", defaultconfigName)
	if err := ko.Load(file.Provider(defaultconfigName), toml.Parser()); err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	// Read the configuration and load it to internal struct.
	err := ko.Unmarshal("", &cfg)
	return cfg, err
}

func gatherDefaultConfig() (Config, error) {
	config := Config{
		SiteName: "demo",
		Title:    "Hello World",
	}
	return config, nil
}
