package middleware

import (
	"net/http"

	util "muse/core/utils/http"

	"github.com/ObertShen/gin"
)

func CORSHandler() gin.HandlerFunc {
	return checkMethod()
}

func checkMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			util.SetOPTIONHeader(c.Writer)
			c.AbortWithStatus(http.StatusOK)
		}
	}
}
