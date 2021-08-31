package model

import (
	"go-boilerplate/views"
)

var todo_ []views.Todo

func ReadByName(username string) []views.Todo {
	con.Where("name = ?", username).Find(&todo_)
	return todo_

}

func ReadById(id string) []views.Todo {
	con.Where("id = ?", id).Find(&todo_)
	return todo_

}

func ReadAll() []views.Todo {
	con.Find(&todo_)
	return todo_

}
