package main

import (
	"fmt"
	"log"

	"ToDoListAPI/handlers"
	"net/http"

	"github.com/gorilla/mux"

	"ToDoListAPI/database"
)

// Map handlers and start the http server
func StartHTTPServer() {
	// New Router Instance
	r := mux.NewRouter()

	// REST: Create
	r.HandleFunc("/api/v1/tasks", handlers.CreateTask).Methods("POST")

	// REST: ReadAll
	// r.HandleFunc("/api/v1/tasks", handlers.ReadAllTasks).Methods("GET")

	// REST: ReadOne
	// r.HandleFunc("/api/v1/tasks/{taskID}", handlers.ReadTask).Methods("GET")

	// REST: Update
	r.HandleFunc("/api/v1/tasks/{taskID}", handlers.UpdateTask).Methods("PATCH")

	// REST: Delete
	r.HandleFunc("/api/v1/tasks/{taskID}", handlers.DeleteTask).Methods("DELETE")

	// Serve the server at port 8080, log fatals
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	// Connect to DB
	database.Connect()

	// defer closing of DB, executed when main func ends
	defer database.DB.Close()

	// ping db (test connection)
	err := database.DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("PostgreSQL DB Successfully connected!")

	// Create "tasks" table in not exist
	database.CreateTaskTable(database.DB)

	// Start HTTP Server
	StartHTTPServer()
}
