package cmd

import (
	"fmt"
	"log"

	"github.com/devkevbot/who-scored/internal/api"

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
			log.Fatal(err)
		}
		fmt.Println(schedule)
	},
}

func init() {
	rootCmd.AddCommand(yesterdayCmd)
}
