package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/devkevbot/who-scored/internal/api"
	"github.com/devkevbot/who-scored/internal/app"

	"github.com/spf13/cobra"
)

var todayCmd = &cobra.Command{
	Use:   "today",
	Args:  cobra.NoArgs,
	Short: "Find scores for NHL games scheduled for today",
	RunE: func(cmd *cobra.Command, args []string) error {
		if teamFilter != "" && !app.IsValidTeam(teamFilter) {
			return fmt.Errorf("invalid team abbreviation %q; valid teams: %s", teamFilter, strings.Join(app.ValidTeamAbbrevs(), ", "))
		}
		scores, err := api.GetScoresForToday()
		if err != nil {
			log.Fatal(err)
		}
		if teamFilter != "" {
			scores.FilterByTeam(teamFilter)
		}
		scores.Short = shortOutput
		fmt.Println(scores)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
}
