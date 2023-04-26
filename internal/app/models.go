package app

import (
	"fmt"
	"strconv"
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
	t := table.NewWriter()
	t.AppendHeader(table.Row{
		"GAME TYPE",
		"START TIME",
		"AWAY TEAM (RECORD)",
		"SCORE",
		"HOME TEAM (RECORD)",
		"SCORE",
		"STATUS",
	})
	t.SuppressEmptyColumns()

	for _, date := range schedule.Dates {
		for _, game := range date.Games {
			localTime, err := formatIsoDateAsLocalTime(game.GameDate)
			if err != nil {
				fmt.Println("Couldn't parse game start time!", err)
			}
			if localTime == nil {
				localTime = new(string)
			}

			t.AppendRow(table.Row{
				game.gameTypeCol(),
				*localTime,
				game.awayTeamCol(),
				game.awayTeamScoreCol(),
				game.homeTeamCol(),
				game.homeTeamScoreCol(),
				game.statusCol(),
			})
		}
	}

	return t.Render()
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

func (game *Game) gameTypeCol() string {
	return parseGameType(game.GameType)
}

func (game *Game) isPlayoffGame() bool {
	return parseGameType(game.GameType) == "Playoffs"
}

func (game *Game) statusCol() string {
	return game.Status.DetailedState
}

func (game *Game) awayTeamCol() string {
	if game.isPlayoffGame() {
		return fmt.Sprintf(
			"%s (%s-%s)",
			game.awayTeamName(),
			game.awayTeamWins(),
			game.awayTeamLosses(),
		)
	}

	return fmt.Sprintf(
		"%s (%s-%s-%s)",
		game.awayTeamName(),
		game.awayTeamWins(),
		game.awayTeamLosses(),
		game.awayTeamOvertimeLosses(),
	)
}

func (game *Game) awayTeamName() string {
	return game.Teams.Away.Team.Name
}

func (game *Game) awayTeamWins() string {
	return strconv.FormatInt(game.Teams.Away.LeagueRecord.Wins, 10)
}

func (game *Game) awayTeamLosses() string {
	return strconv.FormatInt(game.Teams.Away.LeagueRecord.Losses, 10)
}

func (game *Game) awayTeamOvertimeLosses() string {
	if game.isPlayoffGame() {
		return ""
	}
	return strconv.FormatInt(game.Teams.Away.LeagueRecord.Overtime, 10)
}

func (game *Game) awayTeamScoreCol() string {
	return strconv.FormatInt(game.Teams.Away.Score, 10)
}

func (game *Game) homeTeamCol() string {
	if game.isPlayoffGame() {
		return fmt.Sprintf(
			"%s (%s-%s)",
			game.homeTeamName(),
			game.homeTeamWins(),
			game.homeTeamLosses(),
		)
	}

	return fmt.Sprintf(
		"%s (%s-%s-%s)",
		game.homeTeamName(),
		game.homeTeamWins(),
		game.homeTeamLosses(),
		game.homeTeamOvertimeLosses(),
	)
}

func (game *Game) homeTeamName() string {
	return game.Teams.Home.Team.Name
}

func (game *Game) homeTeamWins() string {
	return strconv.FormatInt(game.Teams.Home.LeagueRecord.Wins, 10)
}

func (game *Game) homeTeamLosses() string {
	return strconv.FormatInt(game.Teams.Home.LeagueRecord.Losses, 10)
}

func (game *Game) homeTeamOvertimeLosses() string {
	if game.isPlayoffGame() {
		return ""
	}
	return strconv.FormatInt(game.Teams.Home.LeagueRecord.Overtime, 10)
}

func (game *Game) homeTeamScoreCol() string {
	return strconv.FormatInt(game.Teams.Home.Score, 10)
}

var gameTypeAbbrToFullName = map[string]string{
	"P":  "Playoffs",
	"R":  "Regular Season",
	"PR": "Pre-season",
}

func parseGameType(gameType string) string {
	return gameTypeAbbrToFullName[gameType]
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
	Wins     int64  `json:"wins"`
	Losses   int64  `json:"losses"`
	Overtime int64  `json:"ot"`
	Type     string `json:"type"`
}

type Venue struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

type MetaData struct {
	TimeStamp string `json:"timeStamp"`
}
