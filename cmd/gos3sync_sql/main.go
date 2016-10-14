package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/toomore/gos3sync/syncdb"
)

func showHelp() {
	fmt.Println(`
usage: gos3sync_sql [command]

command:
	init [path]
`)
}

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		showHelp()
		os.Exit(0)
	}
	switch flag.Arg(0) {
	case "init":
		syncdb.Init(flag.Arg(1))
	case "path":
		log.Println(filepath.Abs("./"))
	case "add":
		log.Println("Do add file")
	case "commit":
		log.Println("Save to index")
	case "push":
		log.Println("Upload to AWS S3")
	default:
		showHelp()
	}
}
