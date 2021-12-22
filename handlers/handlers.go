package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"ToDoListAPI/env"
	"ToDoListAPI/models"
)

// Function to check if given api key is correct
func validKey(r *http.Request) bool {
	if r.Header.Get("key") == env.Get("API_KEY") {
		return true
	} else {
		return false
	}
}

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
		fmt.Println("Name: ", newTask.Name)
		fmt.Println("Done: ", newTask.Done)
		jsonResponse = struct {
			Message string
			Task    models.Task
			}{
				"Task Created",
				newTask,
			}
	}
	json.NewEncoder(res).Encode(jsonResponse)
}
