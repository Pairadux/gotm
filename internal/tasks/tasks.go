package tasks

// IMPORTS {{{
import (
	"fmt"
	"slices"
	"strings"
	"time"

	// temp
	// "math/rand/v2"

	"github.com/Pairadux/gotm/internal/models"
	"github.com/Pairadux/gotm/internal/storage"
	"github.com/facette/natsort"
	"github.com/spf13/viper"
) // }}}

func InitWorkspaces() map[string]*models.Workspace {
	workspaces, err := storage.LoadWorkspaces()
	if err != nil {
		panic(err)
	}

	if workspaces == nil {
		workspaces = make(map[string]*models.Workspace)
	}

	workspace := viper.GetString("default_workspace")

	if _, exists := workspaces[workspace]; !exists {
		workspaces[workspace] = &models.Workspace{
			Tasks: []models.Task{},
		}
	}
	return workspaces
}

func AddTask(workspaces map[string]*models.Workspace, workspace string, desc string) {
	workspaces[workspace].Tasks = append(workspaces[workspace].Tasks, models.Task{
		Index:       0,
		// Id:          AssignId(t),
		Created:     time.Now().Format(time.DateTime),
		Description: strings.TrimSpace(desc),
		Completed:   false,
	})
}

// TODO: Implement this method
func RemoveTask( /*t *models.TaskList, id int*/ ) {
	fmt.Println("Task Removed")
}

func SortTasks(sortType string, workspaces map[string]*models.Workspace, workspace string) {
	switch sortType {
	case "natural", "nat":
		fmt.Printf("Sorting with: natural sort\n\n")
		slices.SortFunc(workspaces[workspace].Tasks, func(a, b models.Task) int {
			if natsort.Compare(a.Description, b.Description) {
				return -1
			}
			if natsort.Compare(b.Description, a.Description) {
				return 1
			}
			return 0
		})

	// case "id-asc":
	// 	fmt.Printf("Sorting with: id-asc sort\n\n")
	// 	slices.SortFunc(t.Tasks, func(a, b models.Task) int {
	// 		return a.Id - b.Id
	// 	})

	default:
		fmt.Printf("Sort method: %s does not exist. Sorting with: natural sort\n\n", sortType)
		slices.SortFunc(workspaces[workspace].Tasks, func(a, b models.Task) int {
			if natsort.Compare(a.Description, b.Description) {
				return -1
			}
			if natsort.Compare(b.Description, a.Description) {
				return 1
			}
			return 0
		})
	}

	for i := range workspaces[workspace].Tasks {
		workspaces[workspace].Tasks[i].Index = i + 1
	}
}

// func AssignId(t *models.TaskList) int {
// 	// n := 1
// 	// SortTasks("id-asc", t)
// 	// for i, e := range t.Tasks {
// 	// 	if e.Id != n {
// 	// 		fmt.Println(e, i)
// 	// 	}
// 	// 	fmt.Println(i, e)
// 	// }
// 	return rand.IntN(10000)
// }
