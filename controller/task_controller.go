package controller

import (
	"neilsadev/todo-api/database"
	"neilsadev/todo-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all Tasks
func GetTasks(c *gin.Context) {
	// Imagine you have a big box full of tasks, and you want to see all of them.
	var tasks []models.Task
	// We look inside our database and find all the tasks, including their tags.
	database.DB.Preload("Tags").Find(&tasks)
	// We show all the tasks we found to whoever asked for them.
	c.JSON(http.StatusOK, tasks)
}

// Get a Task by ID
func GetTaskByID(c *gin.Context) {
	// We have a special task ID, like a name tag, and we want to find that task.
	id := c.Param("id")
	var task models.Task
	// We search in our database for the task with that special ID.
	if err := database.DB.Preload("Tags").First(&task, id).Error; err != nil {
		// If we can't find it, we say "Task not found."
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	// If we find it, we show the task to whoever asked for it.
	c.JSON(http.StatusOK, task)
}

// Create a Task
func CreateTask(c *gin.Context) {
	// We want to make a new task, like creating a new sticker.
	var task models.Task
	// We try to understand what the new task should look like from the request.
	if err := c.ShouldBindJSON(&task); err != nil {
		// If we can't understand, we say there's a mistake.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// We add the new task to our database, like putting the sticker in our collection.
	database.DB.Create(&task)
	// We show the new task to whoever asked for it.
	c.JSON(http.StatusCreated, task)
}

// Update a Task
func UpdateTask(c *gin.Context) {
	// We have a special task ID, and we want to change that task.
	id := c.Param("id")
	var task models.Task
	// We search in our database for the task with that special ID.
	if err := database.DB.First(&task, id).Error; err != nil {
		// If we can't find it, we say "Task not found."
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	// We try to understand what the updated task should look like from the request.
	if err := c.ShouldBindJSON(&task); err != nil {
		// If we can't understand, we say there's a mistake.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// We save the updated task in our database.
	database.DB.Save(&task)
	// We show the updated task to whoever asked for it.
	c.JSON(http.StatusOK, task)
}

// Delete a Task
func DeleteTask(c *gin.Context) {
	// We have a special task ID, and we want to remove that task.
	id := c.Param("id")
	// We try to delete the task with that special ID from our database.
	if err := database.DB.Delete(&models.Task{}, id).Error; err != nil {
		// If we can't find it, we say "Task not found."
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	// We say "Task deleted" to whoever asked for it.
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
