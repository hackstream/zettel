package main

import (
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
)

// Config represents zettel site specific settings.
type Config struct {
	SiteName      string `toml:"site_name" koanf:"site_name"`
	Description   string `toml:"description" koanf:"description"`
	Pygmentsstyle string `toml:"pygmentsstyle" koanf:"pygmentsstyle"`
	SitePrefix    string `toml:"site_prefix,omitempty" koanf:"site_prefix"`
	StripHTML     bool   `toml:"strip_html" koanf:"strip_html"`
}

// initConfig initializes the app's configuration manager.
func initConfig() (Config, error) {
	var (
		cfg = Config{}
		ko  = koanf.New(".")
	)

	log.Printf("reading config: %s", defaultconfigName)

	if err := ko.Load(file.Provider(defaultconfigName), toml.Parser()); err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// Read the configuration and load it to internal struct.
	err := ko.Unmarshal("", &cfg)

	return cfg, err
}

func gatherDefaultConfig() Config {
	config := Config{
		SiteName:      "My Zettel",
		Description:   "Hello World. This is my zettel notebook",
		Pygmentsstyle: "monokailight",
	}

	return config
}
