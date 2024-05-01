package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/devkevbot/who-scored/internal/api"
)

var yesterdayCmd = &cobra.Command{
	Use:   "yesterday",
	Args:  cobra.NoArgs,
	Short: "Find scores for NHL games scheduled for yesterday",
	Run: func(cmd *cobra.Command, args []string) {
		scores, err := api.GetScoresForYesterday()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(scores)
	},
}

func init() {
	rootCmd.AddCommand(yesterdayCmd)
}
