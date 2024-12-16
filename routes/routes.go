package routes

import (
	"neilsadev/todo-api/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(routes *gin.Engine) {
	userRoutes := routes.Group("/users")
	{
		userRoutes.GET("", controller.GetUsers)
		userRoutes.GET("/:id", controller.GetUserByID)
		userRoutes.POST("", controller.CreateUser)
		userRoutes.PUT("/:id", controller.UpdateUser)
		userRoutes.DELETE("/:id", controller.DeleteUser)
	}

	todoRoutes := routes.Group("/todos")
	{
		todoRoutes.GET("", controller.GetTodoLists)
		todoRoutes.GET("/:id", controller.GetTodoListByID)
		todoRoutes.POST("", controller.CreateTodoList)
		todoRoutes.PUT("/:id", controller.UpdateTodoList)
		todoRoutes.DELETE("/:id", controller.DeleteTodoList)
	}

	taskRoutes := routes.Group("/tasks")
	{
		taskRoutes.GET("", controller.GetTasks)
		taskRoutes.GET("/:id", controller.GetTaskByID)
		taskRoutes.POST("", controller.CreateTask)
		taskRoutes.PUT("/:id", controller.UpdateTask)
		taskRoutes.DELETE("/:id", controller.DeleteTask)
	}

	tagRoutes := routes.Group("/tags")
	{
		tagRoutes.GET("", controller.GetTags)
		tagRoutes.GET("/:id", controller.GetTagByID)
		tagRoutes.POST("", controller.CreateTag)
		tagRoutes.PUT("/:id", controller.UpdateTag)
		tagRoutes.DELETE("/:id", controller.DeleteTag)
	}
}
