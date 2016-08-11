package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/toomore/gos3sync/hashlib"
)

var outputFilename = flag.String("o", "result.txt", "output data")

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
			if err := ioutil.WriteFile(*outputFilename, _data, 0600); err != nil {
				log.Println("Error save", err)
			}
		}
	}
}
