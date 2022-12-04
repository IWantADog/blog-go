package main

import (
	"flag"

	"github.com/blog/src"
	"github.com/blog/src/model"
)

var createDB = flag.Bool("createDB", false, "init db schema")
var runServer = flag.Bool("runServer", false, "start app")

func main() {
	flag.Parse()
	if *createDB {
		db := src.DB
		db.AutoMigrate(&model.Author{})
		db.AutoMigrate(&model.Tag{})
		db.AutoMigrate(&model.Blog{})
	} else if *runServer {
		app := src.App
		app.Run()
	}
}
