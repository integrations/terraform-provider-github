package ghclient

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"golang.org/x/sync/semaphore"
)

func Test_throttlerReadCloser_Close(t *testing.T) {
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
}

func Test_throttler_RoundTrip(t *testing.T) {
	t.Parallel()

	t.Run("handles_acquire_error", func(t *testing.T) {
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
	})

	t.Run("handles_inner_error", func(t *testing.T) {
		t.Parallel()

		inner := &testRoundTripper{err: errors.New("boom")}
		tr := &throttler{
			sema:  semaphore.NewWeighted(1),
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

		if !errors.Is(err, inner.err) {
			t.Fatalf("expected inner transport error, got %v", err)
		}

		if inner.called.Load() != 1 {
			t.Fatalf("expected inner transport to be called once, got %d calls", inner.called.Load())
		}
	})

	t.Run("success", func(t *testing.T) {
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

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("expected status code 200 OK, got %d", resp.StatusCode)
		}

		if inner.called.Load() != 1 {
			t.Fatalf("expected inner transport to be called once, got %d calls", inner.called.Load())
		}

		if ok := sema.TryAcquire(1); ok {
			t.Fatal("expected semaphore to be held until response body is closed")
		}

		if err := resp.Body.Close(); err != nil {
			t.Fatalf("failed to close response body: %v", err)
		}

		if ok := sema.TryAcquire(1); !ok {
			t.Fatal("expected semaphore to be released after closing response body")
		}
	})
}
