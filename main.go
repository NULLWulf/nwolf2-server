package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/nwolf2/status", getTableStatus)
	r.GET("/nwolf2/all", getAllDocuments)
	r.GET("/nwolf2/search", searchTable)

	r.DELETE("/nwolf2/status", invalidRequest)
	r.POST("/nwolf2/status", invalidRequest)
	r.PATCH("/nwolf2/status", invalidRequest)
	r.PUT("/nwolf2/status", invalidRequest)

	r.Any("/nwolf2", notFoundResponse)
	r.Any("/", notFoundResponse)
	r.Any("/nwolf2/help", getHelp)
	err := r.Run()

	if err != nil {
		logInternalError(InternalError{"Error returned from Main gin loading", err.Error()})
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// r.Run(":8080")
}
