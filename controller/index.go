package controller

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/views"
	"net/http"
)

var todos []views.Todo // es una slice de Todo
func GetTodos(c *gin.Context) {
	if todos == nil {
		c.JSON(http.StatusOK, gin.H{})
	} else {
		c.JSON(http.StatusOK, todos)
	}
}

func CreateTodo(c *gin.Context) {
	var reqBody views.Todo
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "invalid request body",
		})
		return
	}
	todos = append(todos, reqBody)
	c.JSON(http.StatusOK, views.Todo{
		Todo: reqBody.Todo,
		Name: reqBody.Name,
	})
}
