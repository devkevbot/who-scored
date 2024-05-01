package app

import (
	"errors"
	"time"
)

func ParseUserProvidedDate(date string) error {
	parsedTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		return errors.New("expected a valid date in YYYY-MM-DD format")
	}

	currentTime := time.Now()
	if parsedTime.After(currentTime) {
		return errors.New("expected a date in the past or present")
	}

	dateOfFirstNHLGame := time.Date(1917, 12, 19, 0, 0, 0, 0, time.UTC)
	if parsedTime.Before(dateOfFirstNHLGame) {
		return errors.New("expected the date to be on or after 1917-12-19, which was the first NHL game")
	}
	return nil
}
