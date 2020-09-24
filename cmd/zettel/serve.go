package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/cli/v2"
)

func (hub *Hub) StartServer() *cli.Command {
	return &cli.Command{
		Name:    "server",
		Aliases: []string{"s", "serve"},
		Usage:   "Starts a server with live reload",
		Action:  hub.MustHaveConfig(hub.startServer),
	}
}

func (hub *Hub) startServer(ctx *cli.Context) error {
	// Call build
	if err := hub.build(ctx); err != nil {
		return err
	}

	hub.Logger.Infof("Server started on :%d", DefaultPort)
	srv := http.FileServer(http.Dir(defaultDistDir))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", DefaultPort), srv); err != nil {
		return err
	}
	return nil
}
