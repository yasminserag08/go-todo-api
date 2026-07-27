package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/todos", returnTasks)

	r.GET("/todos/:id", getTaskById)

	r.POST("/todos", addTask)

	r.PUT("/todos:id", editTask)

	r.DELETE("/todos:id", deleteTask)

	r.Run()
}
