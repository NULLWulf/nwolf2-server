package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/nullwulf/loggly"
	"os"
)

type RequestLog struct {
	Method      string
	SourceIP    string
	RequestPath string
	StatusCode  int
}

func logRequest(c *gin.Context, status int) {
	// Instantiate Loggly Client
	lgglyClient := loggly.New(getTag())

	logMessage := RequestLog{c.Request.Method, c.ClientIP(), c.FullPath(), status}
	jsonMsg, err := json.Marshal(&logMessage)
	if err != nil {
		lgglyClient.EchoSend("error", err.Error())
	}
	lgglyClient.EchoSend("info", string(jsonMsg))
}

func getTag() string {
	tag := os.Getenv("APP_TAG")
	if tag == "" {
		tag = "default-Nwolf2-Server"
	}
	return tag
}
