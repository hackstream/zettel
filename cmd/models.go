package cmd

import "time"

// Config represents zettel site specific settings.
type Config struct {
	SiteName string
	Title    string
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
