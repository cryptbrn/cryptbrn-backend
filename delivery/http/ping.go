package http

import "github.com/gin-gonic/gin"

func Ping(server *gin.Engine) {
	server.GET("/", func(c *gin.Context) {
		c.String(200, "Cryptbrn API")
	})
}
