package main

import (
	"github.com/pelletier/go-toml"
)

func gatherDefaultConfig() ([]byte, error) {
	cfg := []byte(`
SiteName = "demo"
Title = "Hello World"`)

	config := Config{}
	err := toml.Unmarshal(cfg, &config)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
