package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const CliName = "who-scored"

var teamFilter string
var shortOutput bool

var rootCmd = &cobra.Command{
	Use:   CliName,
	Short: fmt.Sprintf("%s finds the score of NHL games for a specific date", CliName),
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&teamFilter, "team", "t", "", "filter by team abbreviation (e.g. TOR, BOS)")
	rootCmd.PersistentFlags().BoolVarP(&shortOutput, "short", "s", false, "show compact output with just teams, score, and status")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
