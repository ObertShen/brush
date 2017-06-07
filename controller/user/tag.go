package user

import (
	"net/http"
	"strconv"

	"brush/service/user"

	"github.com/ObertShen/gin"
)

func getUserTagEndPoint(ctx *gin.Context) {
	if ctx.Query("platform") == "weibo" {
		weiboID, err := strconv.ParseInt(ctx.Query("params"), 10, 0)
		if err != nil || weiboID < 1 {
			ctx.JSON(http.StatusBadRequest, gin.H{"code": 4000, "error": "invalid params"})
			return
		}

		userTag, err := user.GetServiceIns().GetWeiboUserTags(weiboID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"code": 5000, "error": err})
			return
		}

		if userTag == nil {
			ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "info": []string{}})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "info": userTag.Tags})
	} else {
		userTag, err := user.GetServiceIns().GetZhihuUserTags(ctx.Query("params"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"code": 5000, "error": err})
			return
		}

		if userTag == nil {
			ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "info": []string{}})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "info": userTag.Tags})
	}
}
