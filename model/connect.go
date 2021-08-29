package model

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var con *gorm.DB

func Connect() *sql.DB {
	/*db, err := sql.Open("mysql", "root:1234@/database(tcp:localhost:3306)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nConnected to database")
	con = db
	return db*/
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	// db.AutoMigrate(&views.Todo{})

	fmt.Println("\nConnected to database")
	con = db
	sqlDB, err := db.DB()
	return sqlDB
}
