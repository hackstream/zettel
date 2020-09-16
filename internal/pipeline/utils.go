package pipeline

import (
	"bufio"
	"bytes"
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

//SyntaxHighlighter Highlights any code inside the MD Files
func SyntaxHighlighter(html []byte, syntaxStyle string) string {
	byteReader := bytes.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(byteReader)
	if err != nil {
		return ""
	}
	var hlErr error
	doc.Find("code[class*=\"language-\"]").Each(func(i int, s *goquery.Selection) {
		if hlErr != nil {
			return
		}
		class, _ := s.Attr("class")
		lang := strings.TrimPrefix(class, "language-")
		oldCode := s.Text()
		lexer := lexers.Get(lang)
		formatter := synhtml.New(synhtml.WithClasses(false))
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
		return ""
	}
	new, err := doc.Html()
	if err != nil {
		return ""
	}
	// replace unnecessarily added html tags
	return new
}
