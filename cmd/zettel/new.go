package main

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

// NewPost initializes git repo and copies a sample config
func (hub *Hub) NewPost() *cli.Command {
	return &cli.Command{
		Name:      "new",
		Aliases:   []string{"n"},
		Usage:     "Create a new post.",
		Action:    hub.MustHaveConfig(hub.newPost),
		ArgsUsage: "[TITLE]",
		Before: func(c *cli.Context) error {
			if c.Args().First() == "" {
				return errors.New("title is missing")
			}

			return nil
		},
	}
}

func (hub *Hub) newPost(cliCtx *cli.Context) error {
	title := cliCtx.Args().First()
	// fill basic metadata
	var cfg = map[string]interface{}{
		"Date":  time.Now().Format(time.RFC3339),
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

	p := filepath.Join(currentDir, defaultPostDir, sanitizedTitle)

	post, err := os.Create(p)
	if err != nil {
		return err
	}

	// render post template
	err = saveResource("index", []string{path.Join(hub.Fs.TemplatePath, "post.tmpl")}, post, cfg, hub.Fs.Fs)
	if err != nil {
		return err
	}

	hub.Logger.Infof("New post created! %s", p)

	return nil
}
