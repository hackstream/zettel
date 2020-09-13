package pipeline

import (
	"regexp"
	"strings"
)

var (
	// LinkRegex is the regex of links in posts
	LinkRegex = regexp.MustCompile(`(^|\s)(\[\[(\w|-)+\]\])(\s|$)`)
)

func findLinks(body string) []string {
	matches := LinkRegex.FindAllString(body, -1)
	for i, m := range matches {
		matches[i] = strings.TrimSpace(m)
	}
	return matches
}
