package controller

import (
	"github.com/gin-gonic/gin"
)

func Register() *gin.Engine {
	r := gin.Default()
	todoRoutes := r.Group("/")
	{
		todoRoutes.GET("", GetTodos)
		todoRoutes.GET("single/:id", GetTodoById)
		todoRoutes.GET("user/:username", GetTodoFromUser)
		todoRoutes.GET("health", Health)
		todoRoutes.POST("", CreateTodo)
		todoRoutes.PUT("edit/:id", EditTodo)
		todoRoutes.DELETE("delete/:id", DeleteTodo)
	}
	return r
}
