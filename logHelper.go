package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/nullwulf/loggly"
	"os"
)

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
