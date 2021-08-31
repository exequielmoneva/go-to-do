package views

import (
	_ "gorm.io/driver/sqlite"
)

type Todo struct {
	Id   string `gorm:"primarykey"`
	Todo string
	Name string
}
