package main

import (
	"database/sql" // package to encode and decode the json into struct and vice versa
	"fmt"          // used to access the request and response object of the api
	"os"

	// used to read the environment variable
	// package used to covert string into int type

	// used to get the params from the route

	// package used to read the .env file
	_ "github.com/lib/pq" // postgres golang driver
)

// response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}
type Note struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var host = os.Getenv("DB_URL")
var port = os.Getenv("DB_PORT")
var user = os.Getenv("APP_DB_USERNAME")
var password = os.Getenv("APP_DB_PASSWORD")
var dbname = os.Getenv("DB_NAME")

// create connection with postgres db
func createConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}
