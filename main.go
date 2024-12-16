package main

import (
	"neilsadev/todo-api/database"
	"neilsadev/todo-api/routes"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDatabase()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
