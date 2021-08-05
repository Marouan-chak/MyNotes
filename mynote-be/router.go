package main

import (

	// package to encode and decode the json into struct and vice versa
	"log"
	"net/http" // used to access the request and response object of the api

	// used to read the environment variable
	// package used to covert string into int type

	"github.com/gorilla/mux" // used to get the params from the route

	// package used to read the .env file
	_ "github.com/lib/pq" // postgres golang driver
)

func handleRequests() {
	// creates a new instance of a mux ro:w
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/api/store", store).Methods("POST")
	myRouter.HandleFunc("/api/retrieve", retrieve)
	myRouter.HandleFunc("/api/retrieve/{id}", retrieveSingleNote)
	myRouter.HandleFunc("/api/delete/{id}", deleteNote).Methods("DELETE")
	myRouter.HandleFunc("/api/update/{id}", updateNote).Methods("PUT")
	myRouter.HandleFunc("/api/health-check", HealthCheck).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
