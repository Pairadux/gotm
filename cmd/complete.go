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

	"github.com/Pairadux/gotm/internal/taskops"
	"github.com/Pairadux/gotm/internal/storage"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
) // }}}

// completeCmd represents the complete command
// NOTE: This command will deal with both `active` and `completed` tasks
// Given that it has to move the completed active task to the completed workspace
var completeCmd = &cobra.Command{
	Use:     "complete",
	Aliases: []string{"c"},
	Short:   "Mark a task as completed in Gotm",
	Long:    `Mark a task as completed in Gotm with some other information listed as well`,
	Run: func(cmd *cobra.Command, args []string) {
		debugMessage(fmt.Sprintf("Remove called"))

		if len(args) == 0 {
			_ = cmd.Help()
			return
		}

		active := taskops.InitActive()
		workspace := resolveWorkspace(cmd)

		found := false

		i, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		found = taskops.Complete(&active[workspace].Tasks, i)
		if !found {
			fmt.Printf("Task with index %d does not exist.\n", i)
		}

		storage.SaveTasksToFile(viper.GetString("active_path"), active)
		debugMessage(fmt.Sprintf("Active tasks saved to json file: %s", viper.GetString("active_path")))

		// storage.SaveTasksToFile(viper.GetString("completed_path"), )
		// debugMessage(fmt.Sprintf("Completed tasks saved to json file: %s", viper.GetString("completed_path")))
	},
}

func init() { // {{{
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
} // }}}
