package main

import (
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var (
	// Version and date of the build. This is injected at build-time.
	buildVersion = "unknown"
	buildDate    = "unknown"
)

//go:embed templates
var fs embed.FS

// initLogger initializes logger
func initLogger(verbose bool) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Set logger level
	if verbose {
		logger.SetLevel(logrus.DebugLevel)
		logger.Debug("verbose logging enabled")
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}

	return logger
}

func main() {
	// Intialize new CLI app
	app := cli.NewApp()
	app.Name = "zettel"
	app.Usage = "Zettel builds a digital Zettelkasten website for your notes in Markdown."
	app.Version = fmt.Sprintf("%s, %s", buildVersion, buildDate)
	app.Authors = []*cli.Author{
		{
			Name: "Hackstream Devs",
		},
	}
	// Register command line args.
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:  "verbose",
			Usage: "Enable verbose logging",
		},
	}

	var logger = initLogger(true)

	if _, err := fs.Open("templates/index.tmpl"); err != nil {
		log.Fatalln(err)
	}

	// Initialize hub.
	hub := NewHub(logger, fs, buildVersion)

	// Register commands.
	app.Commands = []*cli.Command{
		hub.InitProject(),
		hub.NewPost(),
		hub.BuildSite(),
	}

	// Run the app.
	hub.Logger.Info("Starting zettel...🚀")

	if err := app.Run(os.Args); err != nil {
		logger.Errorf("OOPS: %s", err)
	}
}
