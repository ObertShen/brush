package user

import (
	"net/http"

	"brush/service/user"

	"github.com/ObertShen/gin"
)

func findUserEndPoint(ctx *gin.Context) {
	if ctx.Query("key") == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 4001, "error": "no parameters"})
		return
	}

	weiboUsers, err := user.GetServiceIns().GetWeiboUserListByName(ctx.Query("key"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 5000, "error": err})
		return
	}

	zhihuUsers, err := user.GetServiceIns().GetZhihuUserListByName(ctx.Query("key"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 5000, "error": err})
		return
	}

	result := make([]interface{}, 0)
	if len(weiboUsers) != 0 {
		result = append(result, weiboUsers)
	}

	if len(zhihuUsers) != 0 {
		result = append(result, zhihuUsers)
	}

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "info": result})
}
