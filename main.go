package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Task represents a single task
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var tasks []Task

func initializeTasks() {
	tasks = []Task{
		{ID: 1, Title: "Task 1", Description: "Description 1", Completed: false},
		{ID: 2, Title: "Task 2", Description: "Description 2", Completed: true},
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func main() {
	initializeTasks()
	http.HandleFunc("/health", healthCheckHandler)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
