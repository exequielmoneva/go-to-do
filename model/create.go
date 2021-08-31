package model

import (
	"github.com/gin-gonic/gin"
	"go-boilerplate/views"
	"gorm.io/gorm"
	"net/http"
)

func CreateTodo(name string, todo string, id string) *gorm.DB {
	if err := con.Create(&views.Todo{
		Todo: todo,
		Name: name,
		Id:   id,
	}); err != nil {
		return err
	}
	return nil
}

func EditTodo(id string, reqBody views.TodoResponse, c *gin.Context) {
	/*if err := con.Where("id = ?", id).Find(&todo_); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "invalid id",
		})
		return
	}

	switch{
	case reqBody.Name != "" && reqBody.Todo != "":
		con.Model(&todo_).Select("name, todo").Updates(views.Todo{Todo: reqBody.Todo, Name: reqBody.Name})
	case reqBody.Name != "":
		con.Model(&todo_).Where("id = ?", id).Update("name", reqBody.Name)
	case reqBody.Todo != "":
		con.Model(&todo_).Where("id = ?", id).Update("todo", reqBody.Todo)
	default:
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "invalid request body",
		})
	}*/
	if err := con.Where("id = ?", id).Find(&todo_); err.RowsAffected <= 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "invalid id",
		})
		return
	}
	if reqBody.Name != "" {
		con.Model(&todo_).Where("id = ?", id).Update("name", reqBody.Name)

	}
	if reqBody.Todo != "" {
		con.Model(&todo_).Where("id = ?", id).Update("todo", reqBody.Todo)
	}
	con.Save(&todo_)

}

func DeleteTodo(id string, c *gin.Context) {
	if err := con.Where("id = ?", id).Delete(&todo_); err.RowsAffected <= 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "invalid id",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
