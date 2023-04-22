package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var player = struct {
	Name string `json:"name"`
	City string `json:"city"`
}{
	Name: "Paul Scholes",
	City: "Manchester",
}

func playerJson(c *gin.Context) {

	c.JSON(http.StatusOK, player)
}

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/**")
	router.GET("/api/mats", playerJson)
	router.Run(":8080")

}
