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
	Name string `json:"name" binding:"required"`
	Desc string `json:"desc" binding:"required"`
}

func createAuthor(c *gin.Context) {
	var schema authorSchema
	if err := c.BindJSON(&schema); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	id, err := lib.CreateAuthor(schema.Name, schema.Desc)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}
