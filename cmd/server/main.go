package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spokencafe/backend/internal/config"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to MongoDB (optional for development)
	mongoClient := config.ConnectDataBase(cfg.DatabaseURL)
	if mongoClient != nil {
		defer func() {
			if err := mongoClient.Disconnect(nil); err != nil {
				fmt.Printf("Error disconnecting from MongoDB: %v\n", err)
			}
		}()
	} else {
		fmt.Println("Running without database connection (development mode)")
	}

	// Set Gin mode
	gin.SetMode(gin.ReleaseMode)

	// Initialize Gin router
	router := gin.Default()

	// Print server running message with port from config
	fmt.Printf("Server is running on port %s\n", cfg.Port)

	// Simple health check route
	router.GET("/", func(c *gin.Context) {
		status := "connected"
		if mongoClient == nil {
			status = "disconnected"
		}
		c.JSON(http.StatusOK, gin.H{
			"message":  "Spoken Cafe Backend is running!",
			"database": status,
			"port":     cfg.Port,
		})
	})

	// Run the server on the port from config
	if err := router.Run(":" + cfg.Port); err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
