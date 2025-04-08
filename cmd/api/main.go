package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	// Set Gin mode
	gin.SetMode(gin.ReleaseMode)

	// Initialize router
	router = gin.Default()

	// Basic route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

// Handler is the main entry point for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
