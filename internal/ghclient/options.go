package ghclient

import (
	"fmt"
	"net/url"
	"path/filepath"
	"time"

	"golang.org/x/sync/semaphore"
)

// SourceOptions defines the configuration options for creating a GitHub client source.
type SourceOptions struct {
	BaseURL       string
	IsGHES        bool
	UserAgent     string
	Cache         bool
	CacheBasePath string
	RetryMax      int
	RetryWaitMin  time.Duration
	RetryWaitMax  time.Duration
}

// getRESTClientOptions returns the REST client options derived from the source options.
func (o *SourceOptions) getRESTClientOptions(sema *semaphore.Weighted, cacheRef string) ClientOptions {
	return ClientOptions{
		BaseURL:         o.BaseURL,
		IsGHES:          o.IsGHES,
		UserAgent:       o.UserAgent,
		Cache:           o.Cache,
		CachePath:       filepath.Join(o.CacheBasePath, cacheRef),
		RetryMax:        o.RetryMax,
		RetryWaitMin:    o.RetryWaitMin,
		RetryWaitMax:    o.RetryWaitMax,
		Sema:            sema,
		MaxIdleConns:    maxIdleConnsREST,
		IdleConnTimeout: idleConnTimeoutREST,
	}
}

// getGraphQLClientOptions returns the GraphQL client options derived from the source options.
func (o *SourceOptions) getGraphQLClientOptions(sema *semaphore.Weighted) ClientOptions {
	return ClientOptions{
		BaseURL:         o.BaseURL,
		IsGHES:          o.IsGHES,
		UserAgent:       o.UserAgent,
		RetryMax:        o.RetryMax,
		RetryWaitMin:    o.RetryWaitMin,
		RetryWaitMax:    o.RetryWaitMax,
		Sema:            sema,
		MaxIdleConns:    maxIdleConnsGraphQL,
		IdleConnTimeout: idleConnTimeoutGraphQL,
	}
}

// ClientOptions defines the configuration options for creating a GitHub client.
type ClientOptions struct {
	BaseURL         string
	IsGHES          bool
	UserAgent       string
	Cache           bool
	CachePath       string
	RetryMax        int
	RetryWaitMin    time.Duration
	RetryWaitMax    time.Duration
	Sema            *semaphore.Weighted
	MaxIdleConns    int
	IdleConnTimeout time.Duration
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
