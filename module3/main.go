package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, &gin.H{
			"pong": "module3",
		})
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8083")
}
