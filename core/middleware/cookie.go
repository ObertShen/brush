package middleware

import (
	"brush/model"
	"net/http"

	"github.com/ObertShen/gin"
)

// UserInfo 用于存储用户信息的Key
const UserInfo = "UserInfo"

// CookiesParser 校验 cookies 的中间件
func CookiesParser() gin.HandlerFunc {
	return func(c *gin.Context) {
		value, err := c.Cookie("bigdata_uid")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no cookies"})
			c.Abort()
			return
		}

		userData := &model.User{UnionID: value}
		has, err := model.GetUserDataIns().Get(userData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		if !has {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid cookies"})
			c.Abort()
			return
		}

		c.Set(UserInfo, userData)
	}
}
