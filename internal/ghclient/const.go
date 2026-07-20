package ghclient

import "time"

const (
	// DotComAPIURL is the base API URL for github.com.
	DotComAPIURL = "https://api.github.com/"
)

const (
	// RESTAPIPath is the rest api path for api.github.com & ghe.com.
	RESTAPIPath = "/"
	// GraphQLAPIPath is the graphql api path for api.github.com & ghe.com.
	GraphQLAPIPath = "/graphql"
	// GHESRESTAPISuffix is the rest api suffix for GitHub Enterprise Server.
	GHESRESTAPIPath = "api/v3/"
	// GHESGraphQLAPISuffix is the GraphQL api suffix for GitHub Enterprise Server.
	GHESGraphQLAPIPath = "api/graphql"
)

const (
	// githubSecondaryRateLimitMaxConcurrency defines the maximum number of concurrent requests allowed to the GitHub API under secondary rate limits.
	githubSecondaryRateLimitMaxConcurrency int64 = 100

	// maxConcurrentRequests defines the maximum number of concurrent requests allowed to the GitHub API.
	maxConcurrentRequests int64 = githubSecondaryRateLimitMaxConcurrency

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
