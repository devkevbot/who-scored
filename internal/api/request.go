package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/devkevbot/who-scored/internal/app"
)

func GetScoresForToday() (*app.DailyScores, error) {
	today := getTodayYearMonthDay()

	scores, err := getScores(today)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch scores for today: %w", errors.Unwrap(err))
	}
	return scores, nil
}

func GetScoresForYesterday() (*app.DailyScores, error) {
	yesterday := getYesterdayYearMonthDay()

	scores, err := getScores(yesterday)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch scores for yesterday: %w", errors.Unwrap(err))
	}
	return scores, nil
}

func GetScoresForSingleDay(date string) (*app.DailyScores, error) {
	scores, err := getScores(date)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch scores for date: %w", errors.Unwrap(err))
	}
	return scores, nil
}

func getScores(date string) (*app.DailyScores, error) {
	req, err := http.NewRequest("GET", nhlScoresApiPath(date), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	var scores app.DailyScores
	err = json.NewDecoder(resp.Body).Decode(&scores)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %w", err)
	}
	return &scores, nil
}

func nhlScoresApiPath(date string) string {
	base := "https://api-web.nhle.com/v1/score/"
	return base + date
}

func getTodayYearMonthDay() string {
	today := time.Now()

	year := today.Year()
	month := fmt.Sprintf("%02d", today.Month())
	day := fmt.Sprintf("%02d", today.Day())

	return fmt.Sprintf("%d-%v-%v", year, month, day)

}

func getYesterdayYearMonthDay() string {
	today := time.Now()

	yesterday := today.Add(time.Hour * -24)

	year := yesterday.Year()
	month := fmt.Sprintf("%02d", yesterday.Month())
	day := fmt.Sprintf("%02d", yesterday.Day())

	return fmt.Sprintf("%d-%v-%v", year, month, day)
}
