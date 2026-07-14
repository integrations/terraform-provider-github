package graphql

import (
	"fmt"
	"time"
)

func ParseDate(value, subject string) (time.Time, error) {
	date, err := time.Parse(time.DateOnly, value)
	if err != nil {
		return time.Time{}, fmt.Errorf("parsing %s date: %w", subject, err)
	}
	return date, nil
}
