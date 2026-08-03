package ghclient

import (
	"io"
	"net/http"
	"sync"

	"golang.org/x/sync/semaphore"
)

// throttlerReadCloser is a wrapper around an io.ReadCloser that releases a semaphore weight when the ReadCloser is closed. This is used to ensure that the semaphore controlling concurrent requests is properly released after the response body has been fully read and closed, preventing resource leaks and allowing other requests to proceed.
type throttlerReadCloser struct {
	io.ReadCloser
	sema *semaphore.Weighted
	once sync.Once
}

// Close releases the semaphore weight when the ReadCloser is closed. It ensures that the release operation is only performed once, even if Close is called multiple times, preventing potential double-release issues that could lead to incorrect semaphore state.
func (c *throttlerReadCloser) Close() error {
	err := c.ReadCloser.Close()
	c.once.Do(func() {
		c.sema.Release(1)
	})
	return err
}

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

	res, err := t.inner.RoundTrip(req)
	if err != nil {
		t.sema.Release(1)
		return nil, err
	}

	res.Body = &throttlerReadCloser{
		ReadCloser: res.Body,
		sema:       t.sema,
	}

	return res, nil
}
