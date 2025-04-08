package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vercel/go-bridge/go/bridge"
)

func init() {
	// Set Gin mode
	gin.SetMode(gin.ReleaseMode)

	// Initialize router
	router := gin.Default()

	// Basic route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Handle all routes
	http.Handle("/", router)
}

func main() {
	bridge.Start(http.DefaultServeMux)
}
