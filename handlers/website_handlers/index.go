package websitehandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context)  {
	c.Redirect(http.StatusMovedPermanently,"http://avatar.cdn.sourcelist.top/index.html")
}