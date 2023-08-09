package cmd

import (
	"fmt"
	"os"

	"github.com/devkevbot/who-scored/internal/pkg/api"
	"github.com/spf13/cobra"
)

// yesterdayCmd represents the yesterday command
var yesterdayCmd = &cobra.Command{
	Use:   "yesterday",
	Args:  cobra.NoArgs,
	Short: "Find scores for NHL games scheduled for yesterday",
	Run: func(cmd *cobra.Command, args []string) {
		schedule, err := api.GetScheduleForYesterday()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(schedule)
	},
}

func init() {
	rootCmd.AddCommand(yesterdayCmd)
}
