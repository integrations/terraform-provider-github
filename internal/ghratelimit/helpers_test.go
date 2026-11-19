package ghratelimit

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync/atomic"
	"testing"
	"time"
)

func mustCreateRequest(t *testing.T, ctx context.Context, method, url string, body io.Reader) *http.Request {
	t.Helper()

	if ctx == nil {
		ctx = t.Context()
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	return req
}

type closeRecorder struct {
	closed atomic.Bool
}

func (c *closeRecorder) Read(p []byte) (int, error) { return 0, io.EOF }

func (c *closeRecorder) Close() error {
	c.closed.Store(true)
	return nil
}

type dummyRoundTripper struct {
	http.RoundTripper
}

type testRoundTripper struct {
	resp *http.Response
	err  error
}

func (tr *testRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return tr.resp, tr.err
}

type errorReadSeeker struct {
	io.Closer
}

func (rs *errorReadSeeker) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("error reading")
}

func (rs *errorReadSeeker) Seek(offset int64, whence int) (int64, error) {
	return 0, fmt.Errorf("error seeking")
}

type dummyLimiter struct {
	Limiter
	waitUntil time.Time
}

func (rl *dummyLimiter) BlockedUntil(category string) time.Time {
	return rl.waitUntil
}
