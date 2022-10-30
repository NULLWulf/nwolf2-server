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
		"status":        "404 - Resource Not found",
		"help endpoint": "/nwolf2/help",
	})
}

func getTableStatus(c *gin.Context) {
	tableName := "Top10Cryptos"
	count, err := getDocumentCount()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status": "500 - Internal Server Error"})
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
			"status": "500 - Internal Server Error"})
	} else {
		c.JSONP(http.StatusOK, docs)
	}
}

func searchTable(c *gin.Context) {
	s := c.Query("start")
	e := c.Query("end")
	docs, err := getDocDateRange(s, e)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status": "500 - Internal Server Error"})
	} else {
		c.JSONP(http.StatusOK, docs)
	}
}

func invalidRequest(c *gin.Context) {
	c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{
		"status":        "405 - Method not allowed",
		"help endpoint": "/nwolf2/help"})
}

func getHelp(c *gin.Context) {
	c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{
		"info":           "The following are valid endpoints and their documentation for the Coin Marketplace Pro Top 10 API",
		"details":        "The Coin Marketplace Pro API is polled hourly and stored in UTC time.",
		"/nwolf2/status": "Returns table name as well as total number of records in database.",
		"/nwolf2/all":    "Returns all records from database",
		"/nwolf2/search?start=<MM-DD-YYYY-HH>&end=<MM-DD-YYYY-HH>": "Returns documents within and including specified date range." +
			"HH must be between 0 and 23.  Ex. 10-28-2022-00, 10-29-2022-05, 10-29-2022-15",
	})
}
