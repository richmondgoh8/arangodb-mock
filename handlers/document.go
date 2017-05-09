package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/thedanielforum/arangodb-mock/jwt"
)

type docEdge struct {
	Type uint `json:"type"`
	To string `json:"_to"`
	From string `json:"_from"`
}

func NewDocument(c *gin.Context) {
	if a := jwt.ValidateJWT(c); a == false {
		return
	}

	var json docEdge
	if c.BindJSON(&json) != nil {
		c.JSON(400, gin.H{
			"error"        : true,
			"errorMessage" : "Error Unmarshaling JSON, Invalid document type",
			"code"         : 400,
			"errorNum"     : 1227,
		})
		return
	}
	c.JSON(200, gin.H{
		"_id": "testCol/81622",
		"_key": "81622",
		"_rev": "_U8rEwqa---",
	})
	return
}

func DeleteDocument(c *gin.Context) {
	if a := jwt.ValidateJWT(c); a == false {
		return
	}

	c.JSON(200, gin.H{
		"_id": "testCol/81622",
		"_key": "81622",
		"_rev": "_U8rEwqa---",
	})
}