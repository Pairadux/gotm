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
	"os"

	"github.com/Pairadux/gotm/internal/models"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)// }}}

func resolveWorkspace(cmd *cobra.Command) string {
	if cmd.Flags().Changed("workspace") {
		return cmd.Flag("workspace").Value.String()
	}
	return viper.GetString("default_workspace")
}

func debugMessage(m string) {
	if len(os.Getenv("DEBUG")) > 0 {
		fmt.Printf("%s\n", m)
	}
}

func ValidateWorkspace(active *models.TaskState, workspace string) error {
	if active.Workspaces == nil || active.Workspaces[workspace] == nil {
		return fmt.Errorf("workspace %q not found", workspace)
	}
	return nil
}
