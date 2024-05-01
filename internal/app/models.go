package app

import (
	"fmt"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type DailyScores struct {
	Games []Game `json:"games"`
}

type Name struct {
	Default string `json:"default"`
}

type Team struct {
	ID     int    `json:"id"`
	Name   Name   `json:"name"`
	Abbrev string `json:"abbrev"`
	Score  int    `json:"score"`
	Sog    int    `json:"sog"`
	Logo   string `json:"logo"`
}

type SeriesStatus struct {
	Round                int    `json:"round"`
	SeriesAbbrev         string `json:"seriesAbbrev"`
	SeriesLetter         string `json:"seriesLetter"`
	NeededToWin          int    `json:"neededToWin"`
	TopSeedTeamAbbrev    string `json:"topSeedTeamAbbrev"`
	TopSeedWins          int    `json:"topSeedWins"`
	BottomSeedTeamAbbrev string `json:"bottomSeedTeamAbbrev"`
	BottomSeedWins       int    `json:"bottomSeedWins"`
	GameNumberOfSeries   int    `json:"gameNumberOfSeries"`
}

type Clock struct {
	TimeRemaining    string `json:"timeRemaining"`
	SecondsRemaining int    `json:"secondsRemaining"`
	Running          bool   `json:"running"`
	InIntermission   bool   `json:"inIntermission"`
}

type PeriodDescriptor struct {
	Number     int    `json:"number"`
	PeriodType string `json:"periodType"`
}

type GameOutcome struct {
	LastPeriodType string `json:"lastPeriodType"`
	OtPeriods      int    `json:"otPeriods"`
}

type FirstName struct {
	Default string `json:"default"`
}

type LastName struct {
	Default string `json:"default"`
}

type Assists struct {
	PlayerID      int  `json:"playerId"`
	Name          Name `json:"name"`
	AssistsToDate int  `json:"assistsToDate"`
}

type Goal struct {
	Period           int              `json:"period"`
	PeriodDescriptor PeriodDescriptor `json:"periodDescriptor"`
	TimeInPeriod     string           `json:"timeInPeriod"`
	PlayerID         int              `json:"playerId"`
	Name             Name             `json:"name"`
	FirstName        FirstName        `json:"firstName"`
	LastName         LastName         `json:"lastName"`
	GoalModifier     string           `json:"goalModifier"`
	Assists          []Assists        `json:"assists"`
	Mugshot          string           `json:"mugshot"`
	TeamAbbrev       string           `json:"teamAbbrev"`
	GoalsToDate      int              `json:"goalsToDate"`
	AwayScore        int              `json:"awayScore"`
	HomeScore        int              `json:"homeScore"`
	Strength         string           `json:"strength"`
	HighlightClip    int64            `json:"highlightClip"`
	HighlightClipFr  int64            `json:"highlightClipFr"`
}

type Game struct {
	ID                int              `json:"id"`
	Season            int              `json:"season"`
	GameType          int              `json:"gameType"`
	GameDate          string           `json:"gameDate"`
	StartTimeUTC      time.Time        `json:"startTimeUTC"`
	EasternUTCOffset  string           `json:"easternUTCOffset"`
	VenueUTCOffset    string           `json:"venueUTCOffset"`
	GameState         string           `json:"gameState"`
	GameScheduleState string           `json:"gameScheduleState"`
	AwayTeam          Team             `json:"awayTeam"`
	HomeTeam          Team             `json:"homeTeam"`
	SeriesStatus      SeriesStatus     `json:"seriesStatus"`
	GameCenterLink    string           `json:"gameCenterLink"`
	SeriesURL         string           `json:"seriesUrl"`
	ThreeMinRecap     string           `json:"threeMinRecap"`
	ThreeMinRecapFr   string           `json:"threeMinRecapFr"`
	Clock             Clock            `json:"clock"`
	NeutralSite       bool             `json:"neutralSite"`
	VenueTimezone     string           `json:"venueTimezone"`
	Period            int              `json:"period"`
	PeriodDescriptor  PeriodDescriptor `json:"periodDescriptor"`
	GameOutcome       GameOutcome      `json:"gameOutcome,omitempty"`
	Goals             []Goal           `json:"goals"`
}

func (ds *DailyScores) String() string {
	if len(ds.Games) == 0 {
		return "No game data was available for the requested date(s)."
	}

	t := table.NewWriter()
	t.AppendHeader(table.Row{
		"GAME TYPE",
		"START TIME",
		"AWAY TEAM",
		"SCORE",
		"HOME TEAM",
		"SCORE",
	})
	t.SuppressEmptyColumns()

	for _, game := range ds.Games {
		t.AppendRow(game.toRow())
	}

	return t.Render()
}

func (g *Game) toRow() table.Row {
	return table.Row{
		g.getGameTypeName(),
		g.getUserLocalStartTime(),
		g.getAwayTeamName(),
		g.getAwayTeamScore(),
		g.getHomeTeamName(),
		g.getHomeTeamScore(),
	}
}

func (g *Game) getGameTypeName() string {
	var gameTypeToName = map[int]string{
		2: "Regular Season",
		3: "Playoffs",
	}
	return gameTypeToName[g.GameType]
}

func (g *Game) getUserLocalStartTime() string {
	return g.StartTimeUTC.Local().Format("02 Jan 06 15:04 MST")
}

func (g *Game) getAwayTeamName() string {
	return g.AwayTeam.Name.Default
}

func (g *Game) getAwayTeamScore() string {
	return fmt.Sprint(g.AwayTeam.Score)
}

func (g *Game) getHomeTeamName() string {
	return g.HomeTeam.Name.Default
}

func (g *Game) getHomeTeamScore() string {
	return fmt.Sprint(g.HomeTeam.Score)
}
