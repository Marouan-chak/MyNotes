package main

import (
	"log"      // package to encode and decode the json into struct and vice versa
	"net/http" // used to access the request and response object of the api

	"github.com/gorilla/mux" // used to get the params from the route
	_ "github.com/lib/pq"    // postgres golang driver
)

func handleRequests() {
	// creates a new instance of a mux
	myRouter := mux.NewRouter().StrictSlash(true)
	// create different routes
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/api/store", store).Methods("POST")
	myRouter.HandleFunc("/api/retrieve", retrieve)
	myRouter.HandleFunc("/api/delete/{id}", deleteNote).Methods("DELETE")
	myRouter.HandleFunc("/api/update/{id}", updateNote).Methods("PUT")
	myRouter.HandleFunc("/api/health-check", HealthCheck).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
