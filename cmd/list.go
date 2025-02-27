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

	"github.com/Pairadux/gotm/internal/storage"
	"github.com/Pairadux/gotm/internal/tasks"

	"github.com/spf13/cobra"
) // }}}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List items",
	Long:    `List Items`,
	Run: func(cmd *cobra.Command, args []string) {
		tm, err := storage.LoadTasks()
		if err != nil {
			panic(err)
		}

		tasks.AddTask(tm, "This is task 1")
		tasks.AddTask(tm, "This is task 2")
		tasks.AddTask(tm, "This is task 3")
		tasks.AddTask(tm, "This is task 4")

		fmt.Printf("INDEX\t| ID\t| Description\n")
		for i, e := range(tm.Tasks) {
			fmt.Printf("%d\t  %d\t  %s\n", i, e.Id, e.Description)
		}

		tasks.FindFirstAvailableId(tm)

		fmt.Printf("INDEX\t| ID\t| Description\n")
		for i, e := range(tm.Tasks) {
			fmt.Printf("%d\t  %d\t  %s\n", i, e.Id, e.Description)
		}
	},
}

func init() { // {{{
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
} // }}}
