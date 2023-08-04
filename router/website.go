package router

import (
	websitehandlers "github.com/cmfunc/jipeng/handlers/website_handlers"
	"github.com/gin-gonic/gin"
)

func websiteRouter(g *gin.RouterGroup) {
	g.GET("", websitehandlers.GetIndex) //offcial website
}
