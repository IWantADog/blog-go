package v1

import (
	"net/http"

	"github.com/blog/src/lib"
	"github.com/gin-gonic/gin"
)

func getAuthor(c *gin.Context) {
	authors := lib.GetAuthors()
	c.JSON(http.StatusOK, authors)
}

type authorSchema struct {
	Name     string `json:"name" binding:"required"`
	Desc     string `json:"desc" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func createAuthor(c *gin.Context) {
	var schema authorSchema
	if err := c.BindJSON(&schema); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author, err := lib.CreateAuthor(schema.Name, schema.Password, schema.Desc)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, author)
}
