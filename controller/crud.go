package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-to-do/model"
	"go-to-do/views"
	"net/http"
)

func CreateTodo(c *gin.Context) {
	// Create a New To-do
	var reqBody views.Todo
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "invalid request body",
		})
		return
	}
	reqBody.Id = uuid.New().String()
	model.CreateTodo(reqBody.Name, reqBody.Todo, reqBody.Id)
	c.JSON(http.StatusCreated, views.TodoResponse{
		Todo: reqBody.Todo,
		Name: reqBody.Name,
		Id:   reqBody.Id,
	})
}

func GetTodos(c *gin.Context) {
	var result = model.ReadAll()
	if result == nil {
		c.JSON(http.StatusOK, gin.H{})
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func GetTodoById(c *gin.Context) {
	// Get an existing to-do based on the given id as queryparam
	id := c.Param("id")
	var result = model.ReadById(id)
	if result == nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "invalid id",
		})
		return
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func GetTodoFromUser(c *gin.Context) {
	// Get an existing to-do based on the given username as queryparam
	username := c.Param("username")
	var result = model.ReadByName(username)
	c.JSON(http.StatusOK, result)
}

func EditTodo(c *gin.Context) {
	// Edit an existing to do matching by its id
	id := c.Param("id")
	var reqBody views.TodoResponse
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "invalid request body",
		})
		return
	}
	model.EditTodo(id, reqBody, c)
}

func DeleteTodo(c *gin.Context) {
	// Delete an existing to do matching by its id
	id := c.Param("id")

	model.DeleteTodo(id, c)
}

func Health(c *gin.Context) {
	// Check whether the API is working or not
	c.JSON(http.StatusOK, gin.H{
		"message": "API is up",
	})
}
