package main

import (
	"path"

	"github.com/hackstream/zettel/internal/pipeline"
	"github.com/urfave/cli/v2"
)

// BuildSite initializes a new zettel site copies a sample config.
func (hub *Hub) BuildSite() *cli.Command {
	return &cli.Command{
		Name:    "build",
		Aliases: []string{"b"},
		Usage:   "Builds a static dist of all notes ready to be published on web.",
		Action:  hub.build,
	}
}

func (hub *Hub) build(cliCtx *cli.Context) error {
	posts, err := pipeline.ReadFiles(defaultPostDir)
	if err != nil {
		return err
	}
	err = pipeline.ReplaceLinks(posts)
	if err != nil {
		return err
	}
	err = pipeline.ConvertMarkdownToHTML(posts)
	if err != nil {
		return err
	}
	_, err = pipeline.MakeGraph(posts)
	if err != nil {
		return err
	}
	for _, p := range posts {
		if path.Base(p.FilePath) == defaultindexFileName {
			err := hub.renderIndex(p)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
