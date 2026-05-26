package ghclient

import "time"

// Options defines the configuration options for creating GitHub clients.
type Options struct {
	RESTAPIURL    *string
	RESTUploadURL *string
	GraphQLURL    *string
	CachePath     *string
	CacheRef      *string
	RetryMax      int
	RetryWaitMin  time.Duration
	RetryWaitMax  time.Duration
}
