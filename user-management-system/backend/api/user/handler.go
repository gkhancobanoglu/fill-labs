package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllUsers handles the request to fetch all users from the database
func GetAllUsers(c *gin.Context) {
	users, err := GetUsersFromDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUserByID handles the request to fetch a specific user by their ID
func GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id")) // Convert ID from string to integer
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := GetUserFromDB(id) // Fetch the user from the database
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser handles the request to create a new user
func CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil { // Bind JSON input to the user struct
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := SaveUserToDB(&user) // Save the new user to the database
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// UpdateUser handles the request to update an existing user
func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id")) // Convert ID from string to integer
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user User
	if err := c.ShouldBindJSON(&user); err != nil { // Bind JSON input to the user struct
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = UpdateUserInDB(id, &user) // Update the user in the database
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUser handles the request to delete a user by their ID
func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id")) // Convert ID from string to integer
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = DeleteUserFromDB(id) // Delete the user from the database
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
