package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/devkevbot/who-scored/internal/app"
	"github.com/devkevbot/who-scored/internal/pkg/utils"
)

func GetScheduleForToday() (*app.Schedule, error) {
	return getSchedule(nil)
}

func GetScheduleForYesterday() (*app.Schedule, error) {
	yesterday := utils.GetYesterdayYearMonthDay()
	dateRange := &DateRange{
		StartDate: yesterday,
		EndDate:   yesterday,
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
