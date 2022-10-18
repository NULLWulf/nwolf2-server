package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/nwolf2/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"system time": time.Now().UTC().Format(time.RFC3339),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
