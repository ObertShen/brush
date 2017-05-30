package user

import (
	"net/http"

	"brush/model"
	"brush/util"

	"github.com/ObertShen/gin"
)

func loginEndPoint(ctx *gin.Context) {
	var loginInfo struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.Bind(&loginInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 4002, "error": err})
		return
	}

	userData := &model.User{Name: loginInfo.UserName}
	has, err := model.GetUserDataIns().Get(userData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 5000, "error": err})
		return
	}

	if !has || util.GetSHA256String(loginInfo.Password) != userData.Password {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 4001, "error": "invalid username or password"})
		return
	}

	ctx.SetCookie(
		"bigdata_uid",
		userData.UnionID,
		3600,
		"/bigdata",
		"localhost",
		false,
		true,
	)

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "info": userData})
}

func getUserEndPoint(ctx *gin.Context) {
	userInfo, exist := ctx.Get(UserInfo)
	if !exist {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 5001, "error": "no cookie middleware"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "info": userInfo})
}

// AppendTo 加载 user 模块的路由
func AppendTo(r *gin.RouterGroup) {
	authorized := r.Group("/users")

	authorized.GET("/", CookiesParser(), findUserEndPoint)
	authorized.POST("/login", loginEndPoint)
	authorized.GET("/relation", getRelationEndPoint)
	authorized.GET("/me", CookiesParser(), getUserEndPoint)
}
