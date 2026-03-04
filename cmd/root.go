package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const CliName = "who-scored"

var teamFilter string

var rootCmd = &cobra.Command{
	Use:   CliName,
	Short: fmt.Sprintf("%s finds the score of NHL games for a specific date", CliName),
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&teamFilter, "team", "t", "", "filter by team abbreviation (e.g. TOR, BOS)")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
