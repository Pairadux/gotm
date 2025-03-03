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
	"path/filepath"

	"github.com/Pairadux/gotm/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
) // }}}

var (
	cfgFile   string
	workspace string

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "gotm",
		Short: "A TUI Task Manager",
		Long:  `A Vim-like TUI Task Manager with deep integration for so and so`,

		Run: func(cmd *cobra.Command, args []string) {
			// DEBUGGING
			if len(os.Getenv("DEBUG")) > 0 {
				f, err := tea.LogToFile("debug.log", "debug")
				if err != nil {
					fmt.Println("fatal:", err)
					os.Exit(1)
				}
				defer f.Close()
			}

			// TUI PROGRAM
			p := tea.NewProgram(tui.InitModel(resolveWorkspace(cmd)))
			if _, err := p.Run(); err != nil {
				fmt.Printf("Alas, there's been an error: %v", err)
				os.Exit(1)
			}
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() { // {{{
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
} // }}}

// Here you will define your flags and configuration settings.
// Cobra supports persistent flags, which, if defined here,
// will be global for your application.
func init() { // {{{
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gotm.yaml)")
	rootCmd.PersistentFlags().StringVar(&workspace, "workspace", "", "workspace to use (default is inbox)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
} // }}}

// initConfig reads in config file and ENV variables if set.
func initConfig() { // {{{

	dataDir := "/Users/austingause/.local/share/"
	appDataDir := filepath.Join(dataDir, "gotm")

	configDir := "/Users/austingause/.config/"
	appConfigDir := filepath.Join(configDir, "gotm")

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		if _, err := os.Stat(appConfigDir); os.IsNotExist(err) {
			cobra.CheckErr(os.MkdirAll(appConfigDir, 0o755))
		}

		if _, err := os.Stat(appDataDir); os.IsNotExist(err) {
			cobra.CheckErr(os.MkdirAll(appDataDir, 0o755))
		}

		viper.AddConfigPath(appConfigDir)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	defaultJSONPath := filepath.Join(appDataDir, "tasks.json")
	viper.SetDefault("json_path", defaultJSONPath)
	viper.SetDefault("default_workspace", "inbox")
	// TODO: Add option for default sorting method

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if cfgFile == "" {
				configFilePath := filepath.Join(appConfigDir, "config.yaml")

				fmt.Println("Config file not found, creating default config file...")
				cobra.CheckErr(viper.SafeWriteConfigAs(configFilePath))
				fmt.Printf("Created default config file at: %s\n", configFilePath)

				cobra.CheckErr(viper.ReadInConfig())
			}
		} else {
			cobra.CheckErr(err)
		}
	}
	debugMessage(fmt.Sprintf("Using config file: %s\n\n", viper.ConfigFileUsed()))
} // }}}
