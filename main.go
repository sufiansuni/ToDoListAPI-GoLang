package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"ToDoListAPI/handlers"
)

// Map handlers and start the http server
func StartHTTPServer() {
	// New Router Instance
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/create_task", handlers.CreateTask).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	StartHTTPServer()
}
