package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/toomore/gos3sync/hashlib"
)

// FILEMOD default
const FILEMOD = 0600

var hashPath = make(map[string]string)

func walkfunc(path string, info os.FileInfo, err error) error {
	if err == nil {
		if !info.IsDir() {
			f, err := os.Open(path)
			defer f.Close()
			if err == nil {
				hashPath[string(hashlib.Sum(f))] = path
				fmt.Printf(".")
			}
		}
	} else {
		log.Println(err)
	}
	return nil
}

func saveData(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, FILEMOD)
}

func main() {
	flag.Parse()
	for _, path := range flag.Args() {
		if err := filepath.Walk(path, walkfunc); err != nil {
			fmt.Println(err)
		}
	}
	if len(hashPath) > 0 {
		fmt.Println()
		var _data = make([]byte, 0, 100)
		for k, v := range hashPath {
			fmt.Println("Key: ", []byte(k))
			fmt.Println("Path: ", v)
			//_data = append(_data, fmt.Sprintf("%x", k)...)
			_data = append(_data, hashlib.String([]byte(k))...)
			_data = append(_data, " "...)
			_data = append(_data, []byte(v)...)
			_data = append(_data, "\n"...)
		}
		if err := saveData("t.txt", _data); err != nil {
			log.Println("Error save", err)
		}
	}
}
