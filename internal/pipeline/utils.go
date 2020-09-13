package pipeline

import "regexp"

var (
	// LinkRegex is the regex of links in posts
	LinkRegex = regexp.MustCompile(`\[\[.+\]\]`)
)

func findLinks(body string) []string {
	return LinkRegex.FindAllString(body, -1)
}
