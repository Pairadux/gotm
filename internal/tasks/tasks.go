package tasks

// IMPORTS {{{
import (
	"fmt"
	"slices"
	"time"

	// temp
	"math/rand/v2"

	"github.com/Pairadux/gotm/internal/models"
	"github.com/Pairadux/gotm/internal/storage"
	"github.com/facette/natsort"
	// "github.com/spf13/viper"
) // }}}

func InitTasks() *models.TaskList {
	t, err := storage.LoadTasks()
	if err != nil {
		panic(err)
	}
	return t
}

func AddTask(t *models.TaskList, desc string) {
	t.Tasks = append(t.Tasks, models.Task{
		Id:          FindFirstAvailableId(t),
		Created:     time.Now().Format(time.DateTime),
		Description: desc,
		Completed:   false,
	})
}

// TODO: Implement this method
func RemoveTask( /*t *models.TaskList, id int*/ ) {
	fmt.Println("Task Removed")
}

func GetTasks(t models.TaskList) []models.Task {
	return t.Tasks
}

func SortTasks(sortType string, t *models.TaskList) {
	switch sortType {
	case "natural":
		fmt.Printf("\nSorting with: natural sort\n\n")
		slices.SortFunc(t.Tasks, func(a, b models.Task) int {
			if natsort.Compare(a.Description, b.Description) {
				return -1
			}
			if natsort.Compare(b.Description, a.Description) {
				return 1
			}
			return 0
		})
	case "id-asc":
		fmt.Printf("\nSorting with: id-asc sort\n\n")
		slices.SortFunc(t.Tasks, func(a, b models.Task) int {
			return a.Id - b.Id
		})
	}
}

func FindFirstAvailableId(t *models.TaskList) int {
	// n := 1
	SortTasks("id-asc", t)
	// for i, e := range t.Tasks {
	// 	if e.Id != n {
	// 		fmt.Println(e, i)
	// 	}
	// 	fmt.Println(i, e)
	// }
	return rand.IntN(10000)
}
