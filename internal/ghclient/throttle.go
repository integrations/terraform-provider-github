package ghclient

import (
	"net/http"

	"golang.org/x/sync/semaphore"
)

// throttler is an HTTP RoundTripper that limits the number of concurrent requests to a specified maximum. It uses a weighted semaphore to control access to the underlying RoundTripper, ensuring that no more than the allowed number of requests are in flight at any given time. This is useful for preventing overwhelming a server or API with too many simultaneous requests.
type throttler struct {
	sema  *semaphore.Weighted
	inner http.RoundTripper
}

// RoundTrip implements the http.RoundTripper interface for the throttler. It acquires a semaphore weight before proceeding with the request, ensuring that the number of concurrent requests does not exceed the specified limit. After the request is completed, it releases the semaphore weight, allowing other requests to proceed. If acquiring the semaphore fails, it returns an error.
func (t *throttler) RoundTrip(req *http.Request) (*http.Response, error) {
	if err := t.sema.Acquire(req.Context(), 1); err != nil {
		return nil, err
	}
	defer t.sema.Release(1)
	return t.inner.RoundTrip(req)
}
