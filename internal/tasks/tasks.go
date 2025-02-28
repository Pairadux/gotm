package tasks

// IMPORTS {{{
import (
	"fmt"
	"strings"
	"time"

	// temp
	// "math/rand/v2"

	"github.com/Pairadux/gotm/internal/models"
	"github.com/Pairadux/gotm/internal/storage"
	"github.com/Pairadux/gotm/internal/utility"

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

func Add(workspaces map[string]*models.Workspace, workspace string, desc string) {
	workspaces[workspace].Tasks = append(workspaces[workspace].Tasks, models.Task{
		Index:       0,
		// Id:          AssignId(t),
		Created:     time.Now(),
		Description: strings.TrimSpace(desc),
		Completed:   false,
	})
}

// TODO: Implement this method
func Remove( /*t *models.TaskList, id int*/ ) {
	fmt.Println("Task Removed")
}

func Sort(sortType string, ts []models.Task) {
	switch sortType {
	case "natural", "nat":
		fmt.Printf("Sorting with: natural sort\n\n")
		utility.NaturalSort(ts)
	default:
		fmt.Printf("Sort method: %s does not exist. Sorting with: natural sort\n\n", sortType)
		utility.NaturalSort(ts)
	}

	for i := range ts {
		ts[i].Index = i + 1
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
