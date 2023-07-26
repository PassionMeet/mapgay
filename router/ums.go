package router

import (
	umshandlers "github.com/cmfunc/jipeng/handlers/ums_handlers"
	"github.com/gin-gonic/gin"
)

// user manager system router
func cmsRouter(g *gin.RouterGroup) {
	g.POST("/login", umshandlers.Login)                 //cms login api
	g.Use(umshandlers.Auth)                             //cms auth middleware
	g.POST("/stronghold", umshandlers.InsertStronghold) //insert a stronghold into sys
}
