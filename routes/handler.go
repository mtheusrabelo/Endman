package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var json map[string]interface{}

func defaultHandler(c *gin.Context) {
	c.BindJSON(&json)
	c.JSON(http.StatusOK, gin.H{
		"body": json,
	})
}

func defaultGetHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
