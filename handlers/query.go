package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/thedanielforum/arangodb-mock/jwt"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"github.com/thedanielforum/arangodb-mock/redirects"
	"github.com/apex/log"
)

type query struct {
	Count       bool     `json:"count"`
	BatchSize   uint     `json:"batchSize"`
	Cache       bool     `json:"cache"`
	MemoryLimit uint     `json:"memoryLimit"`
	Ttl         uint     `json:"ttl"`
	Query       string   `json:"query" binding:"required"`
	Options     *options `json:"options"`
}

type options struct {
	Profile           bool   `json:"profile"`
	Optimizer         string `json:"optimizer"`
	SatelliteSyncWait uint   `json:"satelliteSyncWait"`
	FullCount         bool   `json:"fullCount"`
	MaxPlans          uint   `json:"maxPlans"`
}

func Query(c *gin.Context) {
	if a := jwt.ValidateJWT(c); a == false {
		return
	}

	var rawJson query
	if c.BindJSON(&rawJson) != nil {
		c.JSON(400, gin.H{
			"error": true,
			"errorMessage": "Error Unmarshaling JSON",
			"code": 400,
			"errorNum": 400,
		})
		return
	}
	if redirects.GetConfigPath() == "" {
		log.Error("mount at least 1 json file before continuing")
		c.JSON(400, gin.H{
			"error": true,
			"errorMessage": "json file not mounted",
			"code": 400,
			"errorNum": 1337,
		})
		return
	}

	raw, err := ioutil.ReadFile(redirects.GetConfigPath())
	if err != nil {
		panic(err)
	}

	var jsonMap map[string]*json.RawMessage
	if err := json.Unmarshal(raw, &jsonMap); err !=nil{
		fmt.Println("error unmarshalleing")
	}


	mapResult := make(map[string]interface{})
	for value := range jsonMap {
		var anon interface{}
		err = json.Unmarshal(*jsonMap[value], &anon)
		if err != nil {
			fmt.Println("error unmarshalleing")
		}
		mapResult[value] = anon
	}

	var warnings []gin.H
	var results []gin.H
	results = append(results,mapResult)

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
