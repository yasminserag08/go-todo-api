package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{
		ID:        1,
		Title:     "task1",
		Completed: true,
	},
}

var nextID int = 1

// handlers
func returnTasks(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func getTaskById(c *gin.Context) {
	// TODO
}

func addTask(c *gin.Context) {
	// TODO
}

func editTask(c *gin.Context) {
	// TODO
}

func deleteTask(c *gin.Context) {
	// TODO
}
