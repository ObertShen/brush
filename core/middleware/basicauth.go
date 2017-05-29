package middleware

import (
	"fmt"
	"net/http"

	"github.com/ObertShen/gin"
)

const (
	clientKey = "2wqwe60e66ce67631gasdqw2"
)

// AuthRequired 校验HTTP Basic Authorization的中间件
func AuthRequired() gin.HandlerFunc {
	return checkBasicAuth()
}

func checkBasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userName, password, ok := c.Request.BasicAuth()
		if !ok || password != "" || !(userName == clientKey) {
			fmt.Println("=========================")
			c.JSON(http.StatusUnauthorized, gin.H{})
			c.Abort()
		}
	}
}
