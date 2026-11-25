package ghclient

import (
	"context"
	"crypto"
	"fmt"
	"net/url"
	"path/filepath"
	"time"

	ghct "github.com/bored-engineer/github-conditional-http-transport"
	"github.com/google/go-github/v92/github"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"

	"github.com/integrations/terraform-provider-github/v6/internal/ghratelimit"
	"github.com/integrations/terraform-provider-github/v6/internal/httpx"
)

// sourceOptions defines the configuration shared by every client created by a [Source], including the base URL, user agent, throttling, caching, retry behavior, and authentication.
type sourceOptions struct {
	baseURL       *url.URL
	isGHES        bool
	userAgent     *string
	throttleLimit int64
	cacheEnabled  bool
	cacheBasePath string
	retryMax      int
	retryWaitMin  time.Duration
	retryWaitMax  time.Duration
	retryJitter   time.Duration
	appAuth       bool
	appClientID   string
	appSigner     crypto.Signer
	tokenSource   oauth2.TokenSource
}

// SourceOption defines a function type for configuring source options.
type SourceOption func(*sourceOptions) error

// WithSourceBaseURL sets the base URL for the source's clients and whether it is a GitHub Enterprise Server URL. If the provided URL is nil, the default GitHub API URL will be used; a nil URL is an error when isGHES is true.
func WithSourceBaseURL(baseURL *url.URL, isGHES bool) SourceOption {
	return func(o *sourceOptions) error {
		if isGHES && baseURL == nil {
			return errGHESBaseURLRequired
		}

		o.baseURL = baseURL
		o.isGHES = isGHES

		return nil
	}
}

// WithSourceUserAgent sets the user agent string for the source's REST clients. If the provided user agent is nil, the default user agent will be used. GraphQL clients never set a user agent.
func WithSourceUserAgent(userAgent *string) SourceOption {
	return func(o *sourceOptions) error {
		o.userAgent = userAgent

		return nil
	}
}

// WithSourceThrottleLimit sets the maximum number of concurrent requests for each client group, which is shared by the REST and GraphQL clients of the group. If the provided value is less than or equal to 0, no concurrency limit will be applied.
func WithSourceThrottleLimit(throttleLimit int64) SourceOption {
	return func(o *sourceOptions) error {
		o.throttleLimit = throttleLimit

		return nil
	}
}

// WithSourceCache enables caching of REST responses, storing each client group's cache in a directory under the provided base path. The base path cannot be empty.
func WithSourceCache(basePath string) SourceOption {
	return func(o *sourceOptions) error {
		if basePath == "" {
			return fmt.Errorf("%w: cache base path cannot be empty", ErrInvalidOption)
		}

		o.cacheEnabled = true
		o.cacheBasePath = basePath

		return nil
	}
}

// WithSourceRetry sets the retry behavior for the source's clients, including the maximum number of retries, minimum and maximum wait durations, and jitter. If any of the provided values are negative, they will be set to 0. If the minimum wait is greater than the maximum wait, the maximum wait will be set to the minimum wait.
//
// A zero maxRetries disables retries. The maximum wait also caps how long a client will wait for a rate limit to reset, so a zero waitMax means the client never waits: requests made while rate limited fail immediately and a response with a Retry-After is returned without being retried.
func WithSourceRetry(maxRetries int, waitMin, waitMax, jitter time.Duration) SourceOption {
	return func(o *sourceOptions) error {
		o.retryMax, o.retryWaitMin, o.retryWaitMax, o.retryJitter = normalizeRetryOptions(maxRetries, waitMin, waitMax, jitter)

		return nil
	}
}

// WithSourceAppAuth enables GitHub App authentication using the provided client ID and signer. The client ID and signer cannot be empty. This replaces any token or token source set previously.
func WithSourceAppAuth(clientID string, signer crypto.Signer) SourceOption {
	return func(o *sourceOptions) error {
		if clientID == "" {
			return fmt.Errorf("%w: client id cannot be empty when app authentication is enabled", ErrInvalidOption)
		}

		if signer == nil {
			return fmt.Errorf("%w: signer cannot be nil when app authentication is enabled", ErrInvalidOption)
		}

		o.appAuth = true
		o.appClientID = clientID
		o.appSigner = signer

		return nil
	}
}

// WithSourceTokenSource enables token authentication using the provided token source, which cannot be nil. This replaces any app authentication set previously.
func WithSourceTokenSource(ts oauth2.TokenSource) SourceOption {
	return func(o *sourceOptions) error {
		if ts == nil {
			return fmt.Errorf("%w: token source cannot be nil when token authentication is enabled", ErrInvalidOption)
		}

		o.appAuth = false
		o.tokenSource = ts

		return nil
	}
}

// WithSourceToken enables token authentication using the provided static token, which cannot be empty. This replaces any app authentication set previously.
func WithSourceToken(token string) SourceOption {
	return func(o *sourceOptions) error {
		if token == "" {
			return fmt.Errorf("%w: token cannot be empty when token authentication is enabled", ErrInvalidOption)
		}

		o.appAuth = false
		o.tokenSource = oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})

		return nil
	}
}

// defaultSourceOptions returns a new instance of [sourceOptions] with default values for retry behavior. These defaults can be overridden by using the provided [SourceOption] functions when creating a new [Source].
func defaultSourceOptions() sourceOptions {
	return sourceOptions{
		retryMax:     httpx.DefaultRetryMaxRetries,
		retryWaitMin: httpx.DefaultRetryWaitMin,
		retryWaitMax: httpx.DefaultRetryWaitMax,
		retryJitter:  httpx.DefaultRetryJitter,
	}
}

// sourceClientGroup holds the REST and GraphQL clients that share authentication, throttling, rate limit state and cache store.
type sourceClientGroup struct {
	restClient    *github.Client
	graphqlClient *githubv4.Client
	cleanup       func() error
}

// newSourceClientGroup creates the REST and GraphQL clients for the given token source. The REST and GraphQL clients share a throttle limiter and a rate limiter, as GitHub scopes both to the authenticated identity, and the REST client uses a cache store in the cacheDir under the cache base path when caching is enabled.
func newSourceClientGroup(opts sourceOptions, cacheDir string, ts oauth2.TokenSource) (*sourceClientGroup, error) {
	var cacheStore ghct.Storage
	var closeStore func() error
	if opts.cacheEnabled {
		st, cl, err := createCacheStore(filepath.Join(opts.cacheBasePath, cacheDir))
		if err != nil {
			return nil, err
		}
		cacheStore = st
		closeStore = cl
	}

	var throttleLimiter httpx.ThrottleLimiter
	if opts.throttleLimit > 0 {
		throttleLimiter = httpx.NewThrottleLimiter(opts.throttleLimit)
	}
	rateLimiter := ghratelimit.NewLimiter()

	ru, err := restURL(opts.baseURL, opts.isGHES)
	if err != nil {
		return nil, err
	}

	rco := clientOptions{
		url:                   new(ru.String()),
		userAgent:             opts.userAgent,
		maxIdleConns:          maxIdleConnsREST,
		idleConnTimeout:       idleConnTimeoutREST,
		responseHeaderTimeout: responseHeaderTimeout,
		clientTimeout:         clientTimeout,
		throttleLimiter:       throttleLimiter,
		rateLimiter:           rateLimiter,
		cacheEnabled:          opts.cacheEnabled,
		cacheStore:            cacheStore,
		retryMax:              opts.retryMax,
		retryWaitMin:          opts.retryWaitMin,
		retryWaitMax:          opts.retryWaitMax,
		retryJitter:           opts.retryJitter,
		tokenSource:           ts,
	}

	rc, err := newRESTClient(rco)
	if err != nil {
		return nil, err
	}

	gu, err := graphqlURL(opts.baseURL, opts.isGHES)
	if err != nil {
		return nil, err
	}

	gco := clientOptions{
		url:                   new(gu.String()),
		maxIdleConns:          maxIdleConnsGraphQL,
		idleConnTimeout:       idleConnTimeoutGraphQL,
		responseHeaderTimeout: responseHeaderTimeout,
		clientTimeout:         clientTimeout,
		throttleLimiter:       throttleLimiter,
		rateLimiter:           rateLimiter,
		retryMax:              opts.retryMax,
		retryWaitMin:          opts.retryWaitMin,
		retryWaitMax:          opts.retryWaitMax,
		retryJitter:           opts.retryJitter,
		tokenSource:           ts,
	}

	gc, err := newGraphQLClient(gco)
	if err != nil {
		return nil, err
	}

	cleanup := func() error {
		if closeStore != nil {
			if err := closeStore(); err != nil {
				return fmt.Errorf("failed to close cache store: %w", err)
			}
		}
		return nil
	}

	return &sourceClientGroup{
		restClient:    rc,
		graphqlClient: gc,
		cleanup:       cleanup,
	}, nil
}

// NewSource creates a new [Source] with the provided options. If app authentication is configured the source authenticates as the app and its installations, otherwise it uses the configured token or token source, or is unauthenticated if neither is set.
func NewSource(opts ...SourceOption) (Source, error) {
	o := defaultSourceOptions()

	for _, opt := range opts {
		if err := opt(&o); err != nil {
			return nil, fmt.Errorf("failed to apply source option: %w", err)
		}
	}

	if o.appAuth {
		s, err := newAppSource(o)
		if err != nil {
			return nil, err
		}

		return s, nil
	}

	s, err := newTokenSource(o)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// Source provides an interface for obtaining GitHub clients. It abstracts away the details of how the clients are created and authenticated, so an app-based source and a token-based (or unauthenticated) source can be used interchangeably.
type Source interface {
	// RESTClient returns the default GitHub client for the source.
	RESTClient(ctx context.Context) (*github.Client, error)

	// OwnerRESTClient returns a GitHub client for accessing resources owned by the specified owner (which can be either a user or an organization). An app source returns a client authenticated as the owner's installation, while a token source returns its default client.
	OwnerRESTClient(ctx context.Context, owner string) (*github.Client, error)

	// GraphQLClient returns the default GitHub GraphQL client for the source.
	GraphQLClient(ctx context.Context) (*githubv4.Client, error)

	// OwnerGraphQLClient returns a GitHub GraphQL client for accessing resources owned by the specified owner (which can be either a user or an organization). An app source returns a client authenticated as the owner's installation, while a token source returns its default client.
	OwnerGraphQLClient(ctx context.Context, owner string) (*githubv4.Client, error)
}
