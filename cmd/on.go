package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/devkevbot/who-scored/internal/api"
	"github.com/devkevbot/who-scored/internal/app"
	"github.com/spf13/cobra"
)

var onCmd = &cobra.Command{
	Use:   "on <date>",
	Args:  validateArgsOn,
	Short: "Find scores for NHL games scheduled for a specific date",
	Run: func(cmd *cobra.Command, args []string) {
		inputDate := args[0]
		scores, err := api.GetScoresForSingleDay(inputDate)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(scores)
	},
}

func init() {
	rootCmd.AddCommand(onCmd)
}

func validateArgsOn(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("expected one argument: a date")
	}

	return app.ParseUserProvidedDate(args[0])
}
