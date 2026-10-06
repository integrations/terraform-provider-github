package ghclient

import (
	"fmt"
	"net/url"
	"time"

	"golang.org/x/sync/semaphore"
)

// SourceOptions defines the configuration options for creating a GitHub client source.
type SourceOptions struct {
	BaseURL   string
	IsGHES    bool
	UserAgent string
	Cache     CacheOptions
	Retry     RetryOptions
}

// CacheOptions defines the configuration options for caching.
type CacheOptions struct {
	Enabled  bool
	BasePath string
	Ref      string
}

// RetryOptions defines the configuration options for retrying failed requests.
type RetryOptions struct {
	Max     int
	WaitMin time.Duration
	WaitMax time.Duration
	Jitter  time.Duration
}

// ConcurrencyOptions defines the configuration options for concurrency.
type ConcurrencyOptions struct {
	Max  int64
	sema *semaphore.Weighted
}

// getRESTClientOptions returns the REST client options derived from the source options.
func (o *SourceOptions) getRESTClientOptions(sema *semaphore.Weighted, cacheRef string) ClientOptions {
	return ClientOptions{
		BaseURL:         o.BaseURL,
		IsGHES:          o.IsGHES,
		UserAgent:       o.UserAgent,
		MaxIdleConns:    maxIdleConnsREST,
		IdleConnTimeout: idleConnTimeoutREST,
		Concurrency: ConcurrencyOptions{
			Max:  maxConcurrentRequests,
			sema: sema,
		},
		Retry: o.Retry,
		Cache: CacheOptions{
			Enabled:  o.Cache.Enabled,
			BasePath: o.Cache.BasePath,
			Ref:      cacheRef,
		},
	}
}

// getGraphQLClientOptions returns the GraphQL client options derived from the source options.
func (o *SourceOptions) getGraphQLClientOptions(sema *semaphore.Weighted) ClientOptions {
	return ClientOptions{
		BaseURL:         o.BaseURL,
		IsGHES:          o.IsGHES,
		UserAgent:       o.UserAgent,
		MaxIdleConns:    maxIdleConnsGraphQL,
		IdleConnTimeout: idleConnTimeoutGraphQL,
		Concurrency: ConcurrencyOptions{
			Max:  maxConcurrentRequests,
			sema: sema,
		},
		Retry: o.Retry,
	}
}

// ClientOptions defines the configuration options for creating a GitHub client.
type ClientOptions struct {
	BaseURL         string
	IsGHES          bool
	UserAgent       string
	MaxIdleConns    int
	IdleConnTimeout time.Duration
	Concurrency     ConcurrencyOptions
	Cache           CacheOptions
	Retry           RetryOptions
}

// getRESTURL returns the REST API URL based on the provided base URL and whether it is a GitHub Enterprise Server instance.
func (o *ClientOptions) getRESTURL() (*string, error) {
	baseURL := o.BaseURL
	if baseURL == "" {
		baseURL = DotComAPIURL
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse base url: %w", err)
	}

	if o.IsGHES {
		u = u.JoinPath(GHESRESTAPIPath)
	} else {
		u = u.JoinPath(RESTAPIPath)
	}

	return new(u.String()), nil
}

// getGraphQLURL returns the GraphQL API URL based on the provided base URL and whether it is a GitHub Enterprise Server instance.
func (o *ClientOptions) getGraphQLURL() (*string, error) {
	baseURL := o.BaseURL
	if baseURL == "" {
		baseURL = DotComAPIURL
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse base url: %w", err)
	}

	if o.IsGHES {
		u = u.JoinPath(GHESGraphQLAPIPath)
	} else {
		u = u.JoinPath(GraphQLAPIPath)
	}

	return new(u.String()), nil
}
