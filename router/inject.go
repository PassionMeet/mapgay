package router

import (
	"github.com/cmfunc/jipeng/handlers/middleware"
	"github.com/gin-gonic/gin"
)

func Inject(srv *gin.Engine) {
	// CORS
	srv.Use(middleware.CORS())
	// inject router
	websiteRouter(srv.Group("/v1/website"))
	miniRouter(srv.Group("/v1/mini"))
	cmsRouter(srv.Group("/v1/ums"))
}
