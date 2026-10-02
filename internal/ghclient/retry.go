package ghclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/sethvargo/go-retry"
)

// newRetryTransport returns a [retryTransport] that wraps the given [http.RoundTripper] using the [RetryOptions] to configure the retry backoff strategy.
func newRetryTransport(inner http.RoundTripper, opts RetryOptions) (http.RoundTripper, error) {
	if opts.Max <= 0 {
		return inner, nil
	}

	if opts.WaitMin <= 0 {
		return nil, fmt.Errorf("wait min must be greater than 0")
	}

	return &retryTransport{
		inner: inner,
		opts:  opts,
	}, nil
}

// retryTransport is a [http.RoundTripper] that wraps another http.RoundTripper and retries failed requests using an exponential backoff strategy.
type retryTransport struct {
	inner http.RoundTripper
	opts  RetryOptions
}

// RoundTrip implements the [http.RoundTripper] interface for the [retryTransport]. It retries failed requests using the configured backoff strategy.
func (t *retryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var prevResp *http.Response

	backoff, err := newBackoff(t.opts)
	if err != nil {
		return nil, err
	}

	resp, err := retry.DoValue(req.Context(), backoff, func(ctx context.Context) (*http.Response, error) {
		// If there was a previous response, drain it and reset the body
		if prevResp != nil {
			_ = drainResponseBody(prevResp)
			prevResp = nil
		}

		resp, err := t.inner.RoundTrip(req)
		if err != nil {
			_ = drainResponseBody(resp)
			if isRetryableNetworkError(err) {
				return nil, retry.RetryableError(err)
			}
			return resp, err
		}

		if isRetryableStatusCode(resp.StatusCode) {
			// Keep the response in case we can't retry and need to return it to the caller
			prevResp = resp
			return nil, retry.RetryableError(fmt.Errorf("server error: %d", resp.StatusCode))
		}

		return resp, err
	})

	return resp, err
}

// newBackoff returns a retry.Backoff configured according to the given RetryOptions.
func newBackoff(opts RetryOptions) (bo retry.Backoff, err error) {
	// Required as package panics on error
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("invalid backoff configuration: %v", r)
		}
	}()

	b := retry.NewExponential(opts.WaitMin)
	b = retry.WithCappedDuration(opts.WaitMax, b)
	b = retry.WithMaxRetries(uint64(opts.Max), b)
	b = retry.WithJitter(opts.Jitter, b)

	return b, nil
}

// isRateLimitStatusCode returns true if the given status code is a rate limit status code.
func isRateLimitStatusCode(code int) bool {
	return code == http.StatusTooManyRequests || code == http.StatusForbidden
}

// isRetryableStatusCode returns true if the given status code is retryable.
func isRetryableStatusCode(code int) bool {
	return code >= http.StatusInternalServerError && code != http.StatusNotImplemented
}

// isRetryableNetworkError returns true if the given error is retryable.
func isRetryableNetworkError(err error) bool {
	// Do not retry on context cancellation/deadline
	if err == nil || errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return false
	}

	// Return true for other network errors
	return true
}
