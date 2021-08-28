package main

import (
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"go-boilerplate/controller"
	"go-boilerplate/model"
	"log"
)

func main() {
	r := controller.Register()
	db := model.Connect()
	defer db.Close()

	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())
	}
}
