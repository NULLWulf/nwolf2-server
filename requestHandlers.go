package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DateQuery struct {
	Date string `form:"date"`
}

func notFoundResponse(c *gin.Context) {
	logRequest(c, 405)
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"status": "404 - Resource Not found",
	})
}

func getTableStatus(c *gin.Context) {
	tableName := "Top10Cryptos"
	count, err := getDocumentCount()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status": "500 - Internal Server Error",
			"error":  err.Error(),
		})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{
			"table":       tableName,
			"recordCount": count,
		})
	}
}

func getAllDocuments(c *gin.Context) {
	docs, err := getAllDynamoDBDocs()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status": "500 - Internal Server Error",
			"error":  err.Error(),
		})
	} else {
		c.JSONP(http.StatusOK, docs)
	}
}

func searchTable(c *gin.Context) {
	cs := c.Query("date")
	c.IndentedJSON(http.StatusUnauthorized, gin.H{
		"qs": cs,
	})
}

func invalidRequest(c *gin.Context) {
	logRequest(c, 405)
	c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{
		"status": "405 - Method not allowed",
	})
}
