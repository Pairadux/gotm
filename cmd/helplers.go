package cmd

// IMPORTS {{{
import (
	"fmt"
	"os"

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
