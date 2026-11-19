package ghratelimit

import (
	"errors"
	"fmt"
	"time"
)

// ErrRateLimitWaitExceeded represents an error indicating that the maximum wait time for a rate limit has been exceeded.
var ErrRateLimitWaitExceeded = errors.New("rate limit wait exceeded")

// WaitExceededError is returned when a request is blocked by a rate limit for longer than the maximum wait. It wraps [ErrRateLimitWaitExceeded].
type WaitExceededError struct {
	// Category is the rate limit category that is blocking the request.
	Category string
	// Until is the time the rate limit is blocked until.
	Until time.Time
	// MaxWait is the maximum time the request was allowed to wait.
	MaxWait time.Duration
}

// Error implements the error interface for [WaitExceededError].
func (e *WaitExceededError) Error() string {
	return fmt.Sprintf("%s: %s rate limit is blocked until %s, which is longer than the maximum wait of %s", ErrRateLimitWaitExceeded, e.Category, e.Until.UTC().Format(time.RFC3339), e.MaxWait.Round(time.Second))
}

// Unwrap returns [ErrRateLimitWaitExceeded] so the error can be matched with [errors.Is].
func (e *WaitExceededError) Unwrap() error {
	return ErrRateLimitWaitExceeded
}
