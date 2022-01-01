package main

import (
	"os"

	"github.com/cryptbrn/cryptbrn-backend/delivery/http"
	"github.com/cryptbrn/cryptbrn-backend/loader"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	loader.Loader(r)

	http.Ping(r)

	r.Run(":" + os.Getenv("PORT"))
}
