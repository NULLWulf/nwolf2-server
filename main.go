package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	status := r.Group("/nwolf2/status")
	{
		status.GET("/", getSystemDate())
		status.DELETE("/", invalidRequest())
		status.POST("/", invalidRequest())
		status.PATCH("/", invalidRequest())
		status.PUT("/", invalidRequest())
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getSystemDate() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"system time": time.Now().UTC().Format(time.RFC3339),
		})
	}
}

func invalidRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"status": http.StatusMethodNotAllowed})
	}
}
