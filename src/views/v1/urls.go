package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(app *gin.Engine) {
	group := app.Group("/v1")
	group.GET("hello/", helloWorld)
	// blogs
	group.GET("blogs/", getAllBlog)

	// authors
	group.GET("authors/", getAuthor)
	group.POST("authors/", createAuthor)

	// login
	group.POST("login/", login)
}

func helloWorld(c *gin.Context) {
	c.String(http.StatusOK, "hello, world")
}
