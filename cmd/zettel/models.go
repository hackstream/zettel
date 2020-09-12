package main

import "time"

// Config represents zettel site specific settings.
type Config struct {
	SiteName string `toml:"site_name"`
	Title    string `toml:"title"`
}

type Post struct {
	Body     string
	Metadata Metadata
	// links    []Link
}

type Metadata struct {
	Date  time.Time
	Title string
}
