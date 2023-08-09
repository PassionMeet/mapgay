package router

import (
	"github.com/cmfunc/jipeng/handlers/middleware"
	minihandlers "github.com/cmfunc/jipeng/handlers/mini_handlers"
	"github.com/gin-gonic/gin"
)

func miniRouter(g *gin.RouterGroup) {
	//login
	g.POST("/login", minihandlers.Login)

	//auth middleware
	g.Use(middleware.MiniAuth)

	//service fucntion api
	g.POST("/geo", minihandlers.UploadGeo)               //upload user's geo location
	g.GET("/users/geo", minihandlers.GetUsersByGeo)      //get users by geo
	g.POST("/user/info", minihandlers.UploadUserinfo)    //upload user self userinfo
	g.GET("/cos/auth", minihandlers.GetCosAuth)          //get tencent auth's key
	g.POST("/leave/message", minihandlers.LeaveAMessage) //leave a message for user
}
