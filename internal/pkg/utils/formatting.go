package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/devkevbot/who-scored/internal/app"
	"github.com/jedib0t/go-pretty/v6/table"
)

// Pretty-prints the schedule in a tabular format.
func FormatScheduleAsTable(schedule *app.Schedule) {
	scheduleTable := table.NewWriter()
	scheduleTable.SetOutputMirror(os.Stdout)
	scheduleTable.AppendHeader(table.Row{"START TIME", "AWAY TEAM (W-L)", "SCORE", "HOME TEAM (W-L)", "SCORE", "STATUS"})

	for _, date := range schedule.Dates {
		for _, game := range date.Games {
			localTime, err := formatIsoDateAsLocalTime(game.GameDate)
			if err != nil {
				fmt.Println("Couldn't parse game start time!", err)
			}
			if localTime == nil {
				localTime = new(string)
			}

			scheduleTable.AppendRow(table.Row{
				*localTime,
				fmt.Sprintf("%s (%d-%d)",
					game.Teams.Away.Team.Name,
					game.Teams.Away.LeagueRecord.Wins,
					game.Teams.Away.LeagueRecord.Losses,
				),
				game.Teams.Away.Score,
				fmt.Sprintf("%s (%d-%d)",
					game.Teams.Home.Team.Name,
					game.Teams.Home.LeagueRecord.Wins,
					game.Teams.Home.LeagueRecord.Losses,
				),
				game.Teams.Home.Score,
				game.Status.DetailedState,
			})
		}
	}

	scheduleTable.Render()
}

func formatIsoDateAsLocalTime(isoDate string) (*string, error) {
	parsedIsoDate, err := time.Parse(time.RFC3339, isoDate)
	if err != nil {
		return nil, err
	}

	localTime := parsedIsoDate.Local().Format("02 Jan 06 15:04 MST")
	return &localTime, nil
}

// Returns yesterday's year, month, and day formatted as YYYY-MM-DD
func GetYesterdayYearMonthDay() string {
	today := time.Now()

	yesterday := today.Add(time.Hour * -24)

	year := yesterday.Year()
	month := fmt.Sprintf("%02d", yesterday.Month())
	day := yesterday.Day()

	return fmt.Sprintf("%d-%v-%d", year, month, day)
}
