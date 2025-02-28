package models

// Workspace Struct
type Workspace struct {
	Tasks []Task `json:"tasks"`
}

// Task Struct
type Task struct {
	Index       int    `json:"index"`
	Created     string `json:"created"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
