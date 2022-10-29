package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func notFoundResponse(c *gin.Context) {
	populateLog(c, 405)
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "404 - Resource Not found",
	})
}

func getTableStatus(c *gin.Context) {
	tableName := "Top10Cryptos"
	tableCount := 50

	populateLog(c, 200)
	c.IndentedJSON(http.StatusOK, gin.H{
		"Top10Cryptos": tableName,
		"recordCount":  tableCount,
	})
}

func invalidRequest(c *gin.Context) {
	populateLog(c, 405)
	c.IndentedJSON(http.StatusMethodNotAllowed, gin.H{
		"message": "405 - Method not allowed",
	})
}

func getALlDocuments(c *gin.Context) {

}

func searchTable(c *gin.Context) {

}
