package pipeline

import (
	"time"
)

// Metadata contains the metadata extracted from the frontmatter of the post.
type Metadata struct {
	Date  time.Time `yaml:"date"`
	Tags  []string  `yaml:"tags"`
	Title string    `yaml:"title"`
}

// Link is the link given to another post.
type Link struct {
	Title, Slug string
}

// Post contains all the necessary things to render a post.
type Post struct {
	Meta     Metadata
	Body     string
	FilePath string

	// Links here only contains the links mentioned in the body.
	// These are not all the links to/from this post, since we need
	// to derive those from the graph.
	Links []Link
}

// NewPost returns a new Post
func NewPost(path string) Post {
	return Post{
		FilePath: path,
		Links:    make([]Link, 0),
		Meta:     Metadata{},
	}
}
