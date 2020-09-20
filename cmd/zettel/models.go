package main

// Config represents zettel site specific settings.
type Config struct {
	SiteName      string `toml:"site_name" koanf:"site_name"`
	Description   string `toml:"description" koanf:"description"`
	Pygmentsstyle string `toml:"pygmentsstyle" koanf:"pygmentsstyle"`
	SitePrefix    string `toml:"site_prefix,omitempty" koanf:"site_prefix"`
}
