package user

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", GetAllUsers)      // Fetch all users
		userRoutes.GET("/:id", GetUserByID)   // Fetch a specific user by ID
		userRoutes.POST("/", CreateUser)      // Create a new user
		userRoutes.PUT("/:id", UpdateUser)    // Update an existing user by ID
		userRoutes.DELETE("/:id", DeleteUser) // Delete a user by ID
	}
}
