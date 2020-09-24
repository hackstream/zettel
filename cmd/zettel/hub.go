package main

import (
	"github.com/fsnotify/fsnotify"
	"github.com/knadh/stuffbin"
	"github.com/sirupsen/logrus"
)

// Hub represents the structure for all app wide functions and structs.
type Hub struct {
	Logger  *logrus.Logger
	Config  Config
	Fs      stuffbin.FileSystem
	Version string
	Watcher *fsnotify.Watcher
}

type hubCfg struct {
	logger       *logrus.Logger
	fs           stuffbin.FileSystem
	buildVersion string
	watcher      *fsnotify.Watcher
}

// NewHub initializes an instance of Hub which holds app wide configuration.
func NewHub(cfg hubCfg) *Hub {
	hub := &Hub{
		Logger:  cfg.logger,
		Fs:      cfg.fs,
		Version: cfg.buildVersion,
		Watcher: cfg.watcher,
	}

	return hub
}
