package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"ToDoListAPI/database"
	"ToDoListAPI/env"
	"ToDoListAPI/models"

	"github.com/gorilla/mux"
)

// Function to check if given api key is correct
func validKey(r *http.Request) bool {
	if r.Header.Get("key") == env.Get("API_KEY") {
		return true
	} else {
		return false
	}
}

// Task CRUD: Create
func CreateTask(res http.ResponseWriter, req *http.Request) {
	if !validKey(req) {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write([]byte("401 - Invalid key"))
		return
	}
	var newTask models.Task
	var jsonResponse interface{}

	err := json.NewDecoder(req.Body).Decode(&newTask)
	switch {
	case err == io.EOF:
		// empty body
		jsonResponse = struct {
			Message string
		}{
			"Body of request cannot be blank",
		}
	case err != nil:
		// other error
		jsonResponse = struct {
			Error error
		}{
			err,
		}
	default:
		fmt.Println("Creating New Task...")

		// Add to DB
		newID, err := database.InsertTask(database.DB, newTask)
		newTask.ID = newID
		jsonResponse = struct {
			Message string
			Task    models.Task
		}{
			"Task Created",
			newTask,
		}

		if err != nil {
			jsonResponse = struct {
				Message string
				Error   error
			}{
				"Error While Creating New Task",
				err,
			}
		}
	}
	json.NewEncoder(res).Encode(jsonResponse)
}

// Task CRUD: Update
func UpdateTask(res http.ResponseWriter, req *http.Request) {
	if !validKey(req) {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write([]byte("401 - Invalid key"))
		return
	}
	
	var myTask models.Task
	var jsonResponse interface{}

	vars := mux.Vars(req)
	taskID, err := strconv.Atoi(vars["taskID"])
	if err != nil {
		panic(err)
	}

	myTask.ID = taskID

	err = json.NewDecoder(req.Body).Decode(&myTask)
	switch {
	case err == io.EOF:
		// empty body
		jsonResponse = struct {
			Message string
		}{
			"Body of request cannot be blank",
		}
	case err != nil:
		// other error
		jsonResponse = struct {
			Error error
		}{
			err,
		}
	default:
		fmt.Println("Updating Task of ID:", myTask.ID)

		// Update DB entry
		rowsAffected, err := database.UpdateTask(database.DB, myTask)
		if rowsAffected > 0 {
			jsonResponse = struct {
				Message string
				Task    models.Task
				}{
					"Task Updated",
					myTask,
				}
		} else {
			jsonResponse = struct {
				Message string
			}{
				"No Task Found",
			}
		}

		if err != nil {
			jsonResponse = struct {
				Message string
				Error   error
			}{
				"Error While Updating Task",
				err,
			}
		}
	}
	json.NewEncoder(res).Encode(jsonResponse)
}
