package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/knadh/stuffbin"
)

type FileSystem interface {
	Glob(string) ([]string, error)
	Read(string) ([]byte, error)
}

// assert that filesystem.StuffBin implements FileSystem in compile time
var _ FileSystem = (stuffbin.FileSystem)(nil)

// initBuiltinFileSystem initializes the stuffbin FileSystem to provide
// access to bunded static assets to the app.
func initBuiltinFileSystem() (stuffbin.FileSystem, error) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exPath := filepath.Dir(ex)
	fs, err := stuffbin.UnStuff(filepath.Join(exPath, filepath.Base(os.Args[0])))
	if err != nil {
		return nil, err
	}

	return fs, nil
}

type DiskFileSystem struct {
	root string
}

func initDiskFileSystem(dir string) (FileSystem, error) {
	dir = strings.TrimSuffix(dir, "/")

	info, err := os.Stat(dir)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", dir)
	}

	return &DiskFileSystem{root: dir}, nil
}

func (disk *DiskFileSystem) Glob(pattern string) ([]string, error) {
	matches, err := filepath.Glob(filepath.Join(disk.root, pattern))
	if err != nil {
		return nil, nil
	}

	for i, value := range matches {
		matches[i] = strings.TrimPrefix(value, disk.root)
	}

	return matches, nil
}

func (disk *DiskFileSystem) Read(relative string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(disk.root, relative))
}
