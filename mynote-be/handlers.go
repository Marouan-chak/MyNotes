package main

import (
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"log"
	"net/http" // used to access the request and response object of the api
	"strconv"  // package used to covert string into int type

	"github.com/gorilla/mux" // used to get the params from the route

	_ "github.com/lib/pq" // postgres golang driver
)

func store(w http.ResponseWriter, r *http.Request) {

	// set the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create an empty note of type Note
	var note Note

	// decode the json request to note
	err := json.NewDecoder(r.Body).Decode(&note)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call storeNote function and pass the note
	insertID := storeNote(note)

	// format a response object
	res := response{
		ID:      insertID,
		Message: "Note created successfully",
	}

	// send the response
	json.NewEncoder(w).Encode(res)

}

// insert one note in the DB
func storeNote(note Note) int64 {

	// create the postgres db connection
	db := createConnection()

	// create the insert sql query
	sqlStatement := `INSERT INTO notes (title, text) VALUES ($1, $2) RETURNING id`

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	err := db.QueryRow(sqlStatement, note.Title, note.Text).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}

func retrieve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get all the notes in the db
	notes, err := getAllNotes()

	if err != nil {
		log.Fatalf("Unable to get all note. %v", err)
	}

	// send all the  notes as response
	json.NewEncoder(w).Encode(notes)
}
func getAllNotes() ([]Note, error) {
	// create the postgres db connection
	db := createConnection()

	var notes []Note

	// create the select sql query
	sqlStatement := `SELECT * FROM notes`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var note Note

		err = rows.Scan(&note.Id, &note.Title, &note.Text)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the note in the notes slice
		notes = append(notes, note)

	}

	// return empty note on error
	return notes, err
}

func updateNote(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the noteid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// create an empty note of type models.Note
	var note Note

	// decode the json request to note
	err = json.NewDecoder(r.Body).Decode(&note)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call update note to update the note
	updatedRows := UpdateNote(int64(id), note)

	// format the message string
	msg := fmt.Sprintf("Note updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)

}

// update note in the DB
func UpdateNote(id int64, note Note) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	//defer db.Close()

	// create the update sql query
	sqlStatement := `UPDATE notes SET title=$2, text=$3 WHERE id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id, note.Title, note.Text)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

func deleteNote(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the noteid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id in string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// call the deleteNote, convert the int to int64
	deletedRows := DeleteNote(int64(id))

	// format the message string
	msg := fmt.Sprintf("Note updated successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	//specify status code
	w.WriteHeader(http.StatusOK)

	//update response writer
	fmt.Fprintf(w, "API is up and running")
}

func DeleteNote(id int64) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	//defer db.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM notes WHERE id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

var Notes []Note

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}
