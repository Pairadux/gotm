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

// IMPORTS {{{
import (
	"fmt"

	// "slices"
	// "strconv"

	// "github.com/Pairadux/gotm/internal/models"
	// "github.com/Pairadux/gotm/internal/taskops"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
) // }}}

// Model represents the application state
type Model struct {
	workspaces   list.Model
	tasks        table.Model
	activePane   int // 0 for workspaces, 1 for tasks
	windowWidth  int
	windowHeight int
}

// Workspace is a group of tasks
type Workspace struct {
	title string
	id    int
	tasks []Task
}

// Task represents a single task
type Task struct {
	id          int
	title       string
	description string
	status      string
	priority    string
}

// Implement the list.Item interface for Workspace
func (w Workspace) Title() string       { return w.title }
func (w Workspace) Description() string { return fmt.Sprintf("%d tasks", len(w.tasks)) }
func (w Workspace) FilterValue() string { return w.title }

// Define keyboard mappings
type keyMap struct {
	Up    key.Binding
	Down  key.Binding
	Left  key.Binding
	Right key.Binding
	Tab   key.Binding
	Enter key.Binding
	Quit  key.Binding
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "focus workspaces"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "focus tasks"),
	),
	Tab: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "switch pane"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q/ctrl+c", "quit"),
	),
}

// Style definitions
var (
	baseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("240")).
			Padding(0, 0, 0, 0)

	workspaceStyle = baseStyle.
			BorderForeground(lipgloss.Color("36"))

	taskStyle = baseStyle.
			BorderForeground(lipgloss.Color("63"))

	focusedWorkspaceStyle = workspaceStyle.
				BorderForeground(lipgloss.Color("86"))

	focusedTaskStyle = taskStyle.
				BorderForeground(lipgloss.Color("87"))

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("86")).
			Bold(true).
			MarginLeft(1)

	statusBar = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			MarginTop(1)

	appStyle = lipgloss.NewStyle().
			Padding(1, 2)
)

// Initialize the model
	// Sample data
	workspaces := []list.Item{
		Workspace{title: "Personal", id: 1, tasks: []Task{}},
		Workspace{title: "Work", id: 2, tasks: []Task{}},
		Workspace{title: "Side Projects", id: 3, tasks: []Task{}},
		Workspace{title: "Learning", id: 4, tasks: []Task{}},
		Workspace{title: "Health", id: 5, tasks: []Task{}},
	}

	// Initialize workspaces list
	workspaceList := list.New(workspaces, list.NewDefaultDelegate(), 0, 0)
	workspaceList.Title = "Workspaces"
	workspaceList.SetShowHelp(false)
	workspaceList.SetShowStatusBar(false)
	workspaceList.SetFilteringEnabled(false)
	workspaceList.Styles.Title = titleStyle

	// Initialize tasks table
	columns := []table.Column{
		{Title: "ID", Width: 4},
		{Title: "Title", Width: 30},
		{Title: "Status", Width: 10},
		{Title: "Priority", Width: 8},
	}

	rows := []table.Row{
		{"1", "Complete the TUI application", "Todo", "High"},
		{"2", "Add persistence layer", "Todo", "Medium"},
		{"3", "Implement task filtering", "Todo", "Low"},
	}

	taskTable := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
	)

	// Set table styles
	taskTable.SetStyles(table.Styles{
		Header: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("87")).
			Padding(0, 1).
			Align(lipgloss.Center),
		Selected: lipgloss.NewStyle().
			Foreground(lipgloss.Color("229")).
			Background(lipgloss.Color("63")).
			Bold(true),
		Cell: lipgloss.NewStyle().
			Padding(0, 1).
			Align(lipgloss.Left),
	})

	return Model{
		workspaces:   workspaceList,
		tasks:        taskTable,
		activePane:   0,
		windowWidth:  0,
		windowHeight: 0,
	}
}

// Init initializes the tea program
func (m Model) Init() tea.Cmd {
	return nil
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
				i := m.workspaces.Index()
				if i != -1 {
					workspace := m.workspaces.Items()[i].(Workspace)
					// Here we would update tasks based on the selected workspace
					// For now, we'll just update the title
					m.tasks.SetRows([]table.Row{
						{"1", fmt.Sprintf("Task 1 in %s", workspace.title), "Todo", "High"},
						{"2", fmt.Sprintf("Task 2 in %s", workspace.title), "In Progress", "Medium"},
						{"3", fmt.Sprintf("Task 3 in %s", workspace.title), "Done", "Low"},
					})
				}
			}
		}

	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height

		// Account for padding and borders
		contentWidth := msg.Width - 4   // -4 for horizontal padding/margins
		contentHeight := msg.Height - 4 // -4 for vertical padding/margins

		// Calculate widths
		workspacesWidth := contentWidth / 4
		tasksWidth := contentWidth - workspacesWidth - 4 // account for separator and borders

		// Update list dimensions
		m.workspaces.SetHeight(contentHeight - 2)
		m.workspaces.SetWidth(workspacesWidth)

		// Update table height and width
		m.tasks.SetHeight(contentHeight - 4)

		// Update column widths
		columns := m.tasks.Columns()
		columns[1].Width = tasksWidth - columns[0].Width - columns[2].Width - columns[3].Width - 10
		m.tasks.SetColumns(columns)
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

// View renders the UI
func (m Model) View() string {
	if m.windowWidth == 0 {
		return "Initializing..."
	}

	// Apply appropriate styles based on active pane
	workspaceView := workspaceStyle
	taskView := taskStyle

	if m.activePane == 0 {
		workspaceView = focusedWorkspaceStyle
	} else {
		taskView = focusedTaskStyle
	}

	// Calculate dimensions accounting for borders and padding
	contentWidth := m.windowWidth - 4
	contentHeight := m.windowHeight - 4

	workspaceWidth := contentWidth/4 - 2
	taskWidth := contentWidth - workspaceWidth - 6

	// Render components with proper dimensions
	workspaceView = workspaceView.Width(workspaceWidth).Height(contentHeight - 2)
	taskView = taskView.Width(taskWidth).Height(contentHeight - 2)

	// Combine views horizontally with proper spacing
	layout := lipgloss.JoinHorizontal(
		lipgloss.Top,
		workspaceView.Render(m.workspaces.View()),
		taskView.Render(m.tasks.View()),
	)

	// Add help text at the bottom
	help := statusBar.Render("j/k: navigate • tab: switch pane • enter: select • q: quit")

	// Apply overall app padding and return
	return appStyle.Render(
		lipgloss.JoinVertical(lipgloss.Left, layout, help),
	)
}
