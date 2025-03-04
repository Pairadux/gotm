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

func Add(tasks *[]models.Task, desc string) {
	*tasks = append(*tasks, models.Task{
		Index:       0,
		Created:     time.Now(),
		Description: strings.TrimSpace(desc),
		Completed:   false,
	})
}

// TODO: complete function
func Complete(tasks *[]models.Task, index int) (models.Task, bool) {
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

func Print(tasks []models.Task) {
	r := strings.Repeat
	fmt.Printf("\n%-5s │ %-40s │ %s\n", "Index", "Description", "Created")
	fmt.Printf("%s┼%s┼%s\n", r("─", 6), r("─", 42), r("─", 20))
	for _, e := range tasks {
		fmt.Printf("%-5d │ %-40.40s │ %s\n", e.Index, e.Description, e.Created.Format(time.DateTime))
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
