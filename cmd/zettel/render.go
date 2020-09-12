package main

import (
	"os"
	"path/filepath"

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
	// persist Post file.
	// create content dir if it doesn't exit
	if _, err := os.Stat(defaultDistDir); os.IsNotExist(err) {
		os.Mkdir(defaultDistDir, 0755)
	}
	path := filepath.Join(defaultDistDir, "index.html")
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	// render post template
	err = saveResource([]string{"templates/layouts/base.tmpl", "templates/layouts/header.tmpl", "templates/layouts/navbar.tmpl", "templates/layouts/footer.tmpl"}, file, tmplContext, hub.Fs)
	if err != nil {
		return err
	}
	return nil
}
