package handlers

import (
	"net/http"

	"github.com/cmfunc/jipeng/cache"
	"github.com/gin-gonic/gin"
)

// AuthMidd 身份认证中间件
func AuthMidd(ctx *gin.Context) {
	openid := ctx.GetHeader("openid")
	sessionKey := ctx.GetHeader("sessionKey")
	if openid == "" || sessionKey == "" {
		ctx.JSON(http.StatusBadRequest, NewResp(ErrAuthFailed, nil))
		ctx.Abort()
		return
	}
	cacheSessionKey, err := cache.GetUserSession(ctx, openid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewResp(ErrAuthFailed, nil))
		ctx.Abort()
		return
	}

	if cacheSessionKey != sessionKey {
		ctx.JSON(http.StatusBadRequest, NewResp(ErrAuthFailed, nil))
		ctx.Abort()
		return
	}

	// 设置ctx中openid
	ctx.Set("openid", openid)
}
