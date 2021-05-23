package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
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
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "templates",
				Aliases: []string{"t"},
				Usage:   "Path to custom templates",
			},
		},
		Usage:  "Builds a static dist of all notes ready to be published on web.",
		Action: hub.MustHaveConfig(hub.build),
	}
}

func (hub *Hub) makeDist() error {
	// Clear dist directory
	if err := os.RemoveAll(defaultDistDir); err != nil {
		return err
	}

	dirs := []string{
		defaultDistDir,
		path.Join(defaultDistDir, "css"),
		path.Join(defaultDistDir, "images"),
		path.Join(defaultDistDir, "posts"),
		path.Join(defaultDistDir, "tags"),
		path.Join(defaultDistDir, "data"),
	}

	for _, d := range dirs {
		if err := createDirectory(d); err != nil {
			return err
		}
	}

	// Copy css, images folders to dist directory
	globs := map[string]string{
		"layouts/css":    path.Join(defaultDistDir, "css"),
		"layouts/images": path.Join(defaultDistDir, "images"),
	}

	for g, dir := range globs {
		p := path.Join(hub.Fs.TemplatePath, g)
		files, err := fs.ReadDir(hub.Fs.Fs, p)
		if err != nil {
			log.Printf("p: %v", p)
			return err
		}

		for _, f := range files {
			b, err := hub.Fs.Fs.Open(path.Join(p, f.Name()))
			if err != nil {
				return err
			}
			defer b.Close()

			data, err := io.ReadAll(b)
			if err != nil {
				return err
			}

			fd, err := os.Create(path.Join(dir, f.Name()))
			if err != nil {
				return err
			}
			defer fd.Close()

			if _, err = fd.Write(data); err != nil {
				return err
			}
		}
	}

	return nil
}

func (hub *Hub) build(cliCtx *cli.Context) error {
	fmt.Printf("config: %v\n", hub.Config)
	// If there is a custom template directory, change fs in hub
	customTemplatesPath := cliCtx.Value("templates").(string)
	if customTemplatesPath != "" {
		fs := os.DirFS(customTemplatesPath)
		hub.Fs = Fs{
			Fs:           fs,
			TemplatePath: "",
		}
	}

	posts, err := pipeline.ReadFiles(defaultPostDir)
	if err != nil {
		return err
	}

	if err = pipeline.ReplaceLinks(posts, hub.Config.SitePrefix, hub.Config.StripHTML); err != nil {
		return err
	}

	if err = pipeline.ConvertMarkdownToHTML(posts, hub.Config.Pygmentsstyle); err != nil {
		return err
	}

	g, err := pipeline.MakeGraph(posts)
	if err != nil {
		return err
	}

	if err = hub.makeDist(); err != nil {
		return err
	}

	// Aggregate all tags
	tags := make(map[string][]pipeline.Post)

	var (
		ps []pipeline.Post
		ok bool
	)

	for i := 0; i < len(posts); i++ {
		p := posts[i]
		// Put this post in the appropriate tags
		for _, t := range p.Meta.Tags {
			ps, ok = tags[t]
			if !ok {
				tags[t] = []pipeline.Post{p}
				continue
			}

			ps = append(ps, p)
			tags[t] = ps
		}

		// If index, call render index
		if path.Base(p.FilePath) == defaultindexFileName {
			if err = hub.renderIndex(p); err != nil {
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

			s := strings.TrimSuffix(path.Base(co.FilePath), ".md")
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

		if err = hub.renderPost(p); err != nil {
			return err
		}
	}

	// Render tags
	for tag, ps := range tags {
		if err = hub.renderTag(tag, ps, false); err != nil {
			return err
		}
	}

	// Render all posts file and remove index
	postsWithoutIndex := make([]pipeline.Post, 0, len(posts)-1)

	for _, p := range posts {
		if path.Base(p.FilePath) == defaultindexFileName {
			continue
		}

		postsWithoutIndex = append(postsWithoutIndex, p)
	}

	err = hub.renderTag("All Posts", postsWithoutIndex, true)
	if err != nil {
		return err
	}

	// Render graph.json
	gd := MakeGraphData(posts, g)

	return hub.renderGraphData(gd)
}
