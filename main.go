package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/toomore/gos3sync/hashlib"
)

//var hashPath = make(map[string]string)
//
//func walkfunc(path string, info os.FileInfo, err error) error {
//	if err == nil {
//		if !info.IsDir() {
//			f, err := os.Open(path)
//			defer f.Close()
//			if err == nil {
//				hashPath[string(hashlib.Sum(f))] = path
//				fmt.Printf(".")
//			}
//		}
//	} else {
//		log.Println(err)
//	}
//	return nil
//}

func main() {
	flag.Parse()
	for _, path := range flag.Args() {
		fw := NewFileWalk(path)
		fw.Walk()
		if len(fw.Paths()) > 0 {
			fmt.Println()
			var _data = make([]byte, 0, 100)
			for k, v := range fw.Paths() {
				//fmt.Println("Key: ", []byte(k))
				//fmt.Println("Path: ", v)
				_data = append(_data, hashlib.String([]byte(k))...)
				_data = append(_data, " "...)
				_data = append(_data, []byte(v)...)
				_data = append(_data, "\n"...)
			}
			if err := ioutil.WriteFile("t.txt", _data, 0600); err != nil {
				log.Println("Error save", err)
			}
		}
	}
}
