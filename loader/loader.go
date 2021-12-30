package loader

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Loader(r *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("authorization")
	r.Use(cors.New(config))

}
