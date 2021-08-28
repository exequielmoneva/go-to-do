package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@/database(tcp:localhost:3306)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nConnected to database")
	con = db
	return db
}
