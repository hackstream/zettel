package main

import (
	"embed"

	"github.com/sirupsen/logrus"
)

// Hub represents the structure for all app wide functions and structs.
type Hub struct {
	Logger  *logrus.Logger
	Config  Config
	Fs      embed.FS
	Version string
}

// NewHub initializes an instance of Hub which holds app wide configuration.
func NewHub(logger *logrus.Logger, fs embed.FS, buildVersion string) *Hub {
	hub := &Hub{
		Logger:  logger,
		Fs:      fs,
		Version: buildVersion,
	}

	return hub
}
