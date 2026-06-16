package ghclient

import "time"

const (
	// maxConcurrentRequests defines the maximum number of concurrent requests allowed to the GitHub API.
	maxConcurrentRequests int64 = 100

	// maxIdleConnsREST defines the maximum number of idle connections for REST API requests.
	maxIdleConnsREST int = 100

	// maxIdleConnsGraphQL defines the maximum number of idle connections for GraphQL API requests.
	maxIdleConnsGraphQL int = 10

	// idleConnTimeoutREST defines the timeout duration for idle REST API connections.
	idleConnTimeoutREST = 90 * time.Second

	// idleConnTimeoutGraphQL defines the timeout duration for idle GraphQL API connections.
	idleConnTimeoutGraphQL = 120 * time.Second

	// clientTimeout defines the timeout duration for GitHub API requests.
	clientTimeout = 5 * time.Minute
)
