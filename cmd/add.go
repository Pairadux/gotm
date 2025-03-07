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
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Pairadux/gotm/internal/models"
	"github.com/Pairadux/gotm/internal/storage"
	"github.com/Pairadux/gotm/internal/taskops"
	"github.com/Pairadux/gotm/internal/utility"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
) // }}}

// addCmd represents the add command
// NOTE: This command will only deal with `active` tasks
var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Args:    cobra.MinimumNArgs(1),
	Short:   "Add a task to Gotm",
	Long:    `Add a task to Gotm with some other information listed as well`,
	Run: func(cmd *cobra.Command, args []string) {
		utility.DebugMessage(fmt.Sprintf("Add called"))

		active := taskops.InitActive()
		workspace, err := cmd.Flags().GetString("workspace")
		cobra.CheckErr(err)

		if err := utility.ValidateWorkspace(active, workspace); err != nil {
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("Workspace '%s', not found. Create it? (y/n): ", workspace)
			input, _ := reader.ReadString('\n')
			input = strings.ToLower(strings.TrimSpace(input))
			if input == "yes" || input == "y" || input == "" {
				active.Workspaces[workspace] = &models.Workspace{
					Name:         workspace,
					LastModified: time.Now(),
					Tasks:        []models.Task{},
				}
				fmt.Printf("Created workspace '%s'.\n", workspace)
			} else {
				fmt.Printf("Workspace creation aborted.\n")
				fmt.Printf("You can manually create a workspace with 'gotm workspace create <workspace>\n")
				os.Exit(0)
			}
		}

		taskops.Add(&active.Workspaces[workspace].Tasks, strings.Join(args, " "))

		active_path := viper.GetString("active_path")
		storage.SaveTasksToFile(active_path, active)
		utility.DebugMessage(fmt.Sprintf("Tasks saved to json file: %s", active_path))
	},
}

func init() { // {{{
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	addCmd.Flags().StringP("workspace", "w", "inbox", "workspace to use (default is inbox)")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
} // }}}
