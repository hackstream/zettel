package pipeline

import (
	"os"
	"reflect"
	"testing"
)

func TestReadFiles(t *testing.T) {
	dir := os.Getenv("TEST_DIR")

	posts, err := ReadFiles(dir)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("posts: %v", posts)
}

func TestFindLinks(t *testing.T) {
	body := "#Post1\n\n- Index: [[index]]\n"

	matches := findLinks(body)

	m := []string{"[[index]]"}

	if !reflect.DeepEqual(matches, m) {
		t.Errorf("expected: %v, got: %v", m, matches)
	}
}

func TestReplaceLinks(t *testing.T) {
	dir := os.Getenv("TEST_DIR")

	posts, err := ReadFiles(dir)
	if err != nil {
		t.Fatal(err)
	}

	err = ReplaceLinks(posts)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("posts: %#v", posts)
}

func TestConvertMarkdownToHTML(t *testing.T) {
	dir := os.Getenv("TEST_DIR")

	posts, err := ReadFiles(dir)
	if err != nil {
		t.Fatal(err)
	}

	err = ReplaceLinks(posts)
	if err != nil {
		t.Fatal(err)
	}

	err = ConvertMarkdownToHTML(posts, "dracula")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("posts: %#v", posts)
}
