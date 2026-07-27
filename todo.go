package main

import (
	"net/http"
	"strconv"

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
	id := c.Param("id")

	// id sent was not a number
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	for _, item := range todos {
		if item.ID == idInt {
			c.JSON(http.StatusOK, item)
			return
		}
	}

	// id was a number but not an id of an existing todo
	c.JSON(http.StatusNotFound, gin.H{})
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
