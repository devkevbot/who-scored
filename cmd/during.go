package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/devkevbot/who-scored/internal/api"
	"github.com/spf13/cobra"
)

// duringCmd represents the during command
var duringCmd = &cobra.Command{
	Use:   "during <start date> <end date>",
	Args:  validateArgsDuring,
	Short: "Find scores for NHL games scheduled between a start date and an end date",
	Run: func(cmd *cobra.Command, args []string) {
		schedule, err := api.GetScheduleForDateRange(args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(schedule)
	},
}

func init() {
	rootCmd.AddCommand(duringCmd)
}

func validateArgsDuring(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return errors.New("expected two arguments: a start date and an end date")
	}
	return nil
}
