package main

import (
	"io/fs"

	"github.com/sirupsen/logrus"
)

// Fs represents a Filesystem
type Fs struct {
	Fs           fs.FS
	TemplatePath string
}

// Hub represents the structure for all app wide functions and structs.
type Hub struct {
	Logger  *logrus.Logger
	Config  Config
	Fs      Fs
	Version string
}

// NewHub initializes an instance of Hub which holds app wide configuration.
func NewHub(logger *logrus.Logger, fs fs.FS, templatePath string, buildVersion string) *Hub {
	f := Fs{
		Fs:           fs,
		TemplatePath: templatePath,
	}
	hub := &Hub{
		Logger:  logger,
		Fs:      f,
		Version: buildVersion,
	}

	return hub
}
