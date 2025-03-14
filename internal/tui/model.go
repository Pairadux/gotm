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
	"github.com/Pairadux/gotm/internal/models"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

// Model represents the application state
type Model struct {
	workspaces   list.Model
	tasks        table.Model
	activePane   int // 0 for workspaces, 1 for tasks
	windowWidth  int
	windowHeight int
}

// Init initializes the tea program
func (m Model) Init() tea.Cmd {
	return nil
}

// InitialModel creates and initializes the application model
func InitialModel(taskState models.TaskState, activeWorkspace string) Model {
	workspaceItems := createWorkspaceItems(taskState)

	// Initialize workspaces list
	workspaceList := initWorkspaceList(workspaceItems)

	// Initialize tasks table
	taskTable := initTaskTable(taskState, activeWorkspace)

	return Model{
		workspaces:   workspaceList,
		tasks:        taskTable,
		activePane:   0,
		windowWidth:  0,
		windowHeight: 0,
	}
}

// Update handles events and updates the model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, keys.Tab), key.Matches(msg, keys.Left), key.Matches(msg, keys.Right):
			// Toggle active pane
			if key.Matches(msg, keys.Tab) {
				m.activePane = (m.activePane + 1) % 2
			} else if key.Matches(msg, keys.Left) {
				m.activePane = 0
			} else if key.Matches(msg, keys.Right) {
				m.activePane = 1
			}

		case key.Matches(msg, keys.Enter):
			if m.activePane == 0 {
				// Handle workspace selection
				m = handleWorkspaceSelection(m)
			}
		}

	case tea.WindowSizeMsg:
		m = handleWindowResize(m, msg)
	}

	// Update active component
	if m.activePane == 0 {
		m.workspaces, cmd = m.workspaces.Update(msg)
		cmds = append(cmds, cmd)
	} else {
		m.tasks, cmd = m.tasks.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
