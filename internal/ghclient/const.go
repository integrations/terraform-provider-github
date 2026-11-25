package ghclient

import (
	"net/url"
	"time"
)

const (
	// DotComAPIURLStr is the base API URL for github.com.
	DotComAPIURLStr = "https://api.github.com/"
)

// dotComAPIURL is [DotComAPIURLStr] as a [url.URL]. It's never modified, as [url.URL.JoinPath] returns a copy.
var dotComAPIURL = url.URL{Scheme: "https", Host: "api.github.com", Path: "/"}

const (
	// restAPIPath is the rest api path for api.github.com & ghe.com.
	restAPIPath = "/"
	// graphQLAPIPath is the graphql api path for api.github.com & ghe.com.
	graphQLAPIPath = "/graphql"
	// ghesRESTAPIPath is the rest api path for GitHub Enterprise Server.
	ghesRESTAPIPath = "api/v3/"
	// ghesGraphQLAPIPath is the GraphQL api path for GitHub Enterprise Server.
	ghesGraphQLAPIPath = "api/graphql"
)

const (
	// maxIdleConnsREST defines the maximum number of idle connections for REST API requests.
	maxIdleConnsREST int = 100

	// maxIdleConnsGraphQL defines the maximum number of idle connections for GraphQL API requests.
	maxIdleConnsGraphQL int = 10

	// idleConnTimeoutREST defines the timeout duration for idle REST API connections.
	idleConnTimeoutREST = 90 * time.Second

	// idleConnTimeoutGraphQL defines the timeout duration for idle GraphQL API connections.
	idleConnTimeoutGraphQL = 120 * time.Second

	// responseHeaderTimeout defines how long to wait for a response's headers once the request has been written. GitHub stops processing a request after about 10 seconds, so this only trips on connections that have stalled.
	responseHeaderTimeout = 60 * time.Second

	// clientTimeout defines the timeout duration for GitHub API requests, this is set to be below the default 20 minutes Terraform context timeout but still long enough for secondary ratelimit retrys.
	clientTimeout = 10 * time.Minute
)
