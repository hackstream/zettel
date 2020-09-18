package pipeline

import (
	"bufio"
	"bytes"
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	synhtml "github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
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

// SyntaxHighlighter Highlights any code-blocks from the Generated HTML files(from markdown)
// based on the style defined in the pygmentsStyle key in zettel.toml file
// It takes in the rendered HTML from markdown, parses it and selects the code blocks
// within the <code></code> tag and the language from the tag's class
// It tokenizes the code based on the language and applies highlighting.
// It uses styles from the chroma library in Go
func SyntaxHighlighter(html []byte, syntaxStyle string) (string, error) {
	byteReader := bytes.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(byteReader)
	if err != nil {
		log.Printf("Error parsing HTML: %s", err.Error())
		return "", err
	}
	var hlErr error
	doc.Find("code[class*=\"language-\"]").Each(func(i int, s *goquery.Selection) {
		if hlErr != nil {
			return
		}
		var (
			class, _  = s.Attr("class")
			lang      = strings.TrimPrefix(class, "language-")
			oldCode   = s.Text()
			lexer     = lexers.Get(lang)
			formatter = synhtml.New(synhtml.WithClasses(false))
		)
		iterator, err := lexer.Tokenise(nil, string(oldCode))
		if err != nil {
			hlErr = err
			return
		}

		b := bytes.Buffer{}
		buf := bufio.NewWriter(&b)

		if err := formatter.Format(buf, styles.Get(syntaxStyle), iterator); err != nil {
			hlErr = err
			return
		}

		if err := buf.Flush(); err != nil {
			hlErr = err
			return
		}

		s.SetHtml(b.String())
	})
	if hlErr != nil {
		return "", hlErr
	}

	// Converts document to HTML.
	out, err := doc.Html()
	if err != nil {
		return "", hlErr
	}

	// replace unnecessarily added html tags
	return out, nil
}
