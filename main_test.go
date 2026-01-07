package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	initDB("test.db")
	code := m.Run()
	os.Remove("test.db")
	os.Exit(code)
}

func TestTasksHandlers(t *testing.T) {
	clearTasks()

	// Test Create
	taskJSON := `{"title":"Task 1","description":"Description 1","completed":false}`
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer([]byte(taskJSON)))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(tasksHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var createdTask Task
	json.Unmarshal(rr.Body.Bytes(), &createdTask)
	if createdTask.ID == 0 {
		t.Errorf("expected created task to have an ID")
	}

	// Test Get All
	req, _ = http.NewRequest("GET", "/tasks", nil)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var tasks []Task
	json.Unmarshal(rr.Body.Bytes(), &tasks)
	if len(tasks) != 1 {
		t.Errorf("expected 1 task, got %v", len(tasks))
	}

	// Test Get One
	req, _ = http.NewRequest("GET", "/tasks/1", nil)
	rr = httptest.NewRecorder()
	taskHandler := http.HandlerFunc(taskHandler)
	taskHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var fetchedTask Task
	json.Unmarshal(rr.Body.Bytes(), &fetchedTask)
	if !reflect.DeepEqual(createdTask, fetchedTask) {
		t.Errorf("handler returned unexpected body: got %v want %v", fetchedTask, createdTask)
	}

	// Test Update
	updateJSON := `{"title":"Updated Task 1","description":"Updated Description 1","completed":true}`
	req, _ = http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer([]byte(updateJSON)))
	rr = httptest.NewRecorder()
	taskHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}

	// Verify Update
	req, _ = http.NewRequest("GET", "/tasks/1", nil)
	rr = httptest.NewRecorder()
	taskHandler.ServeHTTP(rr, req)
	var updatedTask Task
	json.Unmarshal(rr.Body.Bytes(), &updatedTask)
	if updatedTask.Title != "Updated Task 1" {
		t.Errorf("expected title to be 'Updated Task 1', got '%v'", updatedTask.Title)
	}

	// Test Delete
	req, _ = http.NewRequest("DELETE", "/tasks/1", nil)
	rr = httptest.NewRecorder()
	taskHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}

	// Verify Delete
	req, _ = http.NewRequest("GET", "/tasks/1", nil)
	rr = httptest.NewRecorder()
	taskHandler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}
