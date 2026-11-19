package httpx

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"slices"
	"time"

	"github.com/sethvargo/go-retry"
)

// retryOptions defines the configuration for retrying failed requests.
type retryOptions struct {
	maxRetries    int
	waitMin       time.Duration
	waitMax       time.Duration
	jitter        time.Duration
	noRetryErrors []error
}

// RetryOption defines a function type for configuring retry options.
type RetryOption func(*retryOptions)

// WithRetryMaxRetries sets the maximum number of retries for failed requests. If the provided value is less than 0, it will be set to 0. If the value is 0, no retries will be performed.
func WithRetryMaxRetries(m int) RetryOption {
	return func(o *retryOptions) {
		if m < 0 {
			m = 0
		}

		o.maxRetries = m
	}
}

// WithRetryWait sets the minimum and maximum wait durations for retrying failed requests. If the provided minimum wait is less than 0, it will be set to 0. If the provided maximum wait is less than 0, it will be set to 0. If the minimum wait is greater than the maximum wait, the maximum wait will be set to the minimum wait.
func WithRetryWait(minWait, maxWait time.Duration) RetryOption {
	return func(o *retryOptions) {
		if minWait < 0 {
			minWait = 0
		}

		if maxWait < 0 {
			maxWait = 0
		}

		if minWait > maxWait {
			maxWait = minWait
		}

		o.waitMin = minWait
		o.waitMax = maxWait
	}
}

// WithRetryJitter sets the jitter duration for retrying failed requests. If the provided jitter is less than 0, it will be set to 0.
func WithRetryJitter(jitter time.Duration) RetryOption {
	return func(o *retryOptions) {
		if jitter < 0 {
			jitter = 0
		}

		o.jitter = jitter
	}
}

// WithRetryNoRetryErrors sets errors that should not be retried. An error is matched using [errors.Is], so wrapped errors are matched too.
func WithRetryNoRetryErrors(errs ...error) RetryOption {
	return func(o *retryOptions) {
		o.noRetryErrors = slices.Clone(errs)
	}
}

// NewRetryTransport returns a [http.RoundTripper] that wraps the given [http.RoundTripper] and retries failed requests using an exponential backoff strategy. The retry behavior can be configured using the provided [RetryOption] functions.
func NewRetryTransport(inner http.RoundTripper, opts ...RetryOption) http.RoundTripper {
	o := retryOptions{
		maxRetries: DefaultRetryMaxRetries,
		waitMin:    DefaultRetryWaitMin,
		waitMax:    DefaultRetryWaitMax,
		jitter:     DefaultRetryJitter,
	}

	for _, opt := range opts {
		opt(&o)
	}

	if o.maxRetries <= 0 {
		return inner
	}

	return &retryTransport{
		inner: inner,
		opts:  o,
	}
}

// retryTransport is a [http.RoundTripper] that retries failed requests with an exponential backoff, honouring any Retry-After wait the server sends.
type retryTransport struct {
	inner http.RoundTripper
	opts  retryOptions
}

// RoundTrip implements the [http.RoundTripper] interface for the [retryTransport]. It retries failed requests using the configured backoff strategy.
func (t *retryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	backoff := &retryAfterBackoff{inner: newBackoff(t.opts)}
	var bo retry.Backoff = backoff
	if deadline, ok := req.Context().Deadline(); ok {
		bo = &deadlineBackoff{inner: backoff, deadline: deadline}
	}

	var resp *http.Response
	tries := 0

	err := retry.Do(req.Context(), bo, func(ctx context.Context) error {
		// If there was a previous response, close it so the connection can be reused
		if resp != nil {
			if resp.Body != nil {
				_ = resp.Body.Close()
			}
			resp = nil
		}

		// Shallow copy so we never mutate the caller's Request
		reqCopy := new(*req)

		// If this is a retry, rewind the request body so it can be sent again
		if tries > 0 {
			if err := rewindRequestBody(reqCopy); err != nil {
				return fmt.Errorf("failed to rewind request body: %w", err)
			}
		}
		tries++

		var respErr error
		resp, respErr = t.inner.RoundTrip(reqCopy)
		if respErr != nil {
			if resp != nil && resp.Body != nil {
				_ = resp.Body.Close()
			}
			resp = nil
			if t.retryableError(req, respErr) && canRewind(req) {
				return retry.RetryableError(respErr)
			}
			return respErr
		}

		if ok, wait := retryableResponse(req, resp); ok && canRewind(req) {
			if wait != nil {
				if *wait > t.opts.waitMax {
					return nil
				}

				backoff.overrideMin(*wait)
			}

			return retry.RetryableError(errRetryServerError)
		}

		return nil
	})

	// retry.Do checks the context before the first attempt, so if it was already done the inner transport never ran and never closed the body
	if tries == 0 && req.Body != nil {
		_ = req.Body.Close()
	}

	// Cleanup error used for retryable 5xx responses so we don't return a server error when the request was actually successful
	if errors.Is(err, errRetryServerError) {
		err = nil
	}

	// Clean up any lingering 5xx response if the context was canceled during backoff sleep
	if err != nil && resp != nil {
		if resp.Body != nil {
			_ = resp.Body.Close()
		}
		resp = nil
	}

	return resp, err
}

// retryableError returns true if the given request and error are retryable. It checks for the request's context being done, the configured no retry errors, TLS certificate verification errors, network dial errors, and idempotent requests.
func (t *retryTransport) retryableError(req *http.Request, err error) bool {
	if err == nil {
		return false
	}

	// Transport timeouts (dial, TLS handshake, response header) match context.DeadlineExceeded too, so only the request's own context being done is final.
	if req.Context().Err() != nil {
		return false
	}

	for _, noRetryErr := range t.opts.noRetryErrors {
		if errors.Is(err, noRetryErr) {
			return false
		}
	}

	if _, ok := errors.AsType[*tls.CertificateVerificationError](err); ok {
		return false
	}

	if opErr, ok := errors.AsType[*net.OpError](err); ok && opErr.Op == "dial" {
		return true
	}

	return isIdempotent(req)
}

// newBackoff returns a [retry.Backoff] configured according to the given RetryOptions.
func newBackoff(opts retryOptions) retry.Backoff {
	// Required to stop panic
	waitMin := opts.waitMin
	if waitMin <= 0 {
		waitMin = time.Nanosecond
	}

	b := retry.NewExponential(waitMin)
	b = retry.WithCappedDuration(opts.waitMax, b)
	b = retry.WithJitter(opts.jitter, b)
	b = retry.WithMaxRetries(uint64(opts.maxRetries), b)

	return b
}

// retryAfterBackoff is a [retry.Backoff] that wraps another backoff and allows for a per try minimum wait to be set, which is used to implement retry after semantics.
type retryAfterBackoff struct {
	inner    retry.Backoff
	override time.Duration
}

// overrideMin sets the minimum wait duration for the next retry attempt. This is used to implement retry after semantics, where the server may indicate a minimum wait time before retrying.
func (b *retryAfterBackoff) overrideMin(d time.Duration) {
	if d < 0 {
		d = 0
	}

	b.override = d
}

// Next implements the [retry.Backoff] interface for the [retryAfterBackoff]. It returns the next wait duration and whether to stop retrying. It ensures that the wait duration is at least the retry after wait duration set by [overrideMin].
func (b *retryAfterBackoff) Next() (time.Duration, bool) {
	d, stop := b.inner.Next()
	if stop {
		return 0, true
	}

	o := b.override
	b.override = 0

	return max(d, o), false
}

// deadlineBackoff is a [retry.Backoff] that stops retrying when the next wait would not complete before the deadline, so the caller gets the last real response or error rather than a context deadline error.
type deadlineBackoff struct {
	inner    retry.Backoff
	deadline time.Time
}

// Next implements the [retry.Backoff] interface for the [deadlineBackoff]. It stops if sleeping for the next wait would reach the deadline.
func (b *deadlineBackoff) Next() (time.Duration, bool) {
	d, stop := b.inner.Next()
	if stop {
		return 0, true
	}

	if !b.deadline.After(time.Now().Add(d)) {
		return 0, true
	}

	return d, false
}
