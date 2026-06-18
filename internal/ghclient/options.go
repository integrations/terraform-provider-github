package ghclient

import (
	"time"

	"golang.org/x/sync/semaphore"
)

// Options defines the configuration options for creating GitHub clients.
type Options struct {
	RESTAPIURL   string
	GraphQLURL   string
	UserAgent    string
	CachePath    string
	RetryMax     int
	RetryWaitMin time.Duration
	RetryWaitMax time.Duration

	// Set per usage.
	sema            *semaphore.Weighted
	maxIdleConns    int
	idleConnTimeout time.Duration
	cacheRef        string
}
