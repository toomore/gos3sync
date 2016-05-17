package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
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
				hashPath[string(md5sum(f))] = path
				fmt.Printf(".")
			}
		}
	} else {
		log.Println(err)
	}
	return nil
}

func saveData(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, os.ModePerm)
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
			fmt.Println([]byte(k))
			fmt.Println(v)
			//_data = append(_data, fmt.Sprintf("%x", k)...)
			_data = append(_data, hex.EncodeToString([]byte(k))...)
			_data = append(_data, "\n"...)
		}
		f, err := os.Create("tt.txt")
		defer f.Close()
		if err == nil {
			fmt.Printf("%x", _data)
			//f.Write([]byte(fmt.Sprintf("%x", _data)))
			f.Write(_data)
		} else {
			log.Println(">>>>", err)
		}
		fmt.Println(saveData("t.txt", _data))
	}
}
