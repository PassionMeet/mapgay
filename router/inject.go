package router

import "github.com/gin-gonic/gin"

func Inject(srv *gin.Engine)  {
	v1Router(srv.Group("/v1"))
}