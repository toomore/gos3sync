package gos3sync

import (
	"os"
	"path/filepath"
)

// FileWalk struct
type FileWalk struct {
	Paths []string
}

// Walk do walk
func (f *FileWalk) Walk(path string) {
	filepath.Walk(path, f.WalkFunc)
}

// WalkFunc to walk all path
func (f *FileWalk) WalkFunc(path string, info os.FileInfo, err error) error {
	if err == nil {
		f.Paths = append(f.Paths, path)
	}
	return err
}
