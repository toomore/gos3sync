package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func md5sum(f *os.File) []byte {
	hash := md5.New()
	io.Copy(hash, f)
	return hash.Sum(nil)
}

func walkfeed(path string, info os.FileInfo, err error) error {
	if err == nil {
		if info.IsDir() {
			fmt.Println("Folder: ", path)
		} else {
			fmt.Println(info.Mode().String(), info.IsDir(), path, info.Size())
			f, err := os.Open(path)
			defer f.Close()
			if err == nil {
				fmt.Printf("%x\n", md5sum(f))
			}
		}
	} else {
		log.Println(err)
	}
	return nil
}

func main() {
	err := filepath.Walk("./", walkfeed)
	fmt.Println(err)
}
