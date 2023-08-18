package router

import (
	"github.com/cmfunc/jipeng/handlers/middleware"
	umshandlers "github.com/cmfunc/jipeng/handlers/ums_handlers"
	"github.com/gin-gonic/gin"
)

// user manager system router
func cmsRouter(g *gin.RouterGroup) {
	//cms login api
	g.POST("/login", umshandlers.Login)

	//cms auth middleware
	g.Use(middleware.UmsAuth)

	//service fucntions api
	g.POST("/stronghold", umshandlers.InsertStronghold)   //insert a stronghold into sys
	g.GET("/strongholds", umshandlers.GetStrongholds)     //query strongsholds list page by page
	g.GET("/stronghold", umshandlers.GetStrongholdDetail) //query a stronghlod detail
	g.DELETE("/stronghold", umshandlers.DelStronghold)    //delete a stronghlod detail
	g.PUT("/stronghold", umshandlers.UpdateStronghold)    //edit a stronghlod detail

	//leave message
	g.GET("/messages", umshandlers.GetMessages) //query messages list page by page.
}
