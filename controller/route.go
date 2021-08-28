package controller

import (
	"github.com/gin-gonic/gin"
)

func Register() *gin.Engine {
	r := gin.Default()
	todoRoutes := r.Group("/")
	{
		todoRoutes.GET("", GetTodos)
		todoRoutes.POST("", CreateTodo)
		todoRoutes.PUT(":id", EditTodo)
		todoRoutes.DELETE(":id", DeleteTodo)
	}
	return r
}