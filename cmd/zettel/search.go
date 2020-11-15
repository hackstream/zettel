package main

import (
	"fmt"
	"path"
	"strings"

	"github.com/hackstream/zettel/internal/pipeline"
)

// PostData holds the data of a Post. Used for Indexing in Search. Has 3 keys,
// the Title of the page
// the Tags of the page if available
// the Permalink/Link of the Page
type PostData struct {
	Title     string   `json:"title"`
	Permalink string   `json:"permalink"`
	Tags      []string `json:"tags"`
}

// GenerateSearchIndex creates a index file(search.json) for search and indexing
func GenerateSearchIndex(posts []pipeline.Post) []PostData {
	indexData := make([]PostData, 0)
	for _, post := range posts {
		slug := strings.TrimSuffix(path.Base(post.FilePath), ".md")
		genPost := PostData{
			Title:     post.Meta.Title,
			Tags:      post.Meta.Tags,
			Permalink: fmt.Sprintf("posts/%s.html", slug),
		}
		if post.Meta.Title == "Index" {
			continue // Skips adding the Index page to search index
		} else {
			indexData = append(indexData, genPost)
		}
	}
	return indexData
}
