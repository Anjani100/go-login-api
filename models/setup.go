package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"  // Registration of drivers to connect our server to the database
	// '_' is used to inform Go that we still want this included even though
	// we will never directly reference the package in our code.
)

var DB *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "GoUser"
)							// parameters for connecting to local database


func ConnectDatabase() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)	// Connects the database
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()	// sql.Open only validates the arguments provided
	// Ping forces the code to actually open up a connection to the database.	
	if err != nil {
		panic(err)
	}

	DB = db

	fmt.Println("Successfully connected!")
}