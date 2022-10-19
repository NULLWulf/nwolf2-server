package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nullwulf/loggly"
	"log"
	"net/http"
	"os"
	"time"
)

type RequestLog struct {
	Method      string
	SourceIP    string
	RequestPath string
	StatusCode  int
}

func main() {

	// Loads Environmental variables into program
	// e.g AWS, Loggly CMP token.
	err := godotenv.Load()
	// If detects an error loading .env file terminates program
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Instantiate Loggly Client
	r := gin.Default()

	r.GET("/nwolf2/status", getSystemDate)
	r.DELETE("/nwolf2/status", invalidRequest)
	r.POST("/nwolf2/status", invalidRequest)
	r.PATCH("/nwolf2/status", invalidRequest)
	r.PUT("/nwolf2/status", invalidRequest)
	r.Any("/nwolf2", notFoundResponse)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

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

func populateLog(c *gin.Context, status int) {
	// Attempts to get APP_TAG from environment variables file.
	tag := os.Getenv("APP_TAG")
	if tag == "" {
		tag = "default-Nwolf2-Server"
	}
	// Instantiate Loggly Client
	lgglyClient := loggly.New(tag)

	logMessage := RequestLog{c.Request.Method, c.ClientIP(), c.FullPath(), status}
	json, err := json.Marshal(&logMessage)
	if err != nil {
		lgglyClient.EchoSend("error", err.Error())

	}
	lgglyClient.EchoSend("info", string(json))
}
