package minihandlers

import (
	"log"
	"net/http"

	"github.com/cmfunc/jipeng/cache"
	"github.com/gin-gonic/gin"
)

// AuthMidd 身份认证中间件
func AuthMidd(ctx *gin.Context) {
	openid := ctx.GetHeader("openid")
	sessionKey := ctx.GetHeader("sessionKey")
	log.Printf("AuthMidd opeind:%s sessionkey:%s", openid, sessionKey)
	if openid == "" || sessionKey == "" {
		ctx.JSON(http.StatusBadRequest, NewResp(ErrAuthFailed, nil))
		ctx.Abort()
		return
	}
	cacheSessionKey, err := cache.GetUserSession(ctx, openid)
	if err != nil {
		log.Printf("AuthMidd opeind:%s sessionkey:%s err:%s", openid, sessionKey, err)
		ctx.JSON(http.StatusBadRequest, NewResp(ErrAuthFailed, nil))
		ctx.Abort()
		return
	}
	log.Printf("AuthMidd opeind:%s sessionkey:%s cacheSessionKey:%s", openid, sessionKey, cacheSessionKey)
	if cacheSessionKey != sessionKey {
		ctx.JSON(http.StatusBadRequest, NewResp(ErrAuthFailed, nil))
		ctx.Abort()
		return
	}

	// 设置ctx中openid
	ctx.Set("openid", openid)
}
