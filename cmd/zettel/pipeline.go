package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// readFiles reads the given directory and appends into given posts
func readFiles(directory string) ([]Post, error) {
	posts := []Post{}
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		// log.Printf("calling for %s, %v, %v", path, info, err)
		// Return if there is an error
		if err != nil {
			return err
		}
		// Skip all directories
		if info.IsDir() {
			return nil
		}

		post := Post{
			FilePath: path,
			Meta:     Metadata{},
		}

		// Read the file
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		// Extract metadata
		splits := bytes.Split(data, []byte("---\n"))
		frontMatter := splits[1]
		err = yaml.Unmarshal(frontMatter, &post.Meta)
		if err != nil {
			log.Printf("error while yaml unmarshal: %v", err)
			return err
		}

		// Join the body back ignoring the frontmatter
		post.Body = string(bytes.Join(splits[2:], []byte("---\n")))

		// Append the post to the posts
		posts = append(posts, post)
		return nil
	})
	if err != nil {
		return []Post{}, err
	}
	return posts, nil
}

// replaceLinks replaces all the `[[]]` links in the body with the link to the HTML page.
func replaceLinks(posts []Post) error {
	// First we need a map of all the files, so that we know
	// if there's a link to any non existent file and also
	// to get the metadata of the file
	slugs := make(map[string]Metadata)
	for _, p := range posts {
		pth := path.Base(p.FilePath)
		slug := strings.TrimRight(pth, ".md")
		slugs[slug] = p.Meta
	}

	// Loop over and replace the `[[ ]]` link.
	for i, p := range posts {
		// Find all the matches first
		matches := findLinks(p.Body)
		for _, m := range matches {
			sg := strings.TrimLeft(m, "[[")
			sg = strings.TrimRight(sg, "]]")

			// Find if the slug exists
			meta, ok := slugs[sg]
			if !ok {
				return fmt.Errorf("link to an invalid slug: %s", sg)
			}

			link := fmt.Sprintf(`<a href="/posts/%s.html">%s</a>`, sg, meta.Title)

			// Replace the slug with a link
			posts[i].Body = strings.ReplaceAll(p.Body, m, link)
		}
	}

	return nil
}
