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

var hashPath = make(map[string]string)

func walkfunc(path string, info os.FileInfo, err error) error {
	if err == nil {
		if !info.IsDir() {
			f, err := os.Open(path)
			defer f.Close()
			if err == nil {
				hashsum := fmt.Sprintf("%x", md5sum(f))
				hashPath[hashsum] = path
			}
		}
	} else {
		log.Println(err)
	}
	return nil
}

func main() {
	if err := filepath.Walk("./", walkfunc); err != nil {
		fmt.Println(err)
	}
	for k, v := range hashPath {
		fmt.Println(k, v)
	}
}
