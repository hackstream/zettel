package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/hackstream/zettel/internal/pipeline"
)

func (hub *Hub) makeDist() error {
	dirs := []string{
		defaultDistDir,
		path.Join(defaultDistDir, "css"),
		path.Join(defaultDistDir, "images"),
		path.Join(defaultDistDir, "posts"),
	}

	for _, d := range dirs {
		err := createDirectory(d)
		if err != nil {
			return err
		}
	}

	// Copy css, images folders to dist directory
	globs := map[string]string{
		"/templates/layouts/css/*":    path.Join(defaultDistDir, "css"),
		"/templates/layouts/images/*": path.Join(defaultDistDir, "images"),
	}

	for g, dir := range globs {
		files, err := hub.Fs.Glob(g)
		if err != nil {
			return err
		}
		for _, f := range files {
			b, err := hub.Fs.Read(f)
			if err != nil {
				return err
			}

			fd, err := os.Create(path.Join(dir, path.Base(f)))
			if err != nil {
				return err
			}

			_, err = fd.Write(b)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

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
