package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"  // Registration of drivers to connect our server to the database
	// '_' is used to inform Go that we still want this included even though we will never directly reference the package in our code.
)
  
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "GoUser"
)

func ConnectDatabase() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}