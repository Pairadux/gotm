package tasks

import "fmt"
// IMPORTS {{{

) // }}}
}

}

func InitTasks() *Tasks {
	return &Tasks{tasks: []Task{}}
}

func (t *Tasks) AddTask(desc string) {
	t.tasks = append(t.tasks, Task{Description: desc})
	fmt.Println("Task added:", desc)
}
