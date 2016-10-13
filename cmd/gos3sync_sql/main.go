package main

import (
	"flag"

	"github.com/toomore/gos3sync/syncdb"
)

func main() {
	flag.Parse()
	if len(flag.Args()) >= 1 {
		syncdb.Init(flag.Arg(0))
	}
}
