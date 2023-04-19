package app

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
