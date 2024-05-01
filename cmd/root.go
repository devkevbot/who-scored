package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const CliName = "who-scored"

var rootCmd = &cobra.Command{
	Use:   CliName,
	Short: fmt.Sprintf("%s finds the score of NHL games for a specific date", CliName),
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
