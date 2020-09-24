package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/knadh/stuffbin"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var (
	// Version and date of the build. This is injected at build-time.
	buildVersion = "unknown"
	buildDate    = "unknown"
)

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

// initFileSystem initializes the stuffbin FileSystem to provide
// access to bunded static assets to the app.
func initFileSystem() (stuffbin.FileSystem, error) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exPath := filepath.Dir(ex)
	fs, err := stuffbin.UnStuff(filepath.Join(exPath, filepath.Base(os.Args[0])))

	if err != nil {
		return nil, err
	}

	return fs, nil
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

	// Initialize the static file system into which all
	// required static assets (.css, .js files etc.) are loaded.
	fs, err := initFileSystem()
	if err != nil {
		logger.Errorf("error reading stuffed binary: %s", err)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logger.Fatalln(err)
	}
	defer watcher.Close()

	hubcfg := hubCfg{
		logger:       logger,
		buildVersion: buildVersion,
		fs:           fs,
		watcher:      watcher,
	}

	// Initialize hub.
	hub := NewHub(hubcfg)

	// Register commands.
	app.Commands = []*cli.Command{
		hub.InitProject(),
		hub.NewPost(),
		hub.BuildSite(),
		hub.StartServer(),
	}

	// Run the app.
	hub.Logger.Info("Starting zettel...🚀")

	if err = app.Run(os.Args); err != nil {
		logger.Errorf("OOPS: %s", err)
	}
}
