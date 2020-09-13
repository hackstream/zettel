package main

import (
	"path"
	"strings"

	"github.com/hackstream/zettel/internal/pipeline"
	"github.com/urfave/cli/v2"
)

// BuildSite initializes a new zettel site copies a sample config.
func (hub *Hub) BuildSite() *cli.Command {
	return &cli.Command{
		Name:    "build",
		Aliases: []string{"b"},
		Usage:   "Builds a static dist of all notes ready to be published on web.",
		Action:  hub.MustHaveConfig(hub.build),
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
	g, err := pipeline.MakeGraph(posts)
	if err != nil {
		return err
	}
	err = hub.makeDist()
	if err != nil {
		return err
	}
	for i := 0; i < len(posts); i++ {
		p := posts[i]
		// If index, call render index
		if path.Base(p.FilePath) == defaultindexFileName {
			err := hub.renderIndex(p)
			if err != nil {
				return err
			}
			continue
		}

		// Get connection indices from the graph
		conns := []pipeline.Link{}
		do := func(n int, c int64) bool {
			// Check if this post is one of the links mentioned in the body of the post.
			// If it is, we will skip, since we only need connections from other post to this.
			co := posts[n]
			isInnerLink := false
			s := strings.TrimRight(path.Base(co.FilePath), ".md")
			for _, l := range p.Links {
				if l.Slug == s {
					isInnerLink = true
					break
				}
			}
			coLink := pipeline.Link{
				Title: co.Meta.Title,
				Slug:  s,
			}
			if !isInnerLink {
				conns = append(conns, coLink)
			}
			return false
		}

		// Get connections to the post
		g.Visit(i, do)

		p.Connections = conns

		err := hub.renderPost(p)
		if err != nil {
			return err
		}
	}
	return nil
}
