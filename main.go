package main

import (
	"flag"

	"github.com/blog/src"
)

var createDB = flag.Bool("createDB", false, "init db schema")
var runServer = flag.Bool("runServer", false, "start app")

func main() {
	flag.Parse()
	if *createDB {
		db := src.DB
		db.AutoMigrate(&src.Author{})
		db.AutoMigrate(&src.Tag{})
		db.AutoMigrate(&src.Blog{})
	} else if *runServer {
		app := src.App
		app.Run()
	}
}
