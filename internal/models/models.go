package models

import "time"

// Workspace Struct
type Workspace struct {
	Tasks []Task `json:"tasks"`
}

// Task Struct
type Task struct {
	Index       int       `json:"index"`
	Created     time.Time `json:"created"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
}
