package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mangugeng/gatcor-backend/config"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	_ = config.InitDB() // We'll use this later

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

	// Get port from env
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// Run server
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
