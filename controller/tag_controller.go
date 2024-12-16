package controller

import (
	"neilsadev/todo-api/database"
	"neilsadev/todo-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all Tags
func GetTags(c *gin.Context) {
	// Imagine we have a box full of tags, and we want to show them all.
	var tags []models.Tag
	// We look inside our database and find all the tags, including their tasks.
	database.DB.Preload("Tasks").Find(&tags)
	// We show all the tags we found to whoever asked for them.
	c.JSON(http.StatusOK, tags)
}

// Get a Tag by ID
func GetTagByID(c *gin.Context) {
	// We have a special tag ID, like a name tag, and we want to find that tag.
	id := c.Param("id")
	var tag models.Tag
	// We search in our database for the tag with that special ID.
	if err := database.DB.Preload("Tasks").First(&tag, id).Error; err != nil {
		// If we can't find it, we say "Tag not found."
		c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		return
	}
	// If we find it, we show the tag to whoever asked for it.
	c.JSON(http.StatusOK, tag)
}

// Create a Tag
func CreateTag(c *gin.Context) {
	// We want to make a new tag, like creating a new sticker.
	var tag models.Tag
	// We try to understand what the new tag should look like from the request.
	if err := c.ShouldBindJSON(&tag); err != nil {
		// If we can't understand, we say there's a mistake.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// We add the new tag to our database, like putting the sticker in our collection.
	database.DB.Create(&tag)
	// We show the new tag to whoever asked for it.
	c.JSON(http.StatusCreated, tag)
}

// Update a Tag
func UpdateTag(c *gin.Context) {
	// We have a tag ID, and we want to change that tag.
	id := c.Param("id")
	var tag models.Tag
	// We look for the tag with that ID in our database.
	if err := database.DB.First(&tag, id).Error; err != nil {
		// If we can't find it, we say "Tag not found."
		c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		return
	}
	// We try to understand the new changes for the tag from the request.
	if err := c.ShouldBindJSON(&tag); err != nil {
		// If we can't understand, we say there's a mistake.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// We save the changes to the tag in our database.
	database.DB.Save(&tag)
	// We show the updated tag to whoever asked for it.
	c.JSON(http.StatusOK, tag)
}

// Delete a Tag
func DeleteTag(c *gin.Context) {
	// We have a tag ID, and we want to remove that tag.
	id := c.Param("id")
	// We try to delete the tag with that ID from our database.
	if err := database.DB.Delete(&models.Tag{}, id).Error; err != nil {
		// If we can't find it to delete, we say "Tag not found."
		c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		return
	}
	// We say "Tag deleted" to confirm it's gone.
	c.JSON(http.StatusOK, gin.H{"message": "Tag deleted"})
}
