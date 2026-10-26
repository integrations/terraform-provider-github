package ghclient

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/shurcooL/githubv4"
)

// NewGraphQLClient creates a new GitHub GraphQL client.
func NewGraphQLClient(opts ...ClientOption) (*githubv4.Client, error) {
	o := defaultClientOptions()

	for _, opt := range opts {
		if err := opt(&o); err != nil {
			return nil, fmt.Errorf("failed to apply client option: %w", err)
		}
	}

	if o.userAgent != nil {
		return nil, fmt.Errorf("%w: user agent is not supported for graphql clients", ErrInvalidOption)
	}

	if o.cacheEnabled {
		return nil, fmt.Errorf("%w: caching is not supported for graphql clients", ErrInvalidOption)
	}

	return newGraphQLClient(o)
}

// newGraphQLClient creates a new GitHub GraphQL client using the provided [clientOptions].
func newGraphQLClient(opts clientOptions) (*githubv4.Client, error) {
	tr := newTransport(opts)

	client := &http.Client{Transport: tr, Timeout: opts.clientTimeout}

	if opts.url == nil {
		return githubv4.NewClient(client), nil
	}

	return githubv4.NewEnterpriseClient(*opts.url, client), nil
}

// graphqlURL returns the GraphQL API URL for the given base URL, which defaults to github.com when nil.
func graphqlURL(u *url.URL, isGHES bool) (*url.URL, error) {
	return apiURL(u, isGHES, graphQLAPIPath, ghesGraphQLAPIPath)
}
