package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"

	"github.com/devkevbot/who-scored/internal/api"
	"github.com/devkevbot/who-scored/internal/app"
)

var yesterdayCmd = &cobra.Command{
	Use:   "yesterday",
	Args:  cobra.NoArgs,
	Short: "Find scores for NHL games scheduled for yesterday",
	RunE: func(cmd *cobra.Command, args []string) error {
		if teamFilter != "" && !app.IsValidTeam(teamFilter) {
			return fmt.Errorf("invalid team abbreviation %q; valid teams: %s", teamFilter, strings.Join(app.ValidTeamAbbrevs(), ", "))
		}
		scores, err := api.GetScoresForYesterday()
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
	rootCmd.AddCommand(yesterdayCmd)
}
