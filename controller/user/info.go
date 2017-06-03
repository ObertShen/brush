package user

import (
	"fmt"
	"net/http"

	"brush/model"
	"brush/service/user"

	"strconv"

	"github.com/ObertShen/gin"
)

func getZhihuAndWeiboUsers(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	pageNo, _ := strconv.Atoi(ctx.Query("pageNo"))

	fmt.Println(pageSize)
	fmt.Println(pageNo)

	if ctx.Query("platform") == "zhihu" {
		zhihuUsers, err := user.GetServiceIns().GetZhihuUsers(model.ZhihuUser{}, pageSize, pageNo)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"code": 5000, "error": err})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "info": zhihuUsers})
	} else {
		weiboUsers, err := user.GetServiceIns().GetWeiboUsers(model.WeiboUser{}, pageSize, pageNo)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"code": 5000, "error": err})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "info": weiboUsers})
	}
}

func findUserEndPoint(ctx *gin.Context) {
	if ctx.Query("key") == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 4001, "error": "no parameters"})
		return
	}

	weiboUsers, err := user.GetServiceIns().GetWeiboUsers(model.WeiboUser{NickName: ctx.Query("key")}, 10, 0)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 5000, "error": err})
		return
	}

	zhihuUsers, err := user.GetServiceIns().GetZhihuUsers(model.ZhihuUser{NickName: ctx.Query("key")}, 10, 0)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 5000, "error": err})
		return
	}

	result := make([]interface{}, 0)
	if len(weiboUsers) != 0 {
		for _, weiboUser := range weiboUsers {
			result = append(result, weiboUser)
		}
	}

	if len(zhihuUsers) != 0 {
		for _, zhihuUser := range zhihuUsers {
			result = append(result, zhihuUser)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "info": result})
}
