/* LICENSE {{{
Copyright © 2025 Austin Gause <a.gause@outlook.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/ // }}}

package taskops

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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
) // }}}

func InitActive() *models.TaskState {
	active, err := storage.Load(viper.GetString("active_path"))
	cobra.CheckErr(err)

	workspaces := active.Workspaces
	workspace := viper.GetString("default_workspace")

	if _, exists := workspaces[workspace]; !exists {
		workspaces[workspace] = &models.Workspace{
			ID:           "inbox",
			Name:         "Inbox",
			LastModified: time.Now(),
			Tasks:        []models.Task{},
		}
	}
	return active
}

func InitCompleted() *models.TaskState {
	completed, err := storage.Load(viper.GetString("completed_path"))
	cobra.CheckErr(err)

	workspaces := completed.Workspaces
	workspace := viper.GetString("default_workspace")

	if _, exists := workspaces[workspace]; !exists {
		workspaces[workspace] = &models.Workspace{
			ID:           "inbox",
			Name:         "Inbox",
			LastModified: time.Now(),
			Tasks:        []models.Task{},
		}
	}
	return completed
}

func InitAll() models.AppState {
	all := models.AppState{
		Active:    InitActive(),
		Completed: InitCompleted(),
	}

	return all
}

func Add(tasks *[]models.Task, desc string) {
	*tasks = append(*tasks, models.Task{
		Index:       0,
		Created:     time.Now(),
		Description: strings.TrimSpace(desc),
		Completed:   false,
	})
}

func Complete(appState models.AppState, workspace string, index int) (models.AppState, bool) {
	activeWorkspace := appState.Active.Workspaces[workspace]
	completedWorkspace := appState.Completed.Workspaces[workspace]

	for i := range activeWorkspace.Tasks {
		if activeWorkspace.Tasks[i].Index == index {
			activeWorkspace.Tasks[i].Completed = true
			removedTask, _ := Remove(&activeWorkspace.Tasks, i)
			completedWorkspace.Tasks = append(completedWorkspace.Tasks, removedTask)

			now := time.Now()
			completedWorkspace.LastModified = now
			activeWorkspace.LastModified = now

			return appState, true
		}
	}

	return appState, false
}

func Remove(tasks *[]models.Task, index int) (models.Task, bool) {
	for i, t := range *tasks {
		if t.Index == index {
			removedTask := t
			copy((*tasks)[i:], (*tasks)[i+1:])
			*tasks = (*tasks)[:len(*tasks)-1]
			return removedTask, true
		}
	}
	return models.Task{}, false
}

func Sort(sortType string, tasks []models.Task) {
	switch sortType {
	case "natural", "nat":
		// fmt.Printf("Sorting with: natural sort\n\n")
		utility.NaturalSort(tasks)
	default:
		fmt.Printf("Sort method: %s does not exist. Sorting with: natural sort\n\n", sortType)
		utility.NaturalSort(tasks)
	}

	for i := range tasks {
		tasks[i].Index = i + 1
	}
}

func PrintActive(tasks []models.Task) {
	r := strings.Repeat
	fmt.Printf("\n%-5s │ %-40s │ %s\n", "Index", "Description", "Created")
	fmt.Printf("%s┼%s┼%s\n", r("─", 6), r("─", 42), r("─", 20))
	for _, e := range tasks {
		fmt.Printf("%-5d │ %-40.40s │ %s\n", e.Index, e.Description, e.Created.Format(time.DateTime))
	}
}
