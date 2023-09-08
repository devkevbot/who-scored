package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/devkevbot/who-scored/internal/app"
)

// GetScheduleForToday retrieves the schedules for today's NHL games.
func GetScheduleForToday() (*app.Schedule, error) {
	schedule, err := getSchedule(nil)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch NHL schedule for today: %w", errors.Unwrap(err))
	}
	return schedule, nil
}

// GetScheduleForYesterday retrieves the schedules for yesterday's NHL games.
func GetScheduleForYesterday() (*app.Schedule, error) {
	yesterday := getYesterdayYearMonthDay()
	dateRange := &app.DateRange{
		StartDate: yesterday,
		EndDate:   yesterday,
	}
	if err := dateRange.Parse(); err != nil {
		return nil, fmt.Errorf("unable to fetch NHL schedule for yesterday: %w", errors.Unwrap(err))
	}

	schedule, err := getSchedule(dateRange)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch NHL schedule for yesterday: %w", errors.Unwrap(err))
	}
	return schedule, nil
}

// GetScheduleForSingleDay retrieves the schedules for NHL games occurring on startDate.
func GetScheduleForSingleDay(startDate string) (*app.Schedule, error) {
	dateRange := &app.DateRange{
		StartDate: startDate,
		EndDate:   startDate,
	}
	if err := dateRange.Parse(); err != nil {
		return nil, fmt.Errorf("unable to fetch NHL schedule for date: %w", errors.Unwrap(err))
	}

	schedule, err := getSchedule(dateRange)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch NHL schedule for date: %w", errors.Unwrap(err))
	}
	return schedule, nil
}

// GetScheduleForDateRange retrieves the schedules for NHL games occurring during the period formed by startDate and endDate.
func GetScheduleForDateRange(startDate, endDate string) (*app.Schedule, error) {
	dateRange := &app.DateRange{
		StartDate: startDate,
		EndDate:   endDate,
	}
	if err := dateRange.Parse(); err != nil {
		return nil, fmt.Errorf("unable to fetch NHL schedule for date range: %w", errors.Unwrap(err))
	}

	schedule, err := getSchedule(dateRange)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch NHL schedule for date range: %w", errors.Unwrap(err))
	}
	return schedule, nil
}

// Gets the NHL schedule for the date range.
func getSchedule(dateRange *app.DateRange) (*app.Schedule, error) {
	req, err := http.NewRequest("GET", nhlScheduleApiPath(dateRange), nil)
	if err != nil {
		return nil, fmt.Errorf("[getSchedule] error creating request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("[getSchedule] error making request: %w", err)
	}

	var schedule app.Schedule
	err = json.NewDecoder(resp.Body).Decode(&schedule)
	if err != nil {
		return nil, fmt.Errorf("[getSchedule] error decoding JSON response: %w", err)
	}
	return &schedule, nil
}

func nhlScheduleApiPath(dateRange *app.DateRange) string {
	base := "https://statsapi.web.nhl.com/api/v1/schedule"
	if dateRange == nil {
		return base
	}
	return fmt.Sprintf("%s?startDate=%s&endDate=%s", base, dateRange.StartDate, dateRange.EndDate)
}

// Returns yesterday's year, month, and day formatted as YYYY-MM-DD
func getYesterdayYearMonthDay() string {
	today := time.Now()

	yesterday := today.Add(time.Hour * -24)

	year := yesterday.Year()
	month := fmt.Sprintf("%02d", yesterday.Month())
	day := fmt.Sprintf("%02d", yesterday.Day())

	return fmt.Sprintf("%d-%v-%v", year, month, day)
}
