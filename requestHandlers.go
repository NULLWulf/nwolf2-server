package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"time"
)

func notFoundResponse(c *gin.Context) {
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"status":        "404 - Resource Not found",
		"help endpoint": "/nwolf2/help",
	})
	logRequest(c, http.StatusNotFound)
}

func getTableStatus(c *gin.Context) {
	tableName := "Top10Cryptos"
	count, err := getDocumentCount()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status": len("500 - Internal Server Error"),
		})
		logRequest(c, http.StatusInternalServerError)
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{
			"table":         tableName,
			"recordCount":   count,
			"utc-timeblock": time.Now().UTC().Format("01-02-2006-15"),
			"utc-time":      time.Now().UTC().Format(time.RFC822Z),
		})
		logRequest(c, http.StatusOK)
	}
}

func getAllDocuments(c *gin.Context) {
	docs, err := getAllDynamoDBDocs()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status":        "500 - Internal Server Error",
			"utc-timeblock": time.Now().UTC().Format("01-02-2006-15"),
		})
		logRequest(c, http.StatusInternalServerError)
	} else {
		c.JSONP(http.StatusOK, docs)
		logRequest(c, http.StatusOK)
	}
}

func searchTable(c *gin.Context) {
	s, sb := c.GetQuery("start")
	e, eb := c.GetQuery("end")
	if eb == false || sb == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"status":         "400 - Bad Request",
			"info":           "Must include both a start and end query parameter",
			"start included": sb,
			"end included":   eb,
			"help endpoint":  "/nwolf2/help",
		})
		logRequest(c, http.StatusBadRequest)
	} else {
		dReg, _ := regexp.Compile("[0-1][0-9]-[0-3][0-9]-\\d{4}-[0-2][0-9]")
		if dReg.Match([]byte(s)) && dReg.Match([]byte(e)) {
			docs, err := getDocDateRange(s, e)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{
					"status": "500 - Internal Server Error",
				})
				logRequest(c, http.StatusInternalServerError)
			} else {
				c.JSONP(http.StatusOK, docs)
				logRequest(c, http.StatusOK)
			}
		} else {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"status": "400 - Bad Request",
				"info": "Malformed Query Parameters, parameter must follow format of MM-DD-YYYY-HH in UTC Time. " +
					"See help endpoint for more information.",
				"help endpoint": "/nwolf2/help",
			})
			logRequest(c, http.StatusBadRequest)
		}
	}
}

func invalidRequest(c *gin.Context) {
	logRequest(c, http.StatusMethodNotAllowed)
	c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{
		"status":        "405 - Method not allowed",
		"help endpoint": "/nwolf2/help"})
}

func getHelp(c *gin.Context) {
	logRequest(c, http.StatusOK)
	c.IndentedJSON(http.StatusOK, gin.H{
		"info":           "The following are valid endpoints and their documentation for the Coin Marketplace Pro Top 10 API",
		"details":        "The Coin Marketplace Pro API is polled hourly and stored in UTC time.",
		"/nwolf2/status": "Returns table name as well as total number of records in database.",
		"/nwolf2/all":    "Returns all records from database",
		"/nwolf2/search?start=MM-DD-YYYY-HH&end=MM-DD-YYYY-HH": "Returns documents within and including specified date range including hour of day." +
			"HH must be between 0 and 23.  Ex. 10-28-2022-00, 10-29-2022-05, 10-29-2022-15",
		"utc-timeblock": time.Now().UTC().Format("01-02-2006-15"),
		"utc-time":      time.Now().UTC().Format(time.RFC822Z),
	})
}
