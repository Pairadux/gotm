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

	"github.com/charmbracelet/bubbles/list"
)

// WorkspaceItem is a wrapper for workspace data
type WorkspaceItem struct {
	workspace *models.Workspace
	name      string
}

// Implement the list.Item interface for Workspace
func (w WorkspaceItem) Title() string       { return w.name }
func (w WorkspaceItem) Description() string { return fmt.Sprintf("%d tasks", len(w.workspace.Tasks)) }
func (w WorkspaceItem) FilterValue() string { return w.name }

// createWorkspaceItems converts workspaces to list items
func createWorkspaceItems(taskState models.TaskState) []list.Item {
	workspaceItems := []list.Item{}
	for name, workspace := range taskState.Workspaces {
		workspaceItems = append(workspaceItems, WorkspaceItem{
			workspace: workspace,
			name:      name,
		})
	}
	return workspaceItems
}

// initWorkspaceList initializes the workspace list component
func initWorkspaceList(workspaceItems []list.Item) list.Model {
	workspaceList := list.New(workspaceItems, list.NewDefaultDelegate(), 0, 0)
	workspaceList.Title = "Workspaces"
	workspaceList.SetShowHelp(false)
	workspaceList.SetShowStatusBar(false)
	workspaceList.SetFilteringEnabled(false)
	workspaceList.Styles.Title = titleStyle

	return workspaceList
}

// handleWorkspaceSelection updates the task table when a workspace is selected
func handleWorkspaceSelection(m Model) Model {
	i := m.workspaces.Index()
	if i != -1 {
		workspaceItem := m.workspaces.Items()[i].(WorkspaceItem)
		workspace := workspaceItem.workspace

		rows := createTaskRows(workspace.Tasks)
		m.tasks.SetRows(rows)
	}
	return m
}
