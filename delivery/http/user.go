package http

import "github.com/gin-gonic/gin"

func Login(server *gin.Engine) {
	server.GET("/login")
}
