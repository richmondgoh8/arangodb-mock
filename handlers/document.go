package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/thedanielforum/arangodb-mock/jwt"
	"github.com/apex/log"
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
		log.Info(c.BindJSON(&json).Error())
		c.JSON(400, gin.H{
			"error"        : true,
			"errorMessage" : "Error Unmarshaling JSON, Invalid document type",
			"code"         : 400,
			"errorNum"     : 1227,
		})
		return
	}

	if json.To != "" && json.From != "" {
		c.JSON(200, gin.H{
			"_id": "testEdge/31432",
			"_key": "31432",
			"_rev": "_U8rEwqa---",
		})
		return
	}

	if json.To == "" && json.From == "" {
		c.JSON(200, gin.H{
			"_id": "testCol/81622",
			"_key": "81622",
			"_rev": "_U8rEwqa---",
		})
		return
	}
	c.JSON(400, gin.H{
		"error": true,
		"errorMessage": "invalid edge attribute",
		"code": 400,
		"errorNum": 1233,
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