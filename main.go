package main

import (
	"github.com/blog/src/global"
	v1 "github.com/blog/src/views/v1"
	"github.com/gin-gonic/gin"
)

// var createDB = flag.Bool("createDB", false, "init db schema")
// var runServer = flag.Bool("runServer", false, "start app")

func main() {
	global.LoadSettings()
	app := gin.Default()
	v1.RegisterRoute(app)
	app.Run()
}
