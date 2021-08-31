package main

import (
	_ "github.com/gin-gonic/gin"
	"go-to-do/controller"
	"go-to-do/model"
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
