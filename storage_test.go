package main

import (
	"testing"
)

func TestCreateTask(t *testing.T) {
	clearTasks()
	task := Task{Title: "Test Task", Description: "Test Description", Completed: false}
	createdTask, err := createTask(task)
	if err != nil {
		tFatalf(t, "createTask() error = %v", err)
	}
	if createdTask.ID == 0 {
		t.Errorf("Expected task ID to be set, got 0")
	}
}

func TestUpdateTask(t *testing.T) {
	clearTasks()
	task := Task{Title: "Test Task", Description: "Test Description", Completed: false}
	createdTask, _ := createTask(task)

	createdTask.Completed = true
	err := updateTask(createdTask)
	if err != nil {
		t.Fatalf("updateTask() error = %v", err)
	}

	updatedTask, err := getTaskByID(createdTask.ID)
	if err != nil {
		t.Fatalf("getTaskByID() error = %v", err)
	}
	if !updatedTask.Completed {
		t.Errorf("Expected task to be completed, but it was not")
	}
}

func TestDeleteTask(t *testing.T) {
	clearTasks()
	task := Task{Title: "Test Task", Description: "Test Description", Completed: false}
	createdTask, _ := createTask(task)

	err := deleteTask(createdTask.ID)
	if err != nil {
		t.Fatalf("deleteTask() error = %v", err)
	}

	_, err = getTaskByID(createdTask.ID)
	if err == nil {
		t.Errorf("Expected error when getting deleted task, but got nil")
	}
}

func tFatalf(t *testing.T, format string, args ...interface{}) {
	t.Helper()
	t.Fatalf(format, args...)
}
