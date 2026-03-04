package cmd

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/devkevbot/who-scored/internal/api"
	"github.com/devkevbot/who-scored/internal/app"
	"github.com/spf13/cobra"
)

var onCmd = &cobra.Command{
	Use:   "on <date>",
	Args:  validateArgsOn,
	Short: "Find scores for NHL games scheduled for a specific date",
	RunE: func(cmd *cobra.Command, args []string) error {
		if teamFilter != "" && !app.IsValidTeam(teamFilter) {
			return fmt.Errorf("invalid team abbreviation %q; valid teams: %s", teamFilter, strings.Join(app.ValidTeamAbbrevs(), ", "))
		}
		inputDate := args[0]
		scores, err := api.GetScoresForSingleDay(inputDate)
		if err != nil {
			log.Fatal(err)
		}
		if teamFilter != "" {
			scores.FilterByTeam(teamFilter)
		}
		fmt.Println(scores)
		return nil
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
