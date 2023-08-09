package middleware

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins: false,
		AllowOrigins:    []string{"http://localhost:5173", "http://sourcelist.top"},
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://sourcelist.top"
		},
		AllowMethods:           []string{http.MethodPost},
		AllowHeaders:           []string{"Origin","Content-Type"},
		AllowCredentials:       true,
		ExposeHeaders:          []string{"Content-Length"},
		MaxAge:                 0,
		AllowWildcard:          false,
		AllowBrowserExtensions: false,
		AllowWebSockets:        false,
		AllowFiles:             false,
	})
}
