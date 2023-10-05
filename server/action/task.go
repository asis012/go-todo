package action

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"todo/data"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

// GetAllTask get all the task route
func GetAllTaskHandler(w http.ResponseWriter, r *http.Request) {
	payload := data.GetAllTask()
	json.NewEncoder(w).Encode(payload)
}

// CreateTask create task route
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task data.ToDoList
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		//log
		return
	}
	id := data.InsertTask(task)
	task.ID = id
	json.NewEncoder(w).Encode(task)
}

// CompleteTask update task route
func CompleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	success := data.CompleteTask(params["id"])

	// Create a map with the success value
	response := map[string]bool{"success": success}

	// Encode the map as a JSON object and return it as the response
	json.NewEncoder(w).Encode(response)
}

// UndoTask undo the complete task route
func UndoTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	success := data.UndoTask(params["id"])
	// Create a map with the success value
	response := map[string]bool{"success": success}
	json.NewEncoder(w).Encode(response)
}

// DeleteTask delete one task route
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	data.DeleteTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

// DeleteAllTasks delete all tasks route
func DeleteAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	count := data.DeleteAllTasks()
	fmt.Println("a")
	fmt.Println(count)
	json.NewEncoder(w).Encode(count)

}

// GetTask get the task route
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	// Query the database for the task with the given ID
	var task data.ToDoList
	err = data.Collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&task)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	// Return the task as a JSON response
	json.NewEncoder(w).Encode(task)
}

// UpdateTask update the task route
func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var task data.ToDoList
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	success := data.UpdateTask(task, params["id"])
	// Create a map with the success value
	response := map[string]bool{"success": success}
	json.NewEncoder(w).Encode(response)
}
