package arangodb_mock

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/apex/log"
	"github.com/thedanielforum/arangodb-mock/handlers"
	"net/http"
)

func Start(port int, debugMode bool) {
	if port <= 0 {
		log.Fatal("port is required")
	}

	gin.SetMode(gin.ReleaseMode)
	if debugMode {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.Default()
	r.POST("_open/auth", handlers.Auth)
	r.POST("_db/:db/_api/collection", handlers.NewCol)
	r.POST("_db/:db/_api/document/:collection",handlers.NewDocument)
	r.POST("_db/:db/_api/cursor", handlers.Query)
	r.DELETE("_db/:db/_api/document/:collection/:key", handlers.DeleteDocument)
	// Error pages
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": true,
			"errorNum": 1228,
			"errorMessage": "404 not found",
			"code": 404,
		})
	})

	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.WithError(err).Fatal("server failed to start")
	}
}