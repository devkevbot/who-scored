package app

import (
	"sort"
	"strings"
)

var validTeams = map[string]bool{
	"ANA": true, "BOS": true, "BUF": true, "CAR": true,
	"CGY": true, "CHI": true, "COL": true, "CBJ": true,
	"DAL": true, "DET": true, "EDM": true, "FLA": true,
	"LAK": true, "MIN": true, "MTL": true, "NSH": true,
	"NJD": true, "NYI": true, "NYR": true, "OTT": true,
	"PHI": true, "PIT": true, "SEA": true, "SJS": true,
	"STL": true, "TBL": true, "TOR": true, "UTA": true,
	"VAN": true, "VGK": true, "WPG": true, "WSH": true,
}

func IsValidTeam(abbrev string) bool {
	return validTeams[strings.ToUpper(abbrev)]
}

func ValidTeamAbbrevs() []string {
	teams := make([]string, 0, len(validTeams))
	for t := range validTeams {
		teams = append(teams, t)
	}
	sort.Strings(teams)
	return teams
}
