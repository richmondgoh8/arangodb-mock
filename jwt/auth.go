package jwt

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

const jwt = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTY4ODc4MDksImlhdCI6MS40OTQyOTU4MDkyNzQ4MTFlKzYsImlzcyI6ImFyYW5nb2RiIiwicHJlZmVycmVkX3VzZXJuYW1lIjoicm9vdCJ9.vELolGpRMVONw3B5BkQmZLBGsp-Va6a-D391fTy_e0Q="

func GetJWT() string {
	return jwt
}

func ValidateJWT(c *gin.Context) bool {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader != fmt.Sprintf("bearer %s", jwt) {
		c.JSON(401, gin.H{
			"error": true,
			"errorMessage": "Wrong Credentials",
			"code": 401,
			"errorNum": 401,
		})
		return false
	}
	return true
}
