package pipeline

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	mathjax "github.com/athul/goldmark-mathjax"
	"github.com/yourbasic/graph"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
	"gopkg.in/yaml.v3"
)

// ReadFiles reads the given directory and appends into given posts
func ReadFiles(directory string) ([]Post, error) {
	posts := []Post{}
	err := filepath.Walk(directory, func(pth string, info os.FileInfo, err error) error {
		// Return if there is an error
		if err != nil {
			return err
		}
		// Skip all directories
		if info.IsDir() {
			return nil
		}

		// Skip any files which aren't markdown.
		if !strings.HasSuffix(path.Base(pth), ".md") {
			return nil
		}

		post := NewPost(pth)

		// Read the file
		data, err := ioutil.ReadFile(pth)
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

		// If the post is a draft, skip the file
		if post.Meta.Draft {
			return nil
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

// ReplaceLinks replaces all the `[[]]` links in the body with the link to the HTML page.
func ReplaceLinks(posts []Post, sitePrefix string) error {
	// First we need a map of all the files, so that we know
	// if there's a link to any non existent file and also
	// to get the metadata of the file
	slugs := make(map[string]Metadata)
	for _, p := range posts {
		pth := path.Base(p.FilePath)
		slug := strings.TrimSuffix(pth, ".md")
		slugs[slug] = p.Meta
	}

	// Loop over and replace the `[[ ]]` link.
	for i, p := range posts {
		// Find all the matches first
		matches := findLinks(p.Body)
		for _, m := range matches {
			sg := strings.TrimLeft(m, "[[")
			sg = strings.TrimSuffix(sg, "]]")

			// Find if the slug exists
			meta, ok := slugs[sg]
			if !ok {
				return fmt.Errorf("link to an invalid slug: %s", sg)
			}

			link := fmt.Sprintf(`[%s](%s/posts/%s.html)`, meta.Title, sitePrefix, sg)

			// Replace the slug with a link
			posts[i].Body = strings.ReplaceAll(posts[i].Body, m, link)
			l := Link{
				Slug:  sg,
				Title: meta.Title,
			}
			posts[i].Links = append(posts[i].Links, l)
		}
	}

	return nil
}

// ConvertMarkdownToHTML converts post's body into HTML
func ConvertMarkdownToHTML(posts []Post, syntaxStyle string) error {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM, mathjax.MathJax),
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)

	for i, p := range posts {
		html := bytes.NewBuffer([]byte{})

		if err := md.Convert([]byte(p.Body), html); err != nil {
			return err
		}
		body, err := SyntaxHighlighter(html.Bytes(), syntaxStyle)
		if err != nil {
			return err
		}
		posts[i].Body = body
	}

	return nil
}

// MakeGraph returns a graph for a given list of `Posts`.
func MakeGraph(posts []Post) (*graph.Mutable, error) {
	// Make a map of all posts with their index
	// so that it would be easy to lookup the index using the slug
	slugs := make(map[string]int)
	for i, p := range posts {
		pth := path.Base(p.FilePath)
		slug := strings.TrimSuffix(pth, ".md")
		slugs[slug] = i
	}

	g := graph.New(len(posts))

	for i, p := range posts {
		for _, link := range p.Links {
			linkIndex := slugs[link.Slug]

			g.AddBoth(i, linkIndex)
		}
	}

	return g, nil
}
