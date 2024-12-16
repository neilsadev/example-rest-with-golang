package controller

import (
	"neilsadev/todo-api/database" // This helps us talk to our database where we keep our user information.
	"neilsadev/todo-api/models"   // This helps us understand what a user looks like.
	"net/http"                    // This helps us send messages over the internet.

	"github.com/gin-gonic/gin"    // This is a tool that helps us make our web server.
)

// GetUsers shows all the users we have.
func GetUsers(c *gin.Context) {
	// We have a big box called users to keep all our users.
	var users []models.User
	// We look in our database and find all the users, including their profiles and to-do lists with tasks and tags.
	database.DB.Preload("Profile").Preload("TodoLists.Tasks.Tags").Find(&users)
	// We show all the users we found to whoever asked for them.
	c.JSON(http.StatusOK, users)
}

// GetUserByID shows a specific user by their ID.
func GetUserByID(c *gin.Context) {
	// We have a special ID, like a name tag, to find a specific user.
	id := c.Param("id")
	var user models.User
	// We search in our database for the user with that special ID.
	if err := database.DB.Preload("Profile").Preload("TodoLists.Tasks.Tags").First(&user, id).Error; err != nil {
		// If we can't find it, we say "User not found."
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	// If we find it, we show the user to whoever asked for it.
	c.JSON(http.StatusOK, user)
}

// CreateUser makes a new user.
func CreateUser(c *gin.Context) {
	// We want to make a new user, like creating a new character in a game.
	var user models.User
	// We try to understand what the new user should look like from the request.
	if err := c.ShouldBindJSON(&user); err != nil {
		// If we can't understand, we say there's a mistake.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// We add the new user to our database, like adding a new character to our game.
	database.DB.Create(&user)
	// We show the new user to whoever asked for it.
	c.JSON(http.StatusCreated, user)
}

// UpdateUser changes an existing user.
func UpdateUser(c *gin.Context) {
	// We have a special ID to find the user we want to change.
	id := c.Param("id")
	var user models.User
	// We look for the user in our database.
	if err := database.DB.First(&user, id).Error; err != nil {
		// If we can't find it, we say "User not found."
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	// We try to understand what the changes should be from the request.
	if err := c.ShouldBindJSON(&user); err != nil {
		// If we can't understand, we say there's a mistake.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// We save the changes to our database, like updating the character in our game.
	database.DB.Save(&user)
	// We show the updated user to whoever asked for it.
	c.JSON(http.StatusOK, user)
}

// DeleteUser removes a user.
func DeleteUser(c *gin.Context) {
	// We have a special ID to find the user we want to remove.
	id := c.Param("id")
	// We try to delete the user from our database.
	if err := database.DB.Delete(&models.User{}, id).Error; err != nil {
		// If we can't find it, we say "User not found."
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	// We say that the user was successfully deleted.
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
