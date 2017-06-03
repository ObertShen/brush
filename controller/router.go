package controller

import (
	"brush/controller/user"
	"brush/core/middleware"

	"github.com/ObertShen/gin"
)

// Mount 加载所有模块对应的路由和请求处理
func Mount() (r *gin.Engine) {
	r = gin.New()
	// middleware.AuthRequired()
	r.Use(gin.Recovery(), middleware.Logger(), middleware.CORSHandler())

	mainRouter := r.Group("/bigdata")
	user.AppendTo(mainRouter)

	return r
}
