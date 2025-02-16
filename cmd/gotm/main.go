package main

import (
	"fmt"
	"github.com/Pairadux/gotm/internal/tasks"
)

func main() {
	fmt.Println("")
	tm := tasks.InitTasks()
	tm.AddTask("Go is cool")
}
