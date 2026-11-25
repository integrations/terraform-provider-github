package httpx

import (
	"context"
	"net/http"

	"golang.org/x/sync/semaphore"
)

// ThrottleLimiter is an interface that defines methods for acquiring and releasing concurrency limits. It is used to control the number of concurrent requests made to a resource, such as an API, to prevent overwhelming the server or exceeding rate limits.
type ThrottleLimiter interface {
	Acquire(ctx context.Context, n int64) error
	Release(n int64)
}

// NewThrottleLimiter creates a new ThrottleLimiter with the specified maximum concurrency. If the maximum concurrency is less than or equal to zero, it returns a no-op ThrottleLimiter that does not enforce any concurrency limits.
func NewThrottleLimiter(m int64) ThrottleLimiter {
	if m <= 0 {
		return &noopThrottleLimiter{}
	}

	return semaphore.NewWeighted(m)
}

// noopThrottleLimiter is a ThrottleLimiter that does nothing. It is used when no concurrency limit is desired.
type noopThrottleLimiter struct{}

// Acquire is a no-op implementation of the ThrottleLimiter interface. It always returns nil, indicating that the request can proceed without any concurrency limit.
func (*noopThrottleLimiter) Acquire(_ context.Context, _ int64) error {
	return nil
}

// Release is a no-op implementation of the ThrottleLimiter interface. It does nothing, as there is no concurrency limit to release.
func (*noopThrottleLimiter) Release(_ int64) {}

// NewThrottleTransport returns a [throttleTransport] that wraps the given [http.RoundTripper] using the given [ThrottleLimiter] to control concurrency.
func NewThrottleTransport(inner http.RoundTripper, limiter ThrottleLimiter) http.RoundTripper {
	if limiter == nil {
		return inner
	}

	return &throttleTransport{
		inner:   inner,
		limiter: limiter,
	}
}

// throttleTransport is a [http.RoundTripper] that wraps another http.RoundTripper and limits the number of concurrent requests using a [ThrottleLimiter].
type throttleTransport struct {
	inner   http.RoundTripper
	limiter ThrottleLimiter
}

// RoundTrip implements the [http.RoundTripper] interface for the [throttleTransport]. It throttles the number of concurrent requests using the limiter and delegates to the inner http.RoundTripper.
func (t *throttleTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if err := t.limiter.Acquire(req.Context(), 1); err != nil {
		if req.Body != nil {
			_ = req.Body.Close()
		}

		return nil, err
	}
	defer t.limiter.Release(1)

	return t.inner.RoundTrip(req)
}
