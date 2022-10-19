package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/nullwulf/loggly"
	"net/http"
	"time"
)

type RequestLog struct {
	Method      string
	SourceIP    string
	RequestPath string
	StatusCode  int
}

func main() {
	r := gin.Default()
	status := r.Group("/nwolf2/status/")
	{
		status.GET("/", getSystemDate)
		status.DELETE("/", invalidRequest)
		status.POST("/", invalidRequest)
		status.PATCH("/", invalidRequest)
		status.PUT("/", invalidRequest)
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
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

func populateLog(c *gin.Context, status int) {
	logMessage := RequestLog{c.Request.Method, c.ClientIP(), c.FullPath(), status}
	fmt.Println(logMessage)
}
