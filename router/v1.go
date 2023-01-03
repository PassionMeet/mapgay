package router

import (
	"github.com/cmfunc/jipeng/handlers"
	"github.com/gin-gonic/gin"
)

func v1Router(g *gin.RouterGroup) {
	g.POST("/login", handlers.Login)   //login
	g.POST("/geo", handlers.UploadGeo) //upload user's geo location
}
