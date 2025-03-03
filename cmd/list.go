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
	"strings"
	"time"

	"github.com/Pairadux/gotm/internal/storage"
	"github.com/Pairadux/gotm/internal/taskops"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
) // }}}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List items from Gotm",
	Long:    `List Items from Gotm with some other information listed as well`,
	Run: func(cmd *cobra.Command, args []string) {
		debugMessage(fmt.Sprintf("list called\n\n"))

		workspaces := taskops.InitWorkspaces()

		workspace := resolveWorkspace(cmd)

		debugMessage(fmt.Sprintf("Using workspace: %s\n\n", workspace))

		tasks := workspaces[workspace].Tasks

		sortType := "natural"
		if len(args) > 0 {
			sortType = args[0]
		}
		taskops.Sort(sortType, tasks)

		// OUTPUT FORMATTING
		r := strings.Repeat
		fmt.Printf("%-5s │ %-40s │ %s\n", "Index", "Description", "Created")
		fmt.Printf("%s┼%s┼%s\n", r("─", 6), r("─", 42), r("─", 20))
		for _, e := range tasks {
			fmt.Printf("%-5d │ %-40.40s │ %s\n", e.Index, e.Description, e.Created.Format(time.DateTime))
		}

		storage.SaveTasksToFile(viper.GetString("json_path"), workspaces)
		debugMessage(fmt.Sprintf("\nTasks saved to json file: %s", viper.GetString("json_path")))
	},
}

func init() { // {{{
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// NOTE: Possibly add sorting method as a flag rather than argument
	// This would allow the argument to be the workspace used, but thats also already a flag, so some more thought has to be put into this

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
} // }}}
