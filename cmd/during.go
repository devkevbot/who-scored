package cmd

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/devkevbot/who-scored/internal/pkg/api"
	"github.com/spf13/cobra"
)

// duringCmd represents the during command
var duringCmd = &cobra.Command{
	Use:   "during <start date> <end date>",
	Args:  validateArgsDuring,
	Short: "Find scores for NHL games scheduled between a start date and an end date. Dates must be in YYYY-MM-DD format",
	Run: func(cmd *cobra.Command, args []string) {
		schedule, err := api.GetScheduleForDateRange(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
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

	startDate, startDateErr := time.Parse(time.DateOnly, args[0])
	endDate, endDateErr := time.Parse(time.DateOnly, args[1])

	// Parsing errors
	if startDateErr != nil && endDateErr != nil {
		return errors.New("couldn't parse the start and end dates to YYYY-MM-DD format")
	}
	if startDateErr != nil {
		return errors.New("couldn't parse the start date to YYYY-MM-DD format")
	}
	if endDateErr != nil {
		return errors.New("couldn't parse the end date to YYYY-MM-DD format")
	}

	now := time.Now()

	// Input range error
	if startDate.After(now) && endDate.After(now) {
		return errors.New("both the start and end dates occur in the future")
	}
	if startDate.After(now) {
		return errors.New("the start occurs in the future")
	}
	if endDate.After(now) {
		return errors.New("the end occurs in the future")
	}
	if startDate.After(endDate) {
		return errors.New("the start date occurs after the end date")
	}

	endDateLimit := startDate.Add(time.Hour * 24 * 7)
	if endDate.After(endDateLimit) {
		suppliedDateRangeInDays := endDate.Sub(startDate).Hours() / 24
		return fmt.Errorf("the maximum allowed date range is 7 days, but you requested %v days", suppliedDateRangeInDays)
	}

	return nil
}
