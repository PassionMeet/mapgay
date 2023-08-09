package middleware

import (
	"log"
	"net/http"

	"github.com/cmfunc/jipeng/cache"
	"github.com/cmfunc/jipeng/handlers/resp"
	"github.com/gin-gonic/gin"
)

// MiniAuth 身份认证中间件
func MiniAuth(ctx *gin.Context) {
	openid := ctx.GetHeader("openid")
	sessionKey := ctx.GetHeader("sessionKey")
	log.Printf("AuthMidd opeind:%s sessionkey:%s", openid, sessionKey)
	if openid == "" || sessionKey == "" {
		ctx.JSON(http.StatusBadRequest, resp.NewResp(resp.ErrAuthFailed, nil))
		ctx.Abort()
		return
	}
	cacheSessionKey, err := cache.GetUserSession(ctx, openid)
	if err != nil {
		log.Printf("AuthMidd opeind:%s sessionkey:%s err:%s", openid, sessionKey, err)
		ctx.JSON(http.StatusBadRequest, resp.NewResp(resp.ErrAuthFailed, nil))
		ctx.Abort()
		return
	}
	log.Printf("AuthMidd opeind:%s sessionkey:%s cacheSessionKey:%s", openid, sessionKey, cacheSessionKey)
	if cacheSessionKey != sessionKey {
		ctx.JSON(http.StatusBadRequest, resp.NewResp(resp.ErrAuthFailed, nil))
		ctx.Abort()
		return
	}

	// 设置ctx中openid
	ctx.Set("openid", openid)
}
