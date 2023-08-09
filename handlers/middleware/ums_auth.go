package middleware

import (
	"log"
	"net/http"

	"github.com/cmfunc/jipeng/cache"
	"github.com/cmfunc/jipeng/handlers/resp"
	"github.com/gin-gonic/gin"
)

func UmsAuth(c *gin.Context) {
	token, account := c.GetHeader("token"), c.GetHeader("account")
	log.Printf("UmsAuth account:%s token:%s", account, token)
	if account == "" || token == "" {
		c.JSON(http.StatusBadRequest, resp.NewResp(resp.ErrAuthFailed, nil))
		c.Abort()
		return
	}
	cacheToken, err := cache.GetUMSToken(c, account)
	if err != nil {
		log.Printf("UmsAuth account:%s token:%s err:%s", account, token, err)
		c.JSON(http.StatusBadRequest, resp.NewResp(resp.ErrAuthFailed, nil))
		c.Abort()
		return
	}
	log.Printf("UmsAuth account:%s token:%s cacheToken:%s", account, token, cacheToken)
	if token != cacheToken {
		c.JSON(http.StatusBadRequest, resp.NewResp(resp.ErrAuthFailed, nil))
		c.Abort()
		return
	}
	// 设置ctx中openid
	c.Set("account", account)
}
