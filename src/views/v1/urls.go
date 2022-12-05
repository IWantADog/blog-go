package v1

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/blog/src/lib"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(app *gin.Engine) {
	group := app.Group("/v1")
	group.GET("hello/", helloWorld)
	group.GET("blogs/", getAllBlog)
}

func helloWorld(c *gin.Context) {
	c.String(http.StatusOK, "hello, world")
}

func getAllBlog(c *gin.Context) {
	blogs := lib.GetAllBlog()
	blog_json, err := json.Marshal(blogs)
	if err != nil {
		log.Fatalln("get blogs failed")
	}
	c.String(http.StatusOK, string(blog_json))
}
