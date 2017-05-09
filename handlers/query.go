package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/thedanielforum/arangodb-mock/jwt"
)

type query struct {
	Query string `json:"query" binding:"required"`
}

func Query(c *gin.Context) {
	if a := jwt.ValidateJWT(c); a == false {
		return
	}

	var json query
	if c.BindJSON(&json) != nil {
		c.JSON(400, gin.H{
			"error": true,
			"errorMessage": "Error Unmarshaling JSON",
			"code": 400,
			"errorNum": 400,
		})
		return
	}

	var results []gin.H
	results = append(results, gin.H{
		"_key": "81465",
		"_id": "testCol/81465",
		"_rev": "_U8rCaHS---",
		"name": "testCol5",
		"owner": "12345",
	})

	var warnings []gin.H
	c.JSON(200, gin.H{
		"result": results,
		"hasMore": false,
		"cached": false,
		"extra": gin.H{
			"stats": gin.H{
				"writesExecuted": 0,
				"writesIgnored": 0,
				"scannedFull": 1,
				"scannedIndex": 0,
				"filtered": 0,
				"executionTime": 0.00010895729064941406,
			},
			"warnings": warnings,
		},
		"error": false,
		"code": 201,
	})
}
