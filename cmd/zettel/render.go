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

	fpath := filepath.Join(defaultDistDir, "index.html")

	file, err := os.Create(fpath)
	if err != nil {
		return err
	}

	// render post template
	tmpls := []string{
		"layouts/base.tmpl",
		"layouts/header.tmpl",
		"layouts/navbar.tmpl",
		"layouts/post.tmpl",
		"layouts/footer.tmpl",
	}

	// Add custom template path
	for i := 0; i < len(tmpls); i++ {
		tmpls[i] = path.Join(hub.Fs.TemplatePath, tmpls[i])
	}

	return saveResource("base", tmpls, file, tmplContext, hub.Fs.Fs)
}

func (hub *Hub) renderPost(post pipeline.Post) error {
	tmplContext := getInitialTmplContext(hub.Config)
	tmplContext["IsIndex"] = false
	tmplContext["Post"] = post

	slug := strings.TrimSuffix(path.Base(post.FilePath), ".md")
	fpath := filepath.Join(defaultDistDir, "posts", fmt.Sprintf("%s.html", slug))

	file, err := os.Create(fpath)
	if err != nil {
		return err
	}

	// render post template
	tmpls := []string{
		"layouts/base.tmpl",
		"layouts/header.tmpl",
		"layouts/navbar.tmpl",
		"layouts/post.tmpl",
		"layouts/footer.tmpl",
	}

	// Add custom template path
	for i := 0; i < len(tmpls); i++ {
		tmpls[i] = path.Join(hub.Fs.TemplatePath, tmpls[i])
	}

	return saveResource("base", tmpls, file, tmplContext, hub.Fs.Fs)
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
	fpath := filepath.Join(defaultDistDir, "all.html")

	if !isAllPosts {
		fpath = filepath.Join(defaultDistDir, "tags", fmt.Sprintf("%s.html", tag))
	}

	file, err := os.Create(fpath)
	if err != nil {
		return err
	}

	// render post template
	tmpls := []string{
		"layouts/list.tmpl",
		"layouts/header.tmpl",
		"layouts/navbar.tmpl",
		"layouts/footer.tmpl",
	}

	// Add custom template path
	for i := 0; i < len(tmpls); i++ {
		tmpls[i] = path.Join(hub.Fs.TemplatePath, tmpls[i])
	}

	return saveResource("list", tmpls, file, tmplContext, hub.Fs.Fs)
}

func (hub *Hub) renderGraphData(graphData GraphData) error {
	fpath := filepath.Join(defaultDistDir, "data", "graph.json")

	file, err := os.Create(fpath)
	if err != nil {
		return err
	}

	if err = json.NewEncoder(file).Encode(&graphData); err != nil {
		return err
	}

	fpath = filepath.Join(defaultDistDir, "graph.html")

	file, err = os.Create(fpath)
	if err != nil {
		return err
	}

	// render post template
	tmpls := []string{
		"layouts/graph.tmpl",
		"layouts/header.tmpl",
		"layouts/navbar.tmpl",
		"layouts/footer.tmpl",
	}

	// Add custom template path
	for i := 0; i < len(tmpls); i++ {
		tmpls[i] = path.Join(hub.Fs.TemplatePath, tmpls[i])
	}

	tmplContext := getInitialTmplContext(hub.Config)

	return saveResource("graph", tmpls, file, tmplContext, hub.Fs.Fs)
}
