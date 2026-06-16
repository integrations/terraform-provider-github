package ghclient

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"golang.org/x/sync/semaphore"
)

type testRoundTripper struct {
	called atomic.Int32
	resp   *http.Response
	err    error
}

func (r *testRoundTripper) RoundTrip(_ *http.Request) (*http.Response, error) {
	r.called.Add(1)
	if r.err != nil {
		return nil, r.err
	}

	return r.resp, nil
}

func Test_throttlerReadCloser_Close_releases_once(t *testing.T) {
	t.Parallel()

	sema := semaphore.NewWeighted(1)
	if err := sema.Acquire(t.Context(), 1); err != nil {
		t.Fatalf("failed to acquire semaphore for setup: %v", err)
	}

	rc := &throttlerReadCloser{
		ReadCloser: io.NopCloser(strings.NewReader("ok")),
		sema:       sema,
	}

	if err := rc.Close(); err != nil {
		t.Fatalf("failed to close read closer: %v", err)
	}

	if err := rc.Close(); err != nil {
		t.Fatalf("failed to close read closer on second close: %v", err)
	}

	ctx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
	defer cancel()

	if err := sema.Acquire(ctx, 1); err != nil {
		t.Fatalf("expected semaphore to be released after close, got error: %v", err)
	}

	ctx2, cancel2 := context.WithTimeout(t.Context(), 10*time.Millisecond)
	defer cancel2()

	if err := sema.Acquire(ctx2, 1); err == nil {
		t.Fatal("expected second acquire to fail because close should release exactly once")
	}
}

func Test_throttler_RoundTrip_acquire_error(t *testing.T) {
	t.Parallel()

	inner := &testRoundTripper{}
	tr := &throttler{
		sema:  semaphore.NewWeighted(1),
		inner: inner,
	}

	ctx, cancel := context.WithCancel(t.Context())
	cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://example.com", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	_, err = tr.RoundTrip(req)
	if err == nil {
		t.Fatal("expected acquire error from canceled context")
	}

	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context canceled error, got %v", err)
	}

	if inner.called.Load() != 0 {
		t.Fatalf("expected inner transport not to be called, got %d calls", inner.called.Load())
	}
}

func Test_throttler_RoundTrip_inner_error_releases_semaphore(t *testing.T) {
	t.Parallel()

	inner := &testRoundTripper{err: errors.New("boom")}
	sema := semaphore.NewWeighted(1)
	tr := &throttler{
		sema:  sema,
		inner: inner,
	}

	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "https://example.com", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	_, err = tr.RoundTrip(req)
	if err == nil {
		t.Fatal("expected round trip to fail")
	}

	ctx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
	defer cancel()

	if err := sema.Acquire(ctx, 1); err != nil {
		t.Fatalf("expected semaphore to be released when inner transport fails, got %v", err)
	}

	if inner.called.Load() != 1 {
		t.Fatalf("expected inner transport to be called once, got %d", inner.called.Load())
	}
}

func Test_throttler_RoundTrip_success_releases_on_body_close(t *testing.T) {
	t.Parallel()

	inner := &testRoundTripper{resp: &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader("ok"))}}
	sema := semaphore.NewWeighted(1)
	tr := &throttler{
		sema:  sema,
		inner: inner,
	}

	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "https://example.com", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err := tr.RoundTrip(req)
	if err != nil {
		t.Fatalf("expected round trip to succeed, got error: %v", err)
	}

	blockedCtx, blockedCancel := context.WithTimeout(t.Context(), 10*time.Millisecond)
	defer blockedCancel()

	if err := sema.Acquire(blockedCtx, 1); err == nil {
		t.Fatal("expected acquire to block before response body is closed")
	}

	if err := resp.Body.Close(); err != nil {
		t.Fatalf("failed to close response body: %v", err)
	}

	ctx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
	defer cancel()

	if err := sema.Acquire(ctx, 1); err != nil {
		t.Fatalf("expected semaphore to be released after closing response body, got %v", err)
	}

	if inner.called.Load() != 1 {
		t.Fatalf("expected inner transport to be called once, got %d", inner.called.Load())
	}
}
