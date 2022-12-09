package v1

import (
	"net/http"

	"github.com/blog/src/lib"
	"github.com/gin-gonic/gin"
)

func getAllBlog(c *gin.Context) {
	blogs := lib.GetAllBlog()
	c.JSON(http.StatusOK, blogs)
}
