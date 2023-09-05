package app

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// DateRange represents an inclusive range of current or past dates between StartDate and EndDate.
//
// StartDate and EndDate are expected to occur in chronological order and are allowed to overlap.
type DateRange struct {
	StartDate string
	EndDate   string
}

var (
	ErrorNoDate           = errors.New("date is empty")
	ErrorInvalidFormat    = errors.New("date is not formatted as YYYY-MM-DD")
	ErrorFutureDate       = errors.New("date occurs in the future")
	ErrorNotChronological = errors.New("start date occurs before the end date")
)

func (d DateRange) Parse() error {
	d.StartDate = strings.TrimSpace(d.StartDate)
	d.EndDate = strings.TrimSpace(d.EndDate)

	// Check if strings are non-empty
	if d.StartDate == "" || d.EndDate == "" {
		return fmt.Errorf("[DateRange.Parse] error: %w", ErrorNoDate)
	}

	// Check if dates are YYYY-MM-DD formatted
	startDate, startDateErr := time.Parse(time.DateOnly, d.StartDate)
	endDate, endDateErr := time.Parse(time.DateOnly, d.EndDate)
	if startDateErr != nil || endDateErr != nil {
		return fmt.Errorf("[DateRange.Parse] error: %w", ErrorInvalidFormat)
	}

	// Check if start and end dates are not in the future
	now := time.Now()
	if startDate.After(now) || endDate.After(now) {
		return fmt.Errorf("[DateRange.Parse] error: %w", ErrorFutureDate)
	}

	// Check if start and end dates are chronological
	if startDate.After(endDate) {
		return fmt.Errorf("[DateRange.Parse] error: %w", ErrorNotChronological)
	}

	return nil
}
