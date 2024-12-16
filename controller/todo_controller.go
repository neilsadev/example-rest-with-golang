package controller

import (
	"neilsadev/todo-api/database" // This helps us talk to our database where we keep our to-do lists.
	"neilsadev/todo-api/models"   // This helps us understand what a to-do list looks like.
	"net/http"                    // This helps us send messages over the internet.

	"github.com/gin-gonic/gin"    // This is a tool that helps us make our web server.
)

// GetTodoLists shows all the to-do lists we have.
func GetTodoLists(c *gin.Context) {
	// We have a big box called todoLists to keep all our to-do lists.
	var todoLists []models.TodoList
	// We look in our database and find all the to-do lists, including their tasks and tags.
	database.DB.Preload("Tasks.Tags").Find(&todoLists)
	// We show all the to-do lists we found to whoever asked for them.
	c.JSON(http.StatusOK, todoLists)
}

// GetTodoListByID shows a specific to-do list by its ID.
func GetTodoListByID(c *gin.Context) {
	// We have a special ID, like a name tag, to find a specific to-do list.
	id := c.Param("id")
	var todoList models.TodoList
	// We search in our database for the to-do list with that special ID.
	if err := database.DB.Preload("Tasks.Tags").First(&todoList, id).Error; err != nil {
		// If we can't find it, we say "TodoList not found."
		c.JSON(http.StatusNotFound, gin.H{"error": "TodoList not found"})
		return
	}
	// If we find it, we show the to-do list to whoever asked for it.
	c.JSON(http.StatusOK, todoList)
}

// CreateTodoList makes a new to-do list.
func CreateTodoList(c *gin.Context) {
	// We want to make a new to-do list, like creating a new page in our notebook.
	var todoList models.TodoList
	// We try to understand what the new to-do list should look like from the request.
	if err := c.ShouldBindJSON(&todoList); err != nil {
		// If we can't understand, we say there's a mistake.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// We add the new to-do list to our database, like putting the new page in our notebook.
	database.DB.Create(&todoList)
	// We show the new to-do list to whoever asked for it.
	c.JSON(http.StatusCreated, todoList)
}

// UpdateTodoList changes an existing to-do list.
func UpdateTodoList(c *gin.Context) {
	// We have a special ID to find the to-do list we want to change.
	id := c.Param("id")
	var todoList models.TodoList
	// We look for the to-do list in our database.
	if err := database.DB.First(&todoList, id).Error; err != nil {
		// If we can't find it, we say "TodoList not found."
		c.JSON(http.StatusNotFound, gin.H{"error": "TodoList not found"})
		return
	}
	// We try to understand the changes we need to make from the request.
	if err := c.ShouldBindJSON(&todoList); err != nil {
		// If we can't understand, we say there's a mistake.
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// We save the changes to our database.
	database.DB.Save(&todoList)
	// We show the updated to-do list to whoever asked for it.
	c.JSON(http.StatusOK, todoList)
}

// DeleteTodoList removes a to-do list.
func DeleteTodoList(c *gin.Context) {
	// We have a special ID to find the to-do list we want to remove.
	id := c.Param("id")
	// We try to remove the to-do list from our database.
	if err := database.DB.Delete(&models.TodoList{}, id).Error; err != nil {
		// If we can't find it, we say "TodoList not found."
		c.JSON(http.StatusNotFound, gin.H{"error": "TodoList not found"})
		return
	}
	// We say that the to-do list was successfully deleted.
	c.JSON(http.StatusOK, gin.H{"message": "TodoList deleted"})
}
