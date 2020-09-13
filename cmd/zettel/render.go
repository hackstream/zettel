package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/hackstream/zettel/internal/pipeline"
)

const (
	defaultDistDir = "dist"
)

func (hub *Hub) renderIndex(post pipeline.Post) error {
	tmplContext := make(map[string]interface{})
	tmplContext["SiteName"] = hub.Config.SiteName
	tmplContext["Description"] = hub.Config.Description
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
	tmplContext := make(map[string]interface{})
	tmplContext["SiteName"] = hub.Config.SiteName
	tmplContext["Description"] = hub.Config.Description
	tmplContext["IsIndex"] = false
	tmplContext["Post"] = post

	slug := strings.Trim(path.Base(post.FilePath), ".md")
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
			Slug:  strings.TrimRight(path.Base(p.FilePath), ".md"),
			Title: p.Meta.Title,
		}
		links = append(links, l)
	}
	tmplContext := make(map[string]interface{})
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
