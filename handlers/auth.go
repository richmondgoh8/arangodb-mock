package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/thedanielforum/arangodb-mock/jwt"
)

type credentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type jwtCredentials struct {
	jwt            string `json:"jwt"`
	mustChangePass bool   `json:"must_change_pass"`
}

func Auth(c *gin.Context) {
	var json credentials
	if c.BindJSON(&json) != nil {
		c.JSON(401, gin.H{
			"error": true,
			"errorMessage": "Error Unmarshaling JSON",
			"code": 401,
			"errorNum": 401,
		})
		return
	}

	if json.Username == "" || json.Password == "" {
		c.JSON(401, gin.H{
			"error": true,
			"errorMessage": "Wrong Credentials",
			"code": 401,
			"errorNum": 401,
		})
		return
	}
	c.JSON(200, gin.H{
		"jwt": jwt.GetJWT(),
		"mustChangePass": false,
	})
	return
}