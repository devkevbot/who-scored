package app

import (
	"fmt"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type DailyScores struct {
	Games []Game `json:"games"`
}

func (ds *DailyScores) String() string {
	if len(ds.Games) == 0 {
		return "No game data was available for the requested date(s)."
	}

	return ds.toTable()
}

func (ds *DailyScores) toTable() string {
	t := table.NewWriter()
	t.SuppressEmptyColumns()

	t.AppendHeader(table.Row{
		"GAME TYPE",
		"PLAYOFF GAME",
		"PLAYOFF SERIES",
		"START TIME",
		"TEAMS",
		"SCORE",
		"STATUS",
		"GAME-WINNING GOAL",
	})

	for _, game := range ds.Games {
		t.AppendRow(game.toRow())
	}

	return t.Render()
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

func (g *Game) toRow() table.Row {
	return table.Row{
		g.gameTypeCol(),
		g.playoffGameCol(),
		g.playoffSeriesCol(),
		g.startTimeCol(),
		g.teamsCol(),
		g.scoreCol(),
		g.statusCol(),
		g.gwgCol(),
	}
}

func (g *Game) gameTypeCol() string {
	gameTypeToName := map[int]string{
		1: "Preseason",
		2: "Regular Season",
		3: "Playoffs",
	}
	return gameTypeToName[g.GameType]
}

func (g *Game) isPlayoffGame() bool {
	return g.SeriesStatus.Round >= 1
}

func (g *Game) playoffGameCol() string {
	if !g.isPlayoffGame() {
		return ""
	}

	return fmt.Sprintf("R%d, GM %d", g.SeriesStatus.Round, g.SeriesStatus.GameNumberOfSeries)
}

func (g *Game) playoffSeriesCol() string {
	if !g.isPlayoffGame() {
		return ""
	}

	// Series is tied
	if g.SeriesStatus.TopSeedWins == g.SeriesStatus.BottomSeedWins {
		return fmt.Sprintf("TIED %d-%d", g.SeriesStatus.TopSeedWins, g.SeriesStatus.BottomSeedWins)
	}
	// Top seed is leading
	if g.SeriesStatus.TopSeedWins > g.SeriesStatus.BottomSeedWins {
		return fmt.Sprintf("%s %d-%d", g.SeriesStatus.TopSeedTeamAbbrev, g.SeriesStatus.TopSeedWins, g.SeriesStatus.BottomSeedWins)
	}
	// Bototm seed is leading
	return fmt.Sprintf("%s %d-%d", g.SeriesStatus.BottomSeedTeamAbbrev, g.SeriesStatus.BottomSeedWins, g.SeriesStatus.TopSeedWins)
}

func (g *Game) startTimeCol() string {
	return g.StartTimeUTC.Local().Format("02 Jan 06 15:04 MST")
}

func (g *Game) awayAbbrev() string {
	return g.AwayTeam.Abbrev
}

func (g *Game) awayScore() int {
	return g.AwayTeam.Score
}

func (g *Game) homeAbbrev() string {
	return g.HomeTeam.Abbrev
}

func (g *Game) homeScore() int {
	return g.HomeTeam.Score
}

func (g *Game) teamsCol() string {
	awayAbbrev := g.awayAbbrev()
	homeAbbrev := g.homeAbbrev()
	team := awayAbbrev + " at " + homeAbbrev
	return team
}

func (g *Game) scoreCol() string {
	awayScore := g.awayScore()
	homeScore := g.homeScore()
	awayAbbrev := g.awayAbbrev()
	homeAbbrev := g.homeAbbrev()

	if awayScore > homeScore {
		return fmt.Sprintf("%d-%d %s", awayScore, homeScore, awayAbbrev)
	}
	if homeScore > awayScore {
		return fmt.Sprintf("%d-%d %s", homeScore, awayScore, homeAbbrev)
	}

	return fmt.Sprintf("%d-%d", awayScore, homeScore)
}

func (g *Game) isOver() bool {
	return g.GameState == "FINAL" || g.GameState == "OFF"
}

func (g *Game) statusCol() string {
	gameStateToDesc := map[string]string{
		"CRIT":  "In progress",
		"LIVE":  "In progress",
		"FUT":   "Not started",
		"PRE":   "Pre-game",
		"OFF":   "Final",
		"FINAL": "Final",
	}

	var output []string

	gameStateDesc := gameStateToDesc[g.GameState]
	output = append(output, gameStateDesc)

	if g.PeriodDescriptor.PeriodType == "OT" {
		otNum := g.PeriodDescriptor.Number % 3
		if otNum == 1 {
			output = append(output, "(OT)")
		} else {
			output = append(output, fmt.Sprintf("(%dOT)", otNum))
		}
	} else if g.PeriodDescriptor.PeriodType == "SO" {
		output = append(output, "(SO)")
	}

	return strings.Join(output, " ")
}

func (g *Game) gwgCol() string {
	if !g.isOver() {
		return ""
	}

	if len(g.Goals) == 0 {
		return "(No Score)"
	}

	var gwg Goal
	awayScore := g.awayScore()
	homeScore := g.homeScore()
	if awayScore > homeScore {
		for _, goal := range g.Goals {
			if goal.AwayScore == homeScore+1 {
				gwg = goal
				break
			}
		}
	} else if homeScore > awayScore {
		for _, goal := range g.Goals {
			if goal.HomeScore == awayScore+1 {
				gwg = goal
				break
			}
		}
	}

	goalScorer := gwg.Name.Default

	goalStrengthToDesc := map[string]string{
		"pp": "PPG",
		"sh": "SHG",
	}
	goalStrength := gwg.Strength

	output := []string{goalScorer}

	gameDecidedInShootout := g.PeriodDescriptor.PeriodType == "SO"
	if !gameDecidedInShootout {
		goalNumber := gwg.GoalsToDate
		output = append(output, fmt.Sprintf("(%d)", goalNumber))
	}
	if goalStrength != "ev" {
		output = append(output, goalStrengthToDesc[gwg.Strength])
	}

	if len(gwg.Assists) > 0 {
		var nameAndNumber []string
		for _, assist := range gwg.Assists {
			playerName := assist.Name.Default
			assistNumber := assist.AssistsToDate
			nameAndNumber = append(nameAndNumber, fmt.Sprintf("%s (%d)", playerName, assistNumber))
		}
		assistsStr := strings.Join(nameAndNumber, ", ")
		output = append(output, "from")
		output = append(output, assistsStr)
	}

	return strings.Join(output, " ")
}
