package model

import (
	"fmt"
	"go-boilerplate/views"
	"gorm.io/gorm"
)

func CreateTodo(name string, todo string, id string) *gorm.DB {
	if err := con.Create(&views.Todo{
		Todo: todo,
		Name: name,
		Id:   id,
	}); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
