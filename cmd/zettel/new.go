package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/urfave/cli"
)

var (
	maxTitleLength = 20
	defaultPostDir = "content"
)

// NewPost initializes git repo and copies a sample config
func (hub *Hub) NewPost(config Config) cli.Command {
	return cli.Command{
		Name:      "new",
		Aliases:   []string{"n"},
		Usage:     "Create a new post.",
		Action:    hub.newPost,
		ArgsUsage: "[TITLE]",
		Before: func(c *cli.Context) error {
			if c.Args().First() == "" {
				return errors.New("Title is missing")
			}
			return nil
		},
	}
}

func (hub *Hub) newPost(cliCtx *cli.Context) error {
	title := cliCtx.Args().First()
	// fill basic metadata
	var cfg = map[string]interface{}{
		"Date":  time.Now().Format("2006-01-02T15:04:05"),
		"Title": title,
	}
	// clean up title
	sanitizedTitle := strings.ToLower(title)
	sanitizedTitle = strings.ReplaceAll(sanitizedTitle, " ", "-")
	if len(sanitizedTitle) > maxTitleLength {
		sanitizedTitle = sanitizedTitle[:maxTitleLength]
	}
	sanitizedTitle = sanitizedTitle + ".md"

	// persist the new post
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	contentDir := filepath.Join(currentDir, defaultPostDir)
	// create content dir if it doesn't exit
	if _, err := os.Stat(contentDir); os.IsNotExist(err) {
		os.Mkdir(contentDir, 0755)
	}
	if err != nil {
		return err
	}
	// persist Post file.
	path := filepath.Join(contentDir, sanitizedTitle)
	post, err := os.Create(path)
	if err != nil {
		return err
	}
	// render post template
	err = saveResource("templates/post.tmpl", post, cfg, hub.Fs)
	if err != nil {
		return err
	}

	hub.Logger.Infof("New post created! %s", path)
	return nil
}
