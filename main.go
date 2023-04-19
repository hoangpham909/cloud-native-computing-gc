package main

import (
	ginning "GC/gin"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", ginning.Home)
	r.Run(":8080")


}