package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

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
