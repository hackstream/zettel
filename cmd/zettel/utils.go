package main

import (
	"fmt"
	"io"
	"os"
	"text/template"

	"github.com/knadh/stuffbin"
)

// createFile takes a default config template file and writes to the current directory
func createFile(cfgFile []byte, configName string) error {
	f, err := os.Create(configName)
	if err != nil {
		return fmt.Errorf("error while creating default config: %v", err)
	}
	_, err = f.Write(cfgFile)
	if err != nil {
		return fmt.Errorf("error while copying default config: %v", err)
	}
	return nil
}

// parse takes in a template path and the variables to be "applied" on it. The rendered template
// is saved to the destination path.
func parse(name string, templateNames []string, fs stuffbin.FileSystem) (*template.Template, error) {
	tmpl := template.New(name)
	for _, t := range templateNames {
		// read template file
		c, err := fs.Read(t)
		if err != nil {
			return nil, fmt.Errorf("error reading template: %v", err)
		}
		tmpl, err = tmpl.Parse(string(c))
		if err != nil {
			return nil, fmt.Errorf("error parsing template: %v", err)
		}
	}
	return tmpl, nil
}

func writeTemplate(tmpl *template.Template, config map[string]interface{}, dest io.Writer) error {
	// apply the variable and save the rendered template to the file.
	err := tmpl.Execute(dest, config)
	if err != nil {
		return err
	}
	return nil
}

func saveResource(name string, templateNames []string, dest io.Writer, config map[string]interface{}, fs stuffbin.FileSystem) error {
	// parse template file
	tmpl, err := parse(name, templateNames, fs)
	if err != nil {
		return err
	}
	err = writeTemplate(tmpl, config, dest)
	if err != nil {
		return err
	}

	return nil
}

func createDirectory(dir string) error {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
		return nil
	}
	return err
}

func getInitialTmplContext(cfg Config) map[string]interface{} {
	tmplContext := make(map[string]interface{})
	tmplContext["SiteName"] = cfg.SiteName
	tmplContext["Description"] = cfg.Description
	tmplContext["SitePrefix"] = cfg.SitePrefix
	return tmplContext
}
