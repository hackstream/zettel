package main

// Config represents zettel site specific settings.
type Config struct {
	SiteName string `toml:"site_name"`
	Title    string `toml:"title"`
}
