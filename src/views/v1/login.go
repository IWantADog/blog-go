package v1

import (
	"net/http"

	"github.com/blog/src/lib"
	"github.com/gin-gonic/gin"
)

type requestCommonHeader struct {
	Token string `header:"x-auth-token"`
}

type loginSchema struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func login(c *gin.Context) {
	var schema loginSchema
	if err := c.BindJSON(&schema); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authorInfo, err := lib.Login(schema.Name, schema.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, authorInfo)
}

func logout(c *gin.Context) {
	var header requestCommonHeader
	if err := c.BindHeader(&header); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lib.Logout(header.Token)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
