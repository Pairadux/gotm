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
	"time"

	"github.com/Pairadux/gotm/internal/models"
	"github.com/Pairadux/gotm/internal/storage"
	"github.com/Pairadux/gotm/internal/taskops"
	"github.com/Pairadux/gotm/internal/utility"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
) // }}}

// completeCmd represents the complete command
// NOTE: This command will deal with both `active` and `completed` tasks
// Given that it has to move the completed active task to the completed workspace
var completeCmd = &cobra.Command{
	Use:     "complete",
	Aliases: []string{"c"},
	Args:    cobra.ExactArgs(1),
	Short:   "Mark a task as completed in Gotm",
	Long:    `Mark a task as completed in Gotm with some other information listed as well`,
	Run: func(cmd *cobra.Command, args []string) {
		utility.DebugMessage(fmt.Sprintf("Complete called"))

		all := taskops.InitAll()
		workspace, err := cmd.Flags().GetString("workspace")
		cobra.CheckErr(err)

		err = utility.ValidateWorkspace(all.Active, workspace)
		cobra.CheckErr(err)

		if err = utility.ValidateWorkspace(all.Completed, workspace); err != nil {
			all.Completed.Workspaces[workspace] = &models.Workspace{
				Name:         all.Active.Workspaces[workspace].Name,
				LastModified: time.Now(),
				Tasks:        []models.Task{},
			}
		}

		found := false

		i, err := strconv.Atoi(args[0])
		cobra.CheckErr(err)

		all, found = taskops.Complete(all, workspace, i)
		if !found {
			fmt.Printf("Task with index %d does not exist.\n", i)
		}

		storage.SaveTasksToFile(viper.GetString("active_path"), all.Active)
		utility.DebugMessage(fmt.Sprintf("Active tasks saved to json file: %s", viper.GetString("active_path")))

		storage.SaveTasksToFile(viper.GetString("completed_path"), all.Completed)
		utility.DebugMessage(fmt.Sprintf("Completed tasks saved to json file: %s", viper.GetString("completed_path")))
	},
}

func init() { // {{{
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	completeCmd.Flags().StringP("workspace", "w", "inbox", "workspace to use (default is inbox)")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
} // }}}
