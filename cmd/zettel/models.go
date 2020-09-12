package main

// Config represents zettel site specific settings.
type Config struct {
	SiteName    string `toml:"site_name"`
	Description string `toml:"description"`
}
