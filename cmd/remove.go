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

package cmd

// IMPORTS {{{
import (
	"fmt"
	"strconv"

	"github.com/Pairadux/gotm/internal/models"
	"github.com/Pairadux/gotm/internal/storage"
	"github.com/Pairadux/gotm/internal/taskops"
	"github.com/Pairadux/gotm/internal/utility"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
) // }}}

// removeCmd represents the remove command
// NOTE: This command can remove `active` or `completed` tasks
var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"r"},
	Args:    cobra.ExactArgs(1),
	Short:   "Remove a task from Gotm",
	Long:    `Remove a task from Gotm with some other information listed as well`,
	Run: func(cmd *cobra.Command, args []string) {
		utility.DebugMessage(fmt.Sprintf("Remove called"))

		workspace, err := cmd.Flags().GetString("workspace")
		cobra.CheckErr(err)
		var taskState *models.TaskState
		var path string

		isCompleted, err := cmd.Flags().GetBool("completed")
		cobra.CheckErr(err)

		if isCompleted {
			taskState = taskops.InitCompleted()
			path = viper.GetString("completed_path")
		} else {
			taskState = taskops.InitActive()
			path = viper.GetString("active_path")
		}

		err = utility.ValidateWorkspace(taskState, workspace)
		cobra.CheckErr(err)

		taskRemoved := models.Task{}
		found := false

		i, err := strconv.Atoi(args[0])
		cobra.CheckErr(err)

		taskRemoved, found = taskops.Remove(&taskState.Workspaces[workspace].Tasks, i)
		if found {
			fmt.Printf("Task removed: %s\n", taskRemoved.Description)
		} else {
			fmt.Printf("Task with index %d does not exist.\n", i)
		}

		storage.SaveTasksToFile(path, taskState)
		utility.DebugMessage(fmt.Sprintf("Tasks saved to json file: %s", path))
	},
}

func init() { // {{{
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	removeCmd.Flags().StringP("workspace", "w", "inbox", "workspace to use (default is inbox)")
	removeCmd.Flags().BoolP("completed", "c", false, "Remove a completed task")
} // }}}
