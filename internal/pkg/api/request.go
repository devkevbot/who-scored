package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/devkevbot/who-scored/internal/app"
)

func GetScheduleForToday() (*app.Schedule, error) {
	return getSchedule(nil)
}

func GetScheduleForYesterday() (*app.Schedule, error) {
	yesterday := getYesterdayYearMonthDay()
	dateRange := &DateRange{
		StartDate: yesterday,
		EndDate:   yesterday,
	}
	return getSchedule(dateRange)
}

func GetScheduleForDate(inputDate string) (*app.Schedule, error) {
	dateRange := &DateRange{
		StartDate: inputDate,
		EndDate:   inputDate,
	}
	return getSchedule(dateRange)
}

type DateRange struct {
	StartDate string
	EndDate   string
}

// Gets the NHL schedule for today.
func getSchedule(dateRange *DateRange) (*app.Schedule, error) {
	url := "https://statsapi.web.nhl.com/api/v1/schedule"
	if dateRange != nil {
		url = fmt.Sprintf("%s?startDate=%s&endDate=%s", url, dateRange.StartDate, dateRange.EndDate)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	defer resp.Body.Close()

	var schedule app.Schedule
	err = json.NewDecoder(resp.Body).Decode(&schedule)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %w", err)
	}

	return &schedule, nil
}

// Returns yesterday's year, month, and day formatted as YYYY-MM-DD
func getYesterdayYearMonthDay() string {
	today := time.Now()

	yesterday := today.Add(time.Hour * -24)

	year := yesterday.Year()
	month := fmt.Sprintf("%02d", yesterday.Month())
	day := yesterday.Day()

	return fmt.Sprintf("%d-%v-%d", year, month, day)
}
