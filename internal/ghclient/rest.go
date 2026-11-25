package ghclient

import (
	"fmt"
	"net/url"

	"github.com/google/go-github/v92/github"
)

// NewRESTClient creates a new GitHub REST client with the provided options. If a token source is provided, the client will be authenticated; otherwise, it will be unauthenticated.
func NewRESTClient(opts ...ClientOption) (*github.Client, error) {
	o := defaultClientOptions()

	for _, opt := range opts {
		if err := opt(&o); err != nil {
			return nil, fmt.Errorf("failed to apply client option: %w", err)
		}
	}

	return newRESTClient(o)
}

// newRESTClient creates a new GitHub REST client using the provided [clientOptions].
func newRESTClient(opts clientOptions) (*github.Client, error) {
	tr := newTransport(opts)

	clientOpts := []github.ClientOptionsFunc{
		github.WithTransport(tr),
		github.WithTimeout(opts.clientTimeout),
		github.WithDisableRateLimitCheck(),
	}

	if opts.userAgent != nil {
		clientOpts = append(clientOpts, github.WithUserAgent(*opts.userAgent))
	}

	if opts.url != nil {
		clientOpts = append(clientOpts, github.WithURLs(opts.url, nil))
	}

	return github.NewClient(clientOpts...)
}

// restURL returns the REST API URL for the given base URL, which defaults to github.com when nil.
func restURL(u *url.URL, isGHES bool) (*url.URL, error) {
	return apiURL(u, isGHES, restAPIPath, ghesRESTAPIPath)
}

// apiURL returns the API URL for the given base URL by joining the path, or the GHES path when isGHES is true. The base URL defaults to github.com when nil, which is an error for GHES.
func apiURL(u *url.URL, isGHES bool, path, ghesPath string) (*url.URL, error) {
	if isGHES && u == nil {
		return nil, errGHESBaseURLRequired
	}

	if u == nil {
		u = &dotComAPIURL
	}

	if isGHES {
		return u.JoinPath(ghesPath), nil
	}

	return u.JoinPath(path), nil
}
