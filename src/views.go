package src

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	g := App.Group("/v1")
	g.GET("hello/", helloWorld)
}

func helloWorld(c *gin.Context) {
	c.String(http.StatusOK, "hello, world")
}

func getAllBlog(c *gin.Context) {
}
