package main

import (
	"log"

	"filllabs-task/user-management-system/backend/api/user"
	"filllabs-task/user-management-system/backend/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	database, err := db.InitDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close() // Ensure the database is closed when the application shuts down

	// Initialize the Gin framework
	router := gin.Default()

	// Add CORS middleware for handling cross-origin requests
	router.Use(cors.Default()) // Default CORS settings

	// Register user-related routes
	user.RegisterRoutes(router)

	// Start the server
	log.Println("Server running at http://localhost:8080")
	router.Run(":8080")
}
