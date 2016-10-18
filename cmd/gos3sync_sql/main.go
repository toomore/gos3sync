package main

import "github.com/toomore/gos3sync/cmd/gos3sync_sql/cmd"

//func showHelp() {
//	fmt.Println(`
//usage: gos3sync_sql [command]
//
//command:
//	init [path]
//`)
//}

//func main() {
//	flag.Parse()
//	if len(flag.Args()) < 1 {
//		showHelp()
//		os.Exit(0)
//	}
//	switch flag.Arg(0) {
//	case "init":
//		syncdb.Init(flag.Arg(1))
//	case "path":
//		log.Println(filepath.Abs("./"))
//	case "add":
//		fw := &gos3sync.FileWalk{}
//		fw.Walk(flag.Arg(1))
//		log.Println(fw.Paths)
//		log.Println("Do add file")
//	case "commit":
//		log.Println("Save to index")
//	case "push":
//		log.Println("Upload to AWS S3")
//	default:
//		showHelp()
//	}
//}

func main() {
	cmd.Execute()
}
