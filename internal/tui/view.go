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
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

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

// handleWindowResize updates component dimensions when window size changes
func handleWindowResize(m Model, msg tea.WindowSizeMsg) Model {
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
	
	return m
}
