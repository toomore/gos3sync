package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/toomore/gos3sync/hashlib"
)

// FileWalk struct
type FileWalk struct {
	hashPath  map[string]string
	StartPath string
}

// NewFileWalk func
func NewFileWalk(StartPath string) *FileWalk {
	return &FileWalk{
		hashPath:  make(map[string]string),
		StartPath: StartPath,
	}
}

// Walk func
func (fw *FileWalk) Walk() {
	filepath.Walk(fw.StartPath, fw.walkfunc)
}

// Paths func
func (fw FileWalk) Paths() map[string]string {
	return fw.hashPath
}

func (fw *FileWalk) walkfunc(path string, info os.FileInfo, err error) error {
	if err == nil {
		if !info.IsDir() {
			f, err := os.Open(path)
			defer f.Close()
			if err == nil {
				fw.hashPath[string(hashlib.Sum(f))] = path
				fmt.Printf(".")
			}
		}
	} else {
		log.Println(err)
	}
	return nil
}
