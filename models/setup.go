package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"  // Registration of drivers to connect our server to the database
	// '_' is used to inform Go that we still want this included even though we will never directly reference the package in our code.
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("postgres", "GoUser.db")

	if err != nil {
		panic("Database connection failed!")
	}

	database.AutoMigrate(&User{})

	DB = database
}