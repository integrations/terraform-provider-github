package httpx

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

func TestNewThrottleLimiter(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		max  int64
	}{
		{
			name: "returns_noop_with_max_zero",
			max:  0,
		},
		{
			name: "returns_noop_with_max_negative",
			max:  -1,
		},
		{
			name: "returns_semaphone_with_positive_max",
			max:  1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := NewThrottleLimiter(tt.max)

			if got == nil {
				t.Fatal("limiter: got nil, want non-nil")
			}

			switch got.(type) {
			case *noopThrottleLimiter:
				if tt.max > 0 {
					t.Fatal("limiter: got *noopThrottleLimiter, want *semaphore.Weighted")
				}
			case *semaphore.Weighted:
				if tt.max <= 0 {
					t.Fatal("limiter: got *semaphore.Weighted, want *noopThrottleLimiter")
				}
			default:
				t.Fatalf("limiter: got %T, want *noopThrottleLimiter or *semaphore.Weighted", got)
			}
		})
	}
}

func Test_noopThrottleLimiter(t *testing.T) {
	t.Parallel()

	limiter := &noopThrottleLimiter{}

	ctx, cancel := context.WithCancel(t.Context())
	cancel()

	for range 3 {
		if err := limiter.Acquire(ctx, 1); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	limiter.Release(1)
}

func TestNewThrottleTransport(t *testing.T) {
	t.Parallel()

	t.Run("returns_inner_if_limiter_is_nil", func(t *testing.T) {
		t.Parallel()

		inner := &dummyRoundTripper{}
		got := NewThrottleTransport(inner, nil)

		if got != inner {
			t.Fatal("transport: got non-inner transport, want inner transport")
		}
	})

	t.Run("returns_throttleTransport_if_limiter_is_not_nil", func(t *testing.T) {
		t.Parallel()

		inner := &dummyRoundTripper{}
		limiter := &noopThrottleLimiter{}
		got := NewThrottleTransport(inner, limiter)

		if got == nil {
			t.Fatal("transport: got nil, want non-nil")
		}

		ttr, ok := got.(*throttleTransport)
		if !ok {
			t.Fatal("transport: got non-throttleTransport, want throttleTransport")
		}

		if ttr.inner != inner {
			t.Fatal("inner: got other transport, want given transport")
		}

		if ttr.limiter != limiter {
			t.Fatal("limiter: got other limiter, want given limiter")
		}
	})
}

func Test_throttleTransport(t *testing.T) {
	t.Parallel()

	t.Run("RoundTrip", func(t *testing.T) {
		t.Parallel()

		t.Run("releases_limiter", func(t *testing.T) {
			t.Parallel()

			inner := &testRoundTripper{resp: &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader("ok"))}}
			limiter := semaphore.NewWeighted(1)
			tr := &throttleTransport{
				inner:   inner,
				limiter: limiter,
			}

			req := mustCreateRequest(t, nil, "GET", "http://example.com", nil)

			resp, err := tr.RoundTrip(req)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if resp != nil && resp.Body != nil {
				defer resp.Body.Close()
			}

			if !limiter.TryAcquire(1) {
				t.Fatal("limiter released: got false, want true")
			}
		})

		t.Run("throttles_concurrent_requests", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				reqs := 5
				inner := &testRoundTripper{delay: 1 * time.Second, resp: &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader("ok"))}}
				tr := &throttleTransport{
					inner:   inner,
					limiter: semaphore.NewWeighted(1),
				}

				for i := range reqs {
					req := mustCreateRequest(t, nil, http.MethodGet, "https://example.com", nil)
					go func() {
						req := req
						resp, err := tr.RoundTrip(req)
						if err != nil {
							t.Errorf("request %d: %v", i, err)
							return
						}
						resp.Body.Close()

						if resp.StatusCode != http.StatusOK {
							t.Errorf("request %d: status code: got %d, want 200", i, resp.StatusCode)
						}
					}()
				}

				// All goroutines are now durably blocked (one sleeping, others on sema.Acquire).
				synctest.Wait()
				if inner.called.Load() != 0 {
					t.Fatalf("calls before time advances: got %d, want 0", inner.called.Load())
				}

				for i := range reqs {
					// Clock jumps 1s: sleeping goroutine wakes, completes, releases semaphore; next goroutine acquires and sleeps.
					time.Sleep(time.Second)
					synctest.Wait()
					if inner.called.Load() != int32(i+1) {
						t.Fatalf("calls: got %d, want %d", inner.called.Load(), i+1)
					}
				}
			})
		})

		t.Run("handles_acquire_error", func(t *testing.T) {
			t.Parallel()

			inner := &dummyRoundTripper{}
			limiter := semaphore.NewWeighted(1)
			tr := &throttleTransport{
				inner:   inner,
				limiter: limiter,
			}

			ctx, cancel := context.WithCancel(t.Context())
			cancel()

			req := mustCreateRequest(t, ctx, "GET", "http://example.com", nil)
			rec := &closeRecorder{}
			req.Body = rec

			resp, err := tr.RoundTrip(req)
			if err == nil {
				t.Fatal("got nil, want error")
			}
			if resp != nil {
				t.Fatal("response: got non-nil, want nil")
			}
			if !rec.closed.Load() {
				t.Error("request body closed: got false, want true")
			}
		})

		t.Run("handles_inner_error", func(t *testing.T) {
			t.Parallel()

			inner := &testRoundTripper{err: errors.New("boom")}
			limiter := semaphore.NewWeighted(1)
			tr := &throttleTransport{
				inner:   inner,
				limiter: limiter,
			}

			req := mustCreateRequest(t, nil, http.MethodGet, "https://example.com", nil)

			resp, err := tr.RoundTrip(req)
			if err == nil {
				t.Fatal("got nil, want error")
			}

			if resp != nil {
				t.Fatal("response: got non-nil, want nil")
			}

			if !limiter.TryAcquire(1) {
				t.Fatal("limiter released: got false, want true")
			}
		})
	})
}
