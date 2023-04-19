package main

import (
	persons "GC/apis"
	ginning "GC/gin"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", ginning.Home)
	r.GET("/api/david", persons.David)
	r.Run(":8080")


}