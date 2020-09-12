package cmd

import (
	"github.com/knadh/stuffbin"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// Hub represents the structure for all app wide functions and structs.
type Hub struct {
	Logger  *logrus.Logger
	Config  Config
	Fs      stuffbin.FileSystem
	Version string
}

// NewHub initializes an instance of Hub which holds app wide configuration.
func NewHub(logger *logrus.Logger, fs stuffbin.FileSystem, buildVersion string) *Hub {
	hub := &Hub{
		Logger:  logger,
		Fs:      fs,
		Version: buildVersion,
	}
	return hub
}

// initApp acts like a middleware to load app managers with Hub before running any command.
// Use this middleware to perform any action before the command is run.
func (hub *Hub) initApp(fn cli.ActionFunc) cli.ActionFunc {
	return func(c *cli.Context) error {
		// var err error
		// Initialize config.
		// hub.Config, err = initConfig(c)
		// if err != nil {
		// 	log.Fatalf("error while reading config: %v", err)
		// }
		return fn(c)
	}
}
