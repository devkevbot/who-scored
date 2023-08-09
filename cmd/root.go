package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const CLI_NAME = "who-scored"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   CLI_NAME,
	Short: fmt.Sprintf("%s finds the score of NHL games for a specific date", CLI_NAME),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
