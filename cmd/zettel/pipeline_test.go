package main

import (
	"os"
	"testing"
)

func TestReadFiles(t *testing.T) {
	dir := os.Getenv("TEST_DIR")

	posts, err := readFiles(dir)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("posts: %v", posts)
}
