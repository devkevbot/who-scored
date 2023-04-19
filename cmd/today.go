package cmd

import (
	"fmt"
	"os"

	"github.com/devkevbot/who-scored/internal/pkg/api"
	"github.com/devkevbot/who-scored/internal/pkg/utils"
	"github.com/spf13/cobra"
)

// todayCmd represents the today command
var todayCmd = &cobra.Command{
	Use:   "today",
	Args:  cobra.NoArgs,
	Short: "Find scores for NHL games scheduled for today",
	Run: func(cmd *cobra.Command, args []string) {
		schedule, err := api.GetScheduleForToday()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		utils.FormatScheduleAsTable(schedule)
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
}
