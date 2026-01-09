package main

import "errors"

var (
	ErrTitleEmpty     = errors.New("title cannot be empty")
	ErrInvalidDueDate = errors.New("due date must be in YYYY-MM-DD format")
)
