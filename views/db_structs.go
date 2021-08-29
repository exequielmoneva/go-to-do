package views

import (
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Todo string
	Name string
	Id   string
}
