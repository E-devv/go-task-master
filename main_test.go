package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetTasksHandler(t *testing.T) {
	initializeTasks()

	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getTasksHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned wrong content type: got %v want %v",
			contentType, expectedContentType)
	}

	var actual []Task
	err = json.Unmarshal(rr.Body.Bytes(), &actual)
	if err != nil {
		t.Fatalf("could not unmarshal response body: %v", err)
	}

	if !reflect.DeepEqual(actual, tasks) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			actual, tasks)
	}
}
