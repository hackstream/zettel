package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/hackstream/zettel/internal/pipeline"
)

func (hub *Hub) renderIndex(post pipeline.Post) error {
	tmplContext := getInitialTmplContext(hub.Config)
	tmplContext["IsIndex"] = true
	tmplContext["Post"] = post

	path := filepath.Join(defaultDistDir, "index.html")
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	// render post template
	tmpls := []string{
		"templates/layouts/base.tmpl",
		"templates/layouts/header.tmpl",
		"templates/layouts/navbar.tmpl",
		"templates/layouts/post.tmpl",
		"templates/layouts/footer.tmpl",
	}
	err = saveResource("base", tmpls, file, tmplContext, hub.Fs)
	if err != nil {
		return err
	}
	return nil
}

func (hub *Hub) renderPost(post pipeline.Post) error {
	tmplContext := getInitialTmplContext(hub.Config)
	tmplContext["IsIndex"] = false
	tmplContext["Post"] = post

	slug := strings.TrimSuffix(path.Base(post.FilePath), ".md")
	path := filepath.Join(defaultDistDir, "posts", fmt.Sprintf("%s.html", slug))
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	// render post template
	tmpls := []string{
		"templates/layouts/base.tmpl",
		"templates/layouts/header.tmpl",
		"templates/layouts/navbar.tmpl",
		"templates/layouts/post.tmpl",
		"templates/layouts/footer.tmpl",
	}
	err = saveResource("base", tmpls, file, tmplContext, hub.Fs)
	if err != nil {
		return err
	}
	return nil
}

func (hub *Hub) renderTag(tag string, posts []pipeline.Post, isAllPosts bool) error {
	// Make links from posts
	links := []pipeline.Link{}
	for _, p := range posts {
		l := pipeline.Link{
			Slug:  strings.TrimSuffix(path.Base(p.FilePath), ".md"),
			Title: p.Meta.Title,
		}
		links = append(links, l)
	}
	tmplContext := getInitialTmplContext(hub.Config)
	tmplContext["TagName"] = tag
	tmplContext["Links"] = links
	path := filepath.Join(defaultDistDir, "all.html")
	if !isAllPosts {
		path = filepath.Join(defaultDistDir, "tags", fmt.Sprintf("%s.html", tag))
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	// render post template
	tmpls := []string{
		"templates/layouts/list.tmpl",
		"templates/layouts/header.tmpl",
		"templates/layouts/navbar.tmpl",
		"templates/layouts/footer.tmpl",
	}
	err = saveResource("list", tmpls, file, tmplContext, hub.Fs)
	if err != nil {
		return err
	}
	return nil
}

func (hub *Hub) renderGraphData(graphData GraphData) error {
	path := filepath.Join(defaultDistDir, "data", "graph.json")
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	err = json.NewEncoder(file).Encode(&graphData)
	if err != nil {
		return err
	}

	path = filepath.Join(defaultDistDir, "graph.html")
	file, err = os.Create(path)
	if err != nil {
		return err
	}
	// render post template
	tmpls := []string{
		"templates/layouts/graph.tmpl",
		"templates/layouts/header.tmpl",
		"templates/layouts/navbar.tmpl",
		"templates/layouts/footer.tmpl",
	}

	tmplContext := getInitialTmplContext(hub.Config)
	err = saveResource("graph", tmpls, file, tmplContext, hub.Fs)
	if err != nil {
		return err
	}
	return nil
}
