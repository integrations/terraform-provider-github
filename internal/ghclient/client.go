package ghclient

import (
	"fmt"
	"time"

	ghct "github.com/bored-engineer/github-conditional-http-transport"
	"golang.org/x/oauth2"

	"github.com/integrations/terraform-provider-github/v6/internal/ghratelimit"
	"github.com/integrations/terraform-provider-github/v6/internal/httpx"
)

// clientOptions defines the configuration options for creating a GitHub API client, including settings for the base URL, user agent, connection pooling, timeouts, concurrency limits, retry behavior, and caching.
type clientOptions struct {
	url                   *string
	userAgent             *string
	maxIdleConns          int
	idleConnTimeout       time.Duration
	responseHeaderTimeout time.Duration
	clientTimeout         time.Duration
	throttleLimiter       httpx.ThrottleLimiter
	rateLimiter           ghratelimit.Limiter
	cacheEnabled          bool
	cacheStore            ghct.Storage
	retryMax              int
	retryWaitMin          time.Duration
	retryWaitMax          time.Duration
	retryJitter           time.Duration
	tokenSource           oauth2.TokenSource
}

// ClientOption defines a function type for configuring client options.
type ClientOption func(*clientOptions) error

// WithClientURL sets the base URL for the GitHub API client. If the provided URL is nil, the default GitHub API URL will be used.
func WithClientURL(u *string) ClientOption {
	return func(o *clientOptions) error {
		o.url = u
		return nil
	}
}

// WithClientUserAgent sets the user agent string for the GitHub API client. If the provided user agent is nil, the default user agent will be used.
func WithClientUserAgent(ua *string) ClientOption {
	return func(o *clientOptions) error {
		o.userAgent = ua

		return nil
	}
}

// WithClientIdleOptions sets the maximum number of idle connections and the idle connection timeout for the GitHub API client. If the provided values are negative, they will be set to 0.
func WithClientIdleOptions(maxIdleConns int, idleConnTimeout time.Duration) ClientOption {
	return func(o *clientOptions) error {
		if maxIdleConns < 0 {
			maxIdleConns = 0
		}

		if idleConnTimeout < 0 {
			idleConnTimeout = 0
		}

		o.maxIdleConns = maxIdleConns
		o.idleConnTimeout = idleConnTimeout

		return nil
	}
}

// WithClientTimeout sets the timeout duration for the GitHub API client. A value of 0 means no timeout. If the provided duration is negative, it will be set to 0.
func WithClientTimeout(timeout time.Duration) ClientOption {
	return func(o *clientOptions) error {
		if timeout < 0 {
			timeout = 0
		}

		o.clientTimeout = timeout

		return nil
	}
}

// WithClientThrottleLimit sets the maximum number of concurrent requests for the GitHub API client. If the provided value is less than or equal to 0, no concurrency limit will be applied.
func WithClientThrottleLimit(throttleLimit int64) ClientOption {
	return func(o *clientOptions) error {
		if throttleLimit <= 0 {
			o.throttleLimiter = nil

			return nil
		}

		o.throttleLimiter = httpx.NewThrottleLimiter(throttleLimit)

		return nil
	}
}

// WithClientCacheOptions sets the caching behavior for the GitHub API client, including whether caching is enabled and the directory path for storing cached responses. If caching is enabled, a cache store will be created at the specified directory path. If an error occurs while creating the cache store, it will be returned.
func WithClientCacheOptions(enabled bool, dirPath string) ClientOption {
	return func(o *clientOptions) error {
		o.cacheEnabled = enabled
		if enabled {
			st, _, err := createCacheStore(dirPath)
			if err != nil {
				return err
			}
			o.cacheStore = st
		}
		return nil
	}
}

// WithClientRetryOptions sets the retry behavior for the GitHub API client, including the maximum number of retries, minimum and maximum wait durations, and jitter. If any of the provided values are negative, they will be set to 0. If the minimum wait is greater than the maximum wait, the maximum wait will be set to the minimum wait.
//
// A zero maxRetries disables retries. The maximum wait also caps how long the client will wait for a rate limit to reset, so a zero waitMax means the client never waits: requests made while rate limited fail immediately and a response with a Retry-After is returned without being retried.
func WithClientRetryOptions(maxRetries int, waitMin, waitMax, jitter time.Duration) ClientOption {
	return func(o *clientOptions) error {
		o.retryMax, o.retryWaitMin, o.retryWaitMax, o.retryJitter = normalizeRetryOptions(maxRetries, waitMin, waitMax, jitter)

		return nil
	}
}

// WithClientTokenSource sets the OAuth2 token source for the GitHub API client. If the provided token source is nil, the client will not be authenticated.
func WithClientTokenSource(ts oauth2.TokenSource) ClientOption {
	return func(o *clientOptions) error {
		if ts == nil {
			return fmt.Errorf("%w: token source cannot be nil when token authentication is enabled", ErrInvalidOption)
		}

		o.tokenSource = ts

		return nil
	}
}

// WithClientToken sets the OAuth2 token for the GitHub API client.
func WithClientToken(token string) ClientOption {
	return func(o *clientOptions) error {
		if token == "" {
			return fmt.Errorf("%w: token cannot be empty when token authentication is enabled", ErrInvalidOption)
		}

		o.tokenSource = oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		return nil
	}
}

// defaultClientOptions returns a new instance of [clientOptions] with default values for connection pooling, timeouts, retry behavior, and caching. These defaults can be overridden by using the provided [ClientOption] functions when creating a new GitHub API client.
func defaultClientOptions() clientOptions {
	return clientOptions{
		maxIdleConns:          maxIdleConnsREST,
		idleConnTimeout:       idleConnTimeoutREST,
		responseHeaderTimeout: responseHeaderTimeout,
		clientTimeout:         clientTimeout,
		retryMax:              httpx.DefaultRetryMaxRetries,
		retryWaitMin:          httpx.DefaultRetryWaitMin,
		retryWaitMax:          httpx.DefaultRetryWaitMax,
		retryJitter:           httpx.DefaultRetryJitter,
	}
}

// normalizeRetryOptions clamps the retry values so none are negative and the maximum wait is not less than the minimum wait.
func normalizeRetryOptions(maxRetries int, waitMin, waitMax, jitter time.Duration) (int, time.Duration, time.Duration, time.Duration) {
	waitMin = max(waitMin, 0)

	return max(maxRetries, 0), waitMin, max(waitMax, waitMin), max(jitter, 0)
}
