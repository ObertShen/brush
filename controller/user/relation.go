package user

import (
	"brush/service/user"
	"net/http"

	"strconv"

	"github.com/ObertShen/gin"
)

func getRelationEndPoint(ctx *gin.Context) {
	nodes := []*user.Node{}
	links := []*user.Link{}
	if ctx.Query("platform") == "weibo" {
		weiboID, err := strconv.ParseInt(ctx.Query("param"), 10, 0)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"code": 4001, "error": err})
			return
		}

		nodes, links, err = user.GetServiceIns().GetWeiboUserByFollower(weiboID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"code": 5001, "error": err})
			return
		}
	} else {
		var err error
		nodes, links, err = user.GetServiceIns().GetZhihuUserByFollower(ctx.Query("param"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"code": 5001, "error": err})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "info": gin.H{"nodes": nodes, "links": links}})
}
