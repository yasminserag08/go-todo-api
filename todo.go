package main

import (
	"net/http"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        int     `json:"id"`
	Title     *string `json:"title"`
	Completed *bool   `json:"completed"`
}

var todos = []todo{}

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, item := range todos {
		if item.ID == idInt {
			c.JSON(http.StatusOK, item)
			return
		}
	}

	// id was a number but not of an existing todo
	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func addTask(c *gin.Context) {
	var newTask todo
	err := c.ShouldBindJSON(&newTask)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(*(newTask.Title)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title cannot be empty"})
		return
	}

	newTask.ID = nextID
	nextID += 1
	todos = append(todos, newTask)
	c.JSON(http.StatusOK, newTask)
}

func editTask(c *gin.Context) {
	id := c.Param("id")
	// id sent was not a number
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedTodo todo

	for i, item := range todos {
		if item.ID == idInt {
			err := c.ShouldBindJSON(&updatedTodo)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				if updatedTodo.Title != nil {
					todos[i].Title = updatedTodo.Title
				}
				if updatedTodo.Completed != nil {
					todos[i].Completed = updatedTodo.Completed
				}
				c.JSON(http.StatusOK, todos[i])
			}
			return
		}
	}

	// id was a number but not of an existing todo
	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func deleteTask(c *gin.Context) {
	id := c.Param("id")

	// id sent was not a number
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, item := range todos {
		if item.ID == idInt {
			todos = slices.Delete(todos, i, i+1)
			c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
			return
		}
	}

	// id was a number but not of an existing todo
	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}
