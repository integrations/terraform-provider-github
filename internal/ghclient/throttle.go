package ghclient

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/semaphore"
)

// newThrottleTransport returns a [throttleTransport] that wraps the given [http.RoundTripper] using the given [ConcurrencyOptions] to configure the concurrency.
func newThrottleTransport(inner http.RoundTripper, opts ConcurrencyOptions) (http.RoundTripper, error) {
	if opts.Max == 0 {
		return inner, nil
	}

	if opts.Max < 0 {
		return nil, fmt.Errorf("max must be positive")
	}

	if opts.sema == nil {
		opts.sema = semaphore.NewWeighted(opts.Max)
	}

	return &throttleTransport{
		inner: inner,
		sema:  opts.sema,
	}, nil
}

// throttleTransport is a [http.RoundTripper] that wraps another http.RoundTripper and limits the number of concurrent requests using a [semaphore.Weighted].
type throttleTransport struct {
	inner http.RoundTripper
	sema  *semaphore.Weighted
}

// RoundTrip implements the [http.RoundTripper] interface for the [throttleTransport]. It throttles the number of concurrent requests using the semaphore and delegates to the inner http.RoundTripper.
func (t *throttleTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if err := t.sema.Acquire(req.Context(), 1); err != nil {
		return nil, err
	}
	defer t.sema.Release(1)
	return t.inner.RoundTrip(req)
}
