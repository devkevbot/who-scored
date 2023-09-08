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
	ErrStartDateEmpty = errors.New("start date is empty")
	ErrEndDateEmpty   = errors.New("end date is empty")

	// "Out of range" errors are different from "malformed" errors in the sense that
	// the date is structurally correct but is otherwise invalid. Example: February 29th
	// only exists on leap years.

	ErrStartDateOutOfRange = errors.New("start date is not a valid date")
	ErrEndDateOutOfRange   = errors.New("end date is not a valid date")

	ErrStartDateMalformed = errors.New("start date is not formatted as YYYY-MM-DD")
	ErrEndDateMalformed   = errors.New("end date is not formatted as YYYY-MM-DD")

	ErrStartDateOccursInFuture = errors.New("start date occurs in the future")
	ErrEndDateOccursInFuture   = errors.New("end date occurs in the future")

	ErrNotChronological = errors.New("start date occurs before the end date")
)

func (d DateRange) Parse() error {
	d.StartDate = strings.TrimSpace(d.StartDate)
	d.EndDate = strings.TrimSpace(d.EndDate)

	// Check if start date is non-empty
	if d.StartDate == "" {
		return fmt.Errorf("[DateRange.Parse] error: %w", ErrStartDateEmpty)
	}

	// Check if start date is YYYY-MM-DD formatted
	startDate, startDateErr := time.Parse(time.DateOnly, d.StartDate)
	if startDateErr != nil {
		if strings.Contains(startDateErr.Error(), "out of range") {
			return fmt.Errorf("[DateRange.Parse] error: %w", ErrStartDateOutOfRange)
		}

		return fmt.Errorf("[DateRange.Parse] error: %w", ErrStartDateMalformed)
	}

	// Check if start date is not in the future
	now := time.Now()
	if startDate.After(now) {
		return fmt.Errorf("[DateRange.Parse] error: %w", ErrStartDateOccursInFuture)
	}

	// If the start date and end date are different, we must check the end date separately.
	if d.StartDate != d.EndDate {
		// Check if end date is non-empty
		if d.EndDate == "" {
			return fmt.Errorf("[DateRange.Parse] error: %w", ErrEndDateEmpty)
		}

		// Check if end date is YYYY-MM-DD formatted
		endDate, endDateErr := time.Parse(time.DateOnly, d.EndDate)
		if endDateErr != nil {
			if strings.Contains(endDateErr.Error(), "out of range") {
				return fmt.Errorf("[DateRange.Parse] error: %w", ErrEndDateOutOfRange)
			}

			return fmt.Errorf("[DateRange.Parse] error: %w", ErrEndDateMalformed)
		}

		// Check if end date is not in the future
		if endDate.After(now) {
			return fmt.Errorf("[DateRange.Parse] error: %w", ErrEndDateOccursInFuture)
		}

		// Check if start and end dates are chronological
		if startDate.After(endDate) {
			return fmt.Errorf("[DateRange.Parse] error: %w", ErrNotChronological)
		}
	}

	return nil
}
