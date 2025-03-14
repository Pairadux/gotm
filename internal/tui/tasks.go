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

package tui

import (
	"fmt"

	"github.com/Pairadux/gotm/internal/models"

	"github.com/charmbracelet/bubbles/table"
)

// Task represents a single task
type Task struct {
	id          int
	title       string
	description string
	status      string
	priority    string
}

// initTaskTable initializes the task table component
func initTaskTable(taskState models.TaskState, activeWorkspace string) table.Model {
	columns := []table.Column{
		{Title: "ID", Width: 4},
		{Title: "Description", Width: 50},
		{Title: "Created", Width: 12},
		{Title: "Status", Width: 8},
	}

	rows := []table.Row{}
	if activeWorkspace != "" && taskState.Workspaces[activeWorkspace] != nil {
		rows = createTaskRows(taskState.Workspaces[activeWorkspace].Tasks)
	}

	taskTable := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
	)

	// Set table styles
	taskTable.SetStyles(GetTableStyles())

	return taskTable
}

// createTaskRows converts tasks to table rows
func createTaskRows(tasks []models.Task) []table.Row {
	rows := []table.Row{}
	for _, task := range tasks {
		status := "Todo"
		if task.Completed {
			status = "Done"
		}
		rows = append(rows, table.Row{
			fmt.Sprintf("%d", task.Index),
			task.Description,
			task.Created.Format("2006-01-02"),
			status,
		})
	}
	return rows
}
