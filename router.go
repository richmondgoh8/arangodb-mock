package arangodb_mock

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/apex/log"
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


	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.WithError(err).Fatal("server failed to start")
	}
}
