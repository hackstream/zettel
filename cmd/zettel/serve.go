package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/urfave/cli/v2"
)

var upgrader = websocket.Upgrader{} // use default options

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
	r := mux.NewRouter()
	r.HandleFunc("/ws", hub.liveReload(ctx))
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(defaultDistDir))))
	hub.Logger.Infof("Server started on :%d", DefaultPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", DefaultPort), r); err != nil {
		return err
	}
	return nil
}

func (hub *Hub) liveReload(ctx *cli.Context) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			hub.Logger.Fatalf("upgrade: %v", err)
			return
		}
		defer c.Close()
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			hub.Logger.Fatalln(err)
		}
		// Listen on events and build whenever a file has added or changed in content dir.
		// Since fsnotify's watcher doesn't support recursively watching the directory,
		// we recursively add the dirs to the watcher
		filepath.Walk(defaultPostDir, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				if err := watcher.Add(path); err != nil {
					hub.Logger.Fatalf("error watching: %v", err)
				}
			}
			return nil
		})
		for {
			select {
			case <-watcher.Events:
				hub.Logger.Info("files modified or added, rebuilding...")
				if err := hub.build(ctx); err != nil {
					hub.Logger.Errorf("error building site: %v", err)
					continue
				}
				watcher.Close()
				// Send a reload
				c.WriteMessage(websocket.TextMessage, []byte("reload"))
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				hub.Logger.Infoln("error:", err)
			}
		}
	}
}
