package app

import (
	"fmt"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Schedule struct {
	Copyright    string   `json:"copyright"`
	TotalItems   int64    `json:"totalItems"`
	TotalEvents  int64    `json:"totalEvents"`
	TotalGames   int64    `json:"totalGames"`
	TotalMatches int64    `json:"totalMatches"`
	MetaData     MetaData `json:"metaData"`
	Wait         int64    `json:"wait"`
	Dates        []Date   `json:"dates"`
}

func (schedule *Schedule) String() string {
	scheduleTable := table.NewWriter()
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

	return scheduleTable.Render()
}

func formatIsoDateAsLocalTime(isoDate string) (*string, error) {
	parsedIsoDate, err := time.Parse(time.RFC3339, isoDate)
	if err != nil {
		return nil, err
	}

	localTime := parsedIsoDate.Local().Format("02 Jan 06 15:04 MST")
	return &localTime, nil
}

type Date struct {
	Date         string        `json:"date"`
	TotalItems   int64         `json:"totalItems"`
	TotalEvents  int64         `json:"totalEvents"`
	TotalGames   int64         `json:"totalGames"`
	TotalMatches int64         `json:"totalMatches"`
	Games        []Game        `json:"games"`
	Events       []interface{} `json:"events"`
	Matches      []interface{} `json:"matches"`
}

type Game struct {
	GamePk   int64   `json:"gamePk"`
	Link     string  `json:"link"`
	GameType string  `json:"gameType"`
	Season   string  `json:"season"`
	GameDate string  `json:"gameDate"`
	Status   Status  `json:"status"`
	Teams    Teams   `json:"teams"`
	Venue    Venue   `json:"venue"`
	Content  Content `json:"content"`
}

type Content struct {
	Link string `json:"link"`
}

type Status struct {
	AbstractGameState string `json:"abstractGameState"`
	CodedGameState    string `json:"codedGameState"`
	DetailedState     string `json:"detailedState"`
	StatusCode        string `json:"statusCode"`
	StartTimeTBD      bool   `json:"startTimeTBD"`
}

type Teams struct {
	Away Away `json:"away"`
	Home Away `json:"home"`
}

type Away struct {
	LeagueRecord LeagueRecord `json:"leagueRecord"`
	Score        int64        `json:"score"`
	Team         Venue        `json:"team"`
}

type LeagueRecord struct {
	Wins   int64  `json:"wins"`
	Losses int64  `json:"losses"`
	Type   string `json:"type"`
}

type Venue struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

type MetaData struct {
	TimeStamp string `json:"timeStamp"`
}
