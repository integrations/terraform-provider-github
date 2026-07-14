package github

import (
	"fmt"
	"time"
)

func validateProjectV2Date(value any, key string) ([]string, []error) {
	date, ok := value.(string)
	if !ok {
		return nil, []error{fmt.Errorf("%q must be a string, got %T", key, value)}
	}
	if _, err := time.Parse(time.DateOnly, date); err != nil {
		return nil, []error{fmt.Errorf("%q must use YYYY-MM-DD format: %w", key, err)}
	}
	return nil, nil
}
