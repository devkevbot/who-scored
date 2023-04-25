package cmd

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/devkevbot/who-scored/internal/pkg/api"
	"github.com/spf13/cobra"
)

// onCmd represents the on command
var onCmd = &cobra.Command{
	Use:   "on",
	Args:  validateArgs,
	Short: "Find scores for NHL games scheduled for on the given date in YYYY-MM-DD format",
	Run: func(cmd *cobra.Command, args []string) {
		inputDate := args[0]
		schedule, err := api.GetScheduleForDate(inputDate)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(schedule)
	},
}

func init() {
	rootCmd.AddCommand(onCmd)
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("expected only a date in YYYY-MM-DD format, but received more arguments")
	}

	now := time.Now()
	inputDate, err := time.Parse(time.DateOnly, args[0])
	if err != nil {
		return errors.New("couldn't parse the input as a date in YYYY-MM-DD format")
	}

	if inputDate.After(now) {
		return errors.New("the input date occurs in the future")
	}

	return nil
}
