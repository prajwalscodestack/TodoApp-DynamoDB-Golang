package main

import (
	"log"
	"todo-app-dynamodb/db"
	"todo-app-dynamodb/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize DynamoDB client
	db.InitDynamoDB()

	router := gin.Default()

	// Routes
	router.POST("/todos", handlers.CreateTodo)
	router.GET("/todos/:id", handlers.GetTodo)
	router.PUT("/todos/:id", handlers.UpdateTodo)
	router.DELETE("/todos/:id", handlers.DeleteTodo)

	log.Println("Server running at http://localhost:8080")
	router.Run(":8080")
}
