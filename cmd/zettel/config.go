package main

func gatherDefaultConfig() (Config, error) {
	config := Config{
		SiteName: "demo",
		Title:    "Hello World",
	}
	return config, nil
}
