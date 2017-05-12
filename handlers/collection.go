package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/thedanielforum/arangodb-mock/jwt"
)

const (
	typeDoc = iota + 2
	typeEdge
)

type colInfo struct {
	JournalSize       uint        `json:"journalSize"`
	ReplicationFactor uint 	      `json:"replicationFactor"`
	KeyOptions        *keyOptions `json:"keyOptions"`
	Name              string      `json:"name"               binding:"required"`
	WaitForSync       bool        `json:"waitForSync"`
	DoCompact         bool	      `json:"doCompact"`
	IsVolatile        bool	      `json:"isVolatile"`
	IsSystem          bool        `json:"isSystem"`
	Type              uint        `json:"type"`
	IndexBuckets      uint        `json:"indexBuckets"`
}

type keyOptions struct {
	AllowUserKeys bool    `json:"allowUserKeys"`
	Type          string  `json:"type"`
	Increment     int     `json:"increment"`
	Offset        int     `json:"offset"`
}

func NewCol(c *gin.Context) {
	if a := jwt.ValidateJWT(c); a == false {
		return
	}

	var json colInfo
	if c.BindJSON(&json) != nil {
		c.JSON(400, gin.H{
			"error"        : true,
			"errorMessage" : "Error Unmarshaling JSON",
			"code"         : 400,
			"errorNum"     : 600,
		})
		return
	}
	if json.Type == typeEdge {
		c.JSON(200, &gin.H{
			"id"          : "9144",
			"name"        : "testCollectionUsers",
			"waitForSync" : false,
			"isVolatile"  : false,
			"isSystem"    : false,
			"status"      : 3,
			"type"        : typeEdge,
			"error"       : false,
			"code"        : 200,
		})
		return
	}

	c.JSON(200, &gin.H{
		"id"          : "9144",
		"name"        : "testCollectionUsers",
		"waitForSync" : false,
		"isVolatile"  : false,
		"isSystem"    : false,
		"status"      : 3,
		"type"        : typeDoc,
		"error"       : false,
		"code"        : 200,
	})
	return
}