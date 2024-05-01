package cmd

import (
	"fmt"
	"log"

	"github.com/devkevbot/who-scored/internal/api"

	"github.com/spf13/cobra"
)

var todayCmd = &cobra.Command{
	Use:   "today",
	Args:  cobra.NoArgs,
	Short: "Find scores for NHL games scheduled for today",
	Run: func(cmd *cobra.Command, args []string) {
		scores, err := api.GetScoresForToday()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(scores)
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
}
