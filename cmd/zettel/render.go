package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/hackstream/zettel/internal/pipeline"
)

const (
	defaultDistDir = "dist"
)

func (hub *Hub) makeDist() error {
	dirs := []string{
		defaultDistDir,
		path.Join(defaultDistDir, "css"),
		path.Join(defaultDistDir, "images"),
	}

	for _, d := range dirs {
		err := createDirectory(d)
		if err != nil {
			return err
		}
	}

	globs := map[string]string{
		"/templates/layouts/css/*":    path.Join(defaultDistDir, "css"),
		"/templates/layouts/images/*": path.Join(defaultDistDir, "images"),
	}

	for g, dir := range globs {
		files, err := hub.Fs.Glob(g)
		if err != nil {
			return err
		}
		fmt.Printf("files: %v\n", files)
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
