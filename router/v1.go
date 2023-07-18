package router

import (
	"github.com/cmfunc/jipeng/handlers"
	"github.com/gin-gonic/gin"
)

func v1Router(g *gin.RouterGroup) {
	g.POST("/login", handlers.Login)              //login
	g.Use(handlers.AuthMidd)                      //auth middleware
	g.POST("/geo", handlers.UploadGeo)            //upload user's geo location
	g.GET("/users/geo", handlers.GetUsersByGeo)   //get users by geo
	g.POST("/user/info", handlers.UploadUserinfo) //upload user self userinfo
	g.GET("/cos/auth", handlers.GetCosAuth)       //get tencent auth's key
}
