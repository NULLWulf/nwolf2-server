package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func notFoundResponse(c *gin.Context) {
	populateLog(c, 405)
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "404 - Resource Not found",
	})
}

func getSystemDate(c *gin.Context) {
	populateLog(c, 200)
	c.IndentedJSON(http.StatusOK, gin.H{
		"system-time": time.Now().UTC().Format(time.RFC3339),
	})
}

func invalidRequest(c *gin.Context) {
	populateLog(c, 405)
	c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{
		"message": "405 - Method not allowed",
	})
}
