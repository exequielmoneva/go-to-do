package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	reqBody.Id = uuid.New().String()
	todos = append(todos, reqBody)
	c.JSON(http.StatusCreated, views.Todo{
		Todo: reqBody.Todo,
		Name: reqBody.Name,
		Id:   reqBody.Id,
	})
}

func EditTodo(c *gin.Context) {
	id := c.Param("id")
	var reqBody views.Todo
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "invalid request body",
		})
		return
	}
	for idx, todo := range todos {
		if todo.Id == id {
			if reqBody.Name != "" {
				todos[idx].Name = reqBody.Name
				c.JSON(http.StatusOK, views.Todo{
					Todo: todo.Todo,
					Name: reqBody.Name,
					Id:   id,
				})
			}
			if reqBody.Todo != "" {
				todos[idx].Todo = reqBody.Todo
				c.JSON(http.StatusOK, views.Todo{
					Todo: reqBody.Todo,
					Name: todo.Name,
					Id:   id,
				})
			}
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "invalid to-do id",
	})
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	for idx, todo := range todos {
		if todo.Id == id {
			// the 3 points at the end extract all the values from the slice
			todos = append(todos[:idx], todos[idx+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "to-do successfully deleted",
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "invalid to-do id",
	})
}
