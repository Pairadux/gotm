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
	"slices"
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

	return active
}

func InitCompleted() *models.TaskState {
	completed, err := storage.Load(viper.GetString("completed_path"))
	cobra.CheckErr(err)

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

	for i, task := range activeWorkspace.Tasks {
		if task.Index == index {
			completedTask := task
			completedTask.Completed = true

			activeWorkspace.Tasks = slices.Delete(activeWorkspace.Tasks, i, i+1)

			completedWorkspace.Tasks = append(completedWorkspace.Tasks, completedTask)

			now := time.Now()

			activeWorkspace.LastModified = now
			completedWorkspace.LastModified = now

			return appState, true
		}
	}

	return appState, false
}

func GetTasks(taskState models.TaskState, workspace string) []models.Task {
	return taskState.Workspaces[workspace].Tasks
}

func Remove(tasks *[]models.Task, index int) (models.Task, bool) {
	for i, t := range *tasks {
		if t.Index == index {
			removedTask := t
			*tasks = slices.Delete(*tasks, i, i+1)
			return removedTask, true
		}
	}
	return models.Task{}, false
}

func Sort(tasks []models.Task, sortType string) {
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

// CreateWorkspace takes a workspace and a taskState and creates the workspace and adds it to the taskState.
func CreateWorkspace(taskState models.TaskState, workspace string) {
	taskState.Workspaces[workspace] = &models.Workspace{
		Name:         workspace,
		LastModified: time.Now(),
		Tasks:        []models.Task{},
	}
}

// DeleteWorkspace takes a workspace and a taskState and deletes the workspace from the taskState.
func DeleteWorkspace(taskState models.TaskState, workspace string) {
	delete(taskState.Workspaces, workspace)
}

