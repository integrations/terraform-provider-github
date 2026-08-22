package ghclient

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"testing/synctest"
	"time"

	"golang.org/x/sync/semaphore"
)

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

		resp, err := tr.RoundTrip(req)
		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}

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

		req := mustCreateRequest(t, http.MethodGet, "https://example.com")

		resp, err := tr.RoundTrip(req)
		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}

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
		tr := &throttler{
			sema:  semaphore.NewWeighted(1),
			inner: inner,
		}

		req := mustCreateRequest(t, http.MethodGet, "https://example.com")

		resp, err := tr.RoundTrip(req)
		if err != nil {
			t.Fatalf("expected round trip to succeed, got error: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("expected status code 200 OK, got %d", resp.StatusCode)
		}

		if inner.called.Load() != 1 {
			t.Fatalf("expected inner transport to be called once, got %d calls", inner.called.Load())
		}
	})

	t.Run("success_after_no_body_close", func(t *testing.T) {
		t.Parallel()

		inner := &testRoundTripper{err: errors.New("boom")}
		sema := semaphore.NewWeighted(1)
		tr := &throttler{
			sema:  sema,
			inner: inner,
		}

		req := mustCreateRequest(t, http.MethodGet, "https://example.com")

		_, err := tr.RoundTrip(req)
		if err == nil || !errors.Is(err, inner.err) {
			t.Fatalf("expected round trip to error with %v, got %v", inner.err, err)
		}

		if !sema.TryAcquire(1) {
			t.Fatal("semaphore permit leaked after body not closed")
		}
	})

	t.Run("throttles_concurrent_requests", func(t *testing.T) {
		t.Parallel()

		synctest.Test(t, func(t *testing.T) {
			reqs := 5
			inner := &testRoundTripper{delay: 1 * time.Second, resp: &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader("ok"))}}
			tr := &throttler{
				sema:  semaphore.NewWeighted(1),
				inner: inner,
			}

			for i := range reqs {
				req := mustCreateRequest(t, http.MethodGet, "https://example.com")
				go func() {
					req := req
					resp, err := tr.RoundTrip(req)
					if err != nil {
						t.Errorf("request %d: %v", i, err)
						return
					}
					resp.Body.Close()

					if resp.StatusCode != http.StatusOK {
						t.Errorf("request %d: expected status code 200 OK, got %d", i, resp.StatusCode)
					}
				}()
			}

			// Both goroutines are now durably blocked (one sleeping, others on sema.Acquire).
			synctest.Wait()
			if inner.called.Load() != 0 {
				t.Fatal("expected no completions before time advances")
			}

			for i := range reqs {
				// Clock jumps 1s: sleeping goroutine wakes, completes, releases semaphore; next goroutine acquires and sleeps.
				time.Sleep(time.Second)
				synctest.Wait()
				if inner.called.Load() != int32(i+1) {
					t.Fatal("expected completions to be throttled to one per second")
				}
			}
		})
	})
}
