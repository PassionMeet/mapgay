package handlers

import (
	"net/http"

	"github.com/cmfunc/jipeng/cache"
	"github.com/gin-gonic/gin"
)

// AuthMidd 身份认证中间件
func AuthMidd(ctx *gin.Context) {
	userid := ctx.GetHeader("userid")
	sessionKey := ctx.GetHeader("sessionKey")
	if userid == "" || sessionKey == "" {
		ctx.JSON(http.StatusBadRequest, nil)
		ctx.Abort()
		return
	}
	cacheSessionKey, err := cache.GetUserSession(ctx, userid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		ctx.Abort()
		return
	}

	if cacheSessionKey != sessionKey {
		ctx.JSON(http.StatusBadRequest, nil)
		ctx.Abort()
		return
	}
}
