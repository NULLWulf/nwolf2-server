package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	getAllDynamoDBDocs()
	// Loads Environmental variables into program
	// e.g AWS, Loggly CMP token.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()

	//r.GET("/nwolf2/status", getTableStatus)
	r.GET("/nwolf2/all", getAllDocuments)
	r.GET("/nwolf2/search", searchTable)

	r.DELETE("/nwolf2/status", invalidRequest)
	r.POST("/nwolf2/status", invalidRequest)
	r.PATCH("/nwolf2/status", invalidRequest)
	r.PUT("/nwolf2/status", invalidRequest)

	r.Any("/nwolf2", notFoundResponse)
	r.Any("/", notFoundResponse)
	err = r.Run()

	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
