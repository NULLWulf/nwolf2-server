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

type InternalError struct {
	Context string
	Error   string
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

func logInternalError(internalErr InternalError) {
	// Instantiate Loggly Client
	lgglyClient := loggly.New(getTag())
	jsonMsg, err := json.Marshal(&internalErr)
	if err != nil {
		lgglyClient.EchoSend("error", err.Error())
	}
	lgglyClient.EchoSend("error", string(jsonMsg))
}
