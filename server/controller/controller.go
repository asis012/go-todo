package controller

import (
	"fmt"
	"log"
	"net/http"
	"todo/action"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() {
	router := mux.NewRouter()

	router.HandleFunc("/api/task", action.GetAllTaskHandler).Methods(action.GET)
	router.HandleFunc("/api/task", action.CreateTaskHandler).Methods(action.POST)
	router.HandleFunc("/api/task/{id}", action.CompleteTask).Methods(action.PUT)
	router.HandleFunc("/api/undoTask/{id}", action.UndoTaskHandler).Methods(action.PUT)
	router.HandleFunc("/api/deleteTask/{id}", action.DeleteTaskHandler).Methods(action.DELETE)
	router.HandleFunc("/api/deleteAllTasks", action.DeleteAllTasksHandler).Methods(action.DELETE)
	router.HandleFunc("/api/task/{id}", action.CompleteTask).Methods(action.PUT, "OPTIONS")
	router.HandleFunc("/api/task/{id}", action.GetTaskHandler).Methods(action.GET, "OPTIONS")
	router.HandleFunc("/api/updateTask/{id}", action.UpdateTaskHandler).Methods(action.PUT)

	handler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"OPTIONS", action.DELETE, action.GET, action.POST, action.PUT}),
		handlers.AllowedHeaders([]string{"Content-Type", "application/x-www-form-urlencoded", "Access-Control-Allow-Headers", "Content-Type"}),
	)

	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler(router)))
}
