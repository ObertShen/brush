package middleware

import (
	"net/http"

	"github.com/ObertShen/gin"
)

func CORSHandler() gin.HandlerFunc {
	return checkMethod()
}

func checkMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Allow-Headers", "Content-Type")
			c.Header("Access-Control-Allow-Headers", "Authorization")
			c.AbortWithStatus(http.StatusOK)
		}
	}
}
