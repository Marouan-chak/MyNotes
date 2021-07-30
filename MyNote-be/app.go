package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "log"
    "encoding/json"
    "github.com/gorilla/mux"
)

type Note struct {
    Id string `json:"Id"`
    Title string `json:"Title"`
    Date string `json:"date"`
    Text string `json:"text"`
}
func store(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // return the string response containing the request body
    reqBody, _ := ioutil.ReadAll(r.Body)
    var note Note
    json.Unmarshal(reqBody, &note)
    Notes = append(Notes, note)
    json.NewEncoder(w).Encode(note)
}
func retrieve(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: retrieve")
    json.NewEncoder(w).Encode(Notes)
}
func retrieveSingleNote(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["id"]

    for _, note := range Notes {
        if note.Id == key {
            json.NewEncoder(w).Encode(note)
        }
    }
}

var Notes []Note
func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    // creates a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
    // replace http.HandleFunc with myRouter.HandleFunc
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/store", store).Methods("POST")
    myRouter.HandleFunc("/retrieve",retrieve)
    myRouter.HandleFunc("/retrieve/{id}", retrieveSingleNote)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
    fmt.Println("Rest API v2.0 - Mux Routers")
    Notes = []Note{
        Note{Id: "1",Title: "Hello", Date: "Note Dateription", Text: "Note Text"},
        Note{Id: "2",Title: "Hello 2", Date: "Note Dateription", Text: "Note Text"},
    }
    handleRequests()
    handleRequests()
}
