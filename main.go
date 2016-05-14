package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func walkfeed(path string, info os.FileInfo, err error) error {
	if err == nil {
		if info.IsDir() {
			fmt.Println("Folder: ", path)
		} else {
			fmt.Println(info.Mode().String(), info.IsDir(), path, info.Size())
		}
	} else {
		log.Println(err)
	}
	return nil
}

func main() {
	err := filepath.Walk("/Volumes/RamDisk/", walkfeed)
	fmt.Println(err)
}
