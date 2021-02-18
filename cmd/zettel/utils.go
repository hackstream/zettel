package main

import (
	"embed"
	"fmt"
	"io"
	"os"
	"text/template"
)

// createFile takes a default config template file and writes to the current directory
func createFile(cfgFile []byte, configName string) error {
	f, err := os.Create(configName)
	if err != nil {
		return fmt.Errorf("error while creating default config: %v", err)
	}

	if _, err = f.Write(cfgFile); err != nil {
		return fmt.Errorf("error while copying default config: %v", err)
	}

	return nil
}

// parse takes in a template path and the variables to be "applied" on it. The rendered template
// is saved to the destination path.
func parse(name string, templateNames []string, fs embed.FS) (*template.Template, error) {
	tmpl := template.New(name)

	for _, t := range templateNames {
		// read template file
		c, err := fs.ReadFile(t)
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
	return tmpl.Execute(dest, config)
}

func saveResource(name string, templateNames []string, dest io.Writer, config map[string]interface{}, fs embed.FS) error {
	// parse template file
	tmpl, err := parse(name, templateNames, fs)
	if err != nil {
		return err
	}

	return writeTemplate(tmpl, config, dest)
}

func createDirectory(dir string) error {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		_ = os.Mkdir(dir, 0750)
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
