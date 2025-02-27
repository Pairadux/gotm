package models

// Task Struct
type Task struct {
	Index       int    `json:"index"`
	Id          int    `json:"id"`
	Created     string `json:"created"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// TaskList Struct
type TaskList struct {
	Tasks []Task `json:"tasklist"`
}
