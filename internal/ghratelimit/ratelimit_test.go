package ghratelimit

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"testing/synctest"
	"time"
)

func TestWaitExceededError(t *testing.T) {
	t.Parallel()

	t.Run("Error", func(t *testing.T) {
		t.Parallel()

		for _, tt := range []struct {
			name string
			err  *WaitExceededError
			want string
		}{
			{
				name: "formats_category_until_and_max_wait",
				err:  &WaitExceededError{Category: rateLimitCategoryCore, Until: time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC), MaxWait: 30 * time.Second},
				want: "rate limit wait exceeded: core rate limit is blocked until 2026-01-02T03:04:05Z, which is longer than the maximum wait of 30s",
			},
			{
				name: "formats_until_in_utc",
				err:  &WaitExceededError{Category: rateLimitCategorySecondary, Until: time.Date(2026, 1, 2, 5, 4, 5, 0, time.FixedZone("UTC+2", 2*60*60)), MaxWait: time.Minute},
				want: "rate limit wait exceeded: secondary rate limit is blocked until 2026-01-02T03:04:05Z, which is longer than the maximum wait of 1m0s",
			},
			{
				name: "rounds_until_down_to_second",
				err:  &WaitExceededError{Category: rateLimitCategoryCore, Until: time.Date(2026, 1, 2, 3, 4, 5, 999_000_000, time.UTC), MaxWait: 30 * time.Second},
				want: "rate limit wait exceeded: core rate limit is blocked until 2026-01-02T03:04:05Z, which is longer than the maximum wait of 30s",
			},
			{
				name: "rounds_max_wait_to_nearest_second",
				err:  &WaitExceededError{Category: rateLimitCategorySearch, Until: time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC), MaxWait: 9500 * time.Millisecond},
				want: "rate limit wait exceeded: search rate limit is blocked until 2026-01-02T03:04:05Z, which is longer than the maximum wait of 10s",
			},
			{
				name: "formats_zero_max_wait",
				err:  &WaitExceededError{Category: rateLimitCategoryGraphQL, Until: time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)},
				want: "rate limit wait exceeded: graphql rate limit is blocked until 2026-01-02T03:04:05Z, which is longer than the maximum wait of 0s",
			},
		} {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				if got := tt.err.Error(); got != tt.want {
					t.Errorf("got %q, want %q", got, tt.want)
				}
			})
		}
	})

	t.Run("Unwrap", func(t *testing.T) {
		t.Parallel()

		for _, tt := range []struct {
			name   string
			err    error
			target error
			want   bool
		}{
			{
				name:   "matches_sentinel",
				err:    &WaitExceededError{Category: rateLimitCategoryCore},
				target: ErrRateLimitWaitExceeded,
				want:   true,
			},
			{
				name:   "matches_sentinel_when_wrapped",
				err:    fmt.Errorf("request failed: %w", &WaitExceededError{Category: rateLimitCategoryCore}),
				target: ErrRateLimitWaitExceeded,
				want:   true,
			},
			{
				name:   "does_not_match_other_error",
				err:    &WaitExceededError{Category: rateLimitCategoryCore},
				target: context.DeadlineExceeded,
				want:   false,
			},
			{
				name:   "does_not_match_other_error_when_wrapped",
				err:    fmt.Errorf("request failed: %w", &WaitExceededError{Category: rateLimitCategoryCore}),
				target: context.DeadlineExceeded,
				want:   false,
			},
		} {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				if got := errors.Is(tt.err, tt.target); got != tt.want {
					t.Errorf("errors.Is(%v, %v): got %v, want %v", tt.err, tt.target, got, tt.want)
				}
			})
		}
	})
}

func TestNewLimiter(t *testing.T) {
	t.Parallel()

	rl := NewLimiter()
	if rl == nil {
		t.Fatal("got nil, want non-nil")
	}
}

func Test_limiter(t *testing.T) {
	t.Parallel()

	t.Run("Block", func(t *testing.T) {
		t.Parallel()

		cur := time.Now()

		for _, tt := range []struct {
			name     string
			cur      *time.Time
			category string
			until    time.Time
			want     time.Time
		}{
			{
				name:     "sets_until_for_new_category",
				category: "test",
				until:    cur.Add(5 * time.Second),
				want:     cur.Add(5 * time.Second),
			},
			{
				name:     "sets_until_for_existing_category_with_lower_wait",
				cur:      &cur,
				category: "test",
				until:    cur.Add(5 * time.Second),
				want:     cur.Add(5 * time.Second),
			},
			{
				name:     "leaves_until_for_existing_category_with_newer_wait",
				cur:      new(cur.Add(5 * time.Second)),
				category: "test",
				until:    cur,
				want:     cur.Add(5 * time.Second),
			},
		} {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				rl := &limiter{blockedUntil: make(map[string]time.Time)}

				if tt.cur != nil {
					rl.blockedUntil[tt.category] = *tt.cur
				}

				rl.Block(tt.category, tt.until)

				if got, ok := rl.blockedUntil[tt.category]; !ok || !got.Equal(tt.want) {
					t.Errorf("category %q blocked until: got %v, want %v", tt.category, got, tt.want)
				}
			})
		}
	})

	t.Run("BlockedUntil", func(t *testing.T) {
		t.Parallel()

		for _, tt := range []struct {
			name string
			wait map[string]time.Duration
			cat  string
			want *time.Duration
		}{
			{
				name: "returns_zero_if_not_blocked",
				cat:  "test",
			},
			{
				name: "returns_blocked_until_for_category",
				wait: map[string]time.Duration{
					"test": 5 * time.Second,
				},
				cat:  "test",
				want: new(5 * time.Second),
			},
			{
				name: "returns_blocked_until_for_secondary_category_if_later",
				wait: map[string]time.Duration{
					"test":                     5 * time.Second,
					rateLimitCategorySecondary: 10 * time.Second,
				},
				cat:  "test",
				want: new(10 * time.Second),
			},
			{
				name: "returns_blocked_until_for_category_if_later_than_secondary",
				wait: map[string]time.Duration{
					"test":                     10 * time.Second,
					rateLimitCategorySecondary: 5 * time.Second,
				},
				cat:  "test",
				want: new(10 * time.Second),
			},
		} {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				synctest.Test(t, func(t *testing.T) {
					rl := &limiter{blockedUntil: make(map[string]time.Time)}

					for cat, wait := range tt.wait {
						rl.blockedUntil[cat] = time.Now().Add(wait)
					}

					got := rl.BlockedUntil(tt.cat)

					if tt.want == nil {
						if got != (time.Time{}) {
							t.Errorf("category %q blocked until: got %v, want zero", tt.cat, got)
						}
						return
					}

					want := time.Now().Add(*tt.want)
					if !got.Equal(want) {
						t.Errorf("category %q blocked until: got %v, want %v", tt.cat, got, want)
					}
				})
			})
		}
	})
}

func TestWithDefaultSecondaryWait(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		d    time.Duration
		want time.Duration
	}{
		{
			name: "sets_default_secondary_wait",
			d:    5 * time.Second,
			want: 5 * time.Second,
		},
		{
			name: "sets_default_secondary_wait_to_zero_if_negative",
			d:    -1 * time.Second,
			want: 0,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			opts := &options{}
			WithDefaultSecondaryWait(tt.d)(opts)

			if opts.defaultSecondaryWait != tt.want {
				t.Errorf("got %v, want %v", opts.defaultSecondaryWait, tt.want)
			}
		})
	}
}

func TestWithWaitMax(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		d    time.Duration
		want time.Duration
	}{
		{
			name: "sets_wait_max",
			d:    10 * time.Second,
			want: 10 * time.Second,
		},
		{
			name: "sets_wait_max_to_zero_if_negative",
			d:    -1 * time.Second,
			want: 0,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			opts := &options{}
			WithWaitMax(tt.d)(opts)

			if opts.waitMax != tt.want {
				t.Errorf("got %v, want %v", opts.waitMax, tt.want)
			}
		})
	}
}

func TestWithBasePath(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		p    string
		want string
	}{
		{
			name: "sets_empty_base_path",
			p:    "",
			want: "",
		},
		{
			name: "sets_base_path",
			p:    "/api/v3",
			want: "/api/v3",
		},
		{
			name: "trims_trailing_slash",
			p:    "/api/v3/",
			want: "/api/v3",
		},
		{
			name: "adds_leading_slash",
			p:    "api/v3",
			want: "/api/v3",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			opts := &options{}
			WithBasePath(tt.p)(opts)

			if opts.basePath != tt.want {
				t.Errorf("got %q, want %q", opts.basePath, tt.want)
			}
		})
	}
}

func TestWithLimiter(t *testing.T) {
	t.Parallel()

	t.Run("sets_custom_rate_limiter", func(t *testing.T) {
		t.Parallel()

		rl := NewLimiter()
		opts := &options{}
		WithLimiter(rl)(opts)

		if opts.limiter != rl {
			t.Errorf("got %v, want %v", opts.limiter, rl)
		}
	})

	t.Run("does_not_set_nil_rate_limiter", func(t *testing.T) {
		t.Parallel()

		opts := &options{}
		WithLimiter(nil)(opts)

		if opts.limiter != nil {
			t.Errorf("got %v, want nil", opts.limiter)
		}
	})
}

func TestNewTransport(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name                     string
		opts                     []Option
		wantDefaultSecondaryWait time.Duration
		wantWaitMax              time.Duration
		wantBasePath             string
	}{
		{
			name:                     "with_defaults",
			opts:                     []Option{},
			wantDefaultSecondaryWait: 60 * time.Second,
			wantWaitMax:              60 * time.Second,
		},
		{
			name: "with_custom_options",
			opts: []Option{
				WithDefaultSecondaryWait(5 * time.Second),
				WithWaitMax(10 * time.Second),
				WithBasePath("/api/v3"),
			},
			wantDefaultSecondaryWait: 5 * time.Second,
			wantWaitMax:              10 * time.Second,
			wantBasePath:             "/api/v3",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			inner := &dummyRoundTripper{}
			got := NewTransport(inner, tt.opts...)

			if got == nil {
				t.Fatal("transport: got nil, want non-nil")
			}

			if got == inner {
				t.Fatal("transport: got inner transport, want wrapping transport")
			}

			tr, ok := got.(*rateLimitTransport)
			if !ok {
				t.Fatalf("transport: got %T, want *rateLimitTransport", got)
			}

			if tr.defaultSecondaryWait != tt.wantDefaultSecondaryWait {
				t.Errorf("default secondary wait: got %v, want %v", tr.defaultSecondaryWait, tt.wantDefaultSecondaryWait)
			}

			if tr.waitMax != tt.wantWaitMax {
				t.Errorf("wait max: got %v, want %v", tr.waitMax, tt.wantWaitMax)
			}

			if tr.basePath != tt.wantBasePath {
				t.Errorf("base path: got %q, want %q", tr.basePath, tt.wantBasePath)
			}

			if tr.limiter == nil {
				t.Error("rate limiter: got nil, want non-nil")
			}
		})
	}
}

func Test_rateLimitTransport(t *testing.T) {
	t.Parallel()

	t.Run("RoundTrip", func(t *testing.T) {
		t.Parallel()

		t.Run("handles_wait_error", func(t *testing.T) {
			t.Parallel()

			inner := &dummyRoundTripper{}
			tr := &rateLimitTransport{
				inner:   inner,
				waitMax: 1 * time.Second,
				limiter: NewLimiter(),
			}
			tr.limiter.Block(rateLimitCategoryCore, time.Now().Add(2*time.Second))

			req := mustCreateRequest(t, nil, "GET", "http://example.com", nil)
			rec := &closeRecorder{}
			req.Body = rec

			resp, err := tr.RoundTrip(req)
			if err == nil {
				t.Fatal("got nil, want error")
			}

			if !errors.Is(err, ErrRateLimitWaitExceeded) {
				t.Fatalf("got %v, want %v", err, ErrRateLimitWaitExceeded)
			}

			if !rec.closed.Load() {
				t.Error("request body closed: got false, want true")
			}

			if resp != nil {
				t.Fatalf("response: got %v, want nil", resp)
			}
		})

		t.Run("handles_inner_error", func(t *testing.T) {
			t.Parallel()

			inner := &testRoundTripper{err: errors.New("inner error")}
			tr := &rateLimitTransport{
				inner:   inner,
				waitMax: 1 * time.Second,
				limiter: &dummyLimiter{},
			}

			req := mustCreateRequest(t, nil, "GET", "http://example.com", nil)

			resp, err := tr.RoundTrip(req)
			if err == nil {
				t.Fatal("got nil, want error")
			}

			if err.Error() != "inner error" {
				t.Fatalf("error: got %q, want %q", err, "inner error")
			}

			if resp != nil {
				t.Fatalf("response: got %v, want nil", resp)
			}
		})

		t.Run("closes_response_body_with_inner_error", func(t *testing.T) {
			t.Parallel()

			innerErr := errors.New("inner error")
			rec := &closeRecorder{}
			inner := &testRoundTripper{resp: &http.Response{StatusCode: http.StatusOK, Body: rec}, err: innerErr}
			tr := &rateLimitTransport{
				inner:   inner,
				waitMax: 1 * time.Second,
				limiter: NewLimiter(),
			}

			req := mustCreateRequest(t, nil, "GET", "http://example.com", nil)

			resp, err := tr.RoundTrip(req)
			if !errors.Is(err, innerErr) {
				t.Fatalf("got %v, want %v", err, innerErr)
			}

			if !rec.closed.Load() {
				t.Error("response body closed: got false, want true")
			}

			if resp != nil {
				t.Fatalf("response: got %v, want nil", resp)
			}
		})

		t.Run("handles_successful_response", func(t *testing.T) {
			t.Parallel()

			inner := &testRoundTripper{resp: &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader("ok"))}}
			tr := &rateLimitTransport{
				inner:   inner,
				waitMax: 1 * time.Second,
				limiter: NewLimiter(),
			}

			req := mustCreateRequest(t, nil, "GET", "http://example.com", nil)

			resp, err := tr.RoundTrip(req)
			if err != nil {
				t.Fatalf("unexpected error from RoundTrip: %v", err)
			}

			if resp == nil {
				t.Fatal("response: got nil, want non-nil")
			}

			if resp.StatusCode != http.StatusOK {
				t.Fatalf("status code: got %d, want %d", resp.StatusCode, http.StatusOK)
			}
		})

		t.Run("blocks_on_primary_rate_limit", func(t *testing.T) {
			t.Parallel()

			inner := &testRoundTripper{resp: &http.Response{
				StatusCode: http.StatusForbidden,
				Header: http.Header{
					headerRateRemaining: []string{"0"},
					headerRateReset:     []string{strconv.FormatInt(time.Now().Add(time.Hour).Unix(), 10)},
				},
			}}
			tr := &rateLimitTransport{inner: inner, waitMax: time.Second, limiter: NewLimiter()}

			req := mustCreateRequest(t, nil, "POST", "http://example.com/repos/test", nil)

			if _, err := tr.RoundTrip(req); err != nil {
				t.Fatalf("unexpected error from first RoundTrip: %v", err)
			}

			if _, err := tr.RoundTrip(req); !errors.Is(err, ErrRateLimitWaitExceeded) {
				t.Fatalf("got %v, want %v", err, ErrRateLimitWaitExceeded)
			}
		})

		t.Run("blocks_on_secondary_rate_limit", func(t *testing.T) {
			t.Parallel()

			inner := &testRoundTripper{resp: &http.Response{
				StatusCode: http.StatusTooManyRequests,
				Header: http.Header{
					headerRetryAfter: []string{"5"},
				},
			}}
			tr := &rateLimitTransport{inner: inner, waitMax: time.Second, limiter: NewLimiter()}

			req := mustCreateRequest(t, nil, "POST", "http://example.com/repos/test", nil)

			if _, err := tr.RoundTrip(req); err != nil {
				t.Fatalf("unexpected error from first RoundTrip: %v", err)
			}

			if _, err := tr.RoundTrip(req); !errors.Is(err, ErrRateLimitWaitExceeded) {
				t.Fatalf("got %v, want %v", err, ErrRateLimitWaitExceeded)
			}
		})

		t.Run("only_blocks_secondary_category_if_secondary_rate_limit_has_no_quota", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				inner := &testRoundTripper{resp: &http.Response{
					StatusCode: http.StatusTooManyRequests,
					Header: http.Header{
						headerRetryAfter:    []string{"5"},
						headerRateRemaining: []string{"0"},
						headerRateReset:     []string{strconv.FormatInt(time.Now().Add(time.Hour).Unix(), 10)},
					},
				}}
				rl := &limiter{blockedUntil: make(map[string]time.Time)}
				tr := &rateLimitTransport{inner: inner, waitMax: time.Second, limiter: rl}

				req := mustCreateRequest(t, nil, "GET", "http://example.com/repos/test", nil)

				if _, err := tr.RoundTrip(req); err != nil {
					t.Fatalf("unexpected error from RoundTrip: %v", err)
				}

				if len(rl.blockedUntil) != 1 {
					t.Fatalf("blocked categories: got %v, want only secondary", rl.blockedUntil)
				}

				got, ok := rl.blockedUntil[rateLimitCategorySecondary]
				if !ok {
					t.Fatalf("blocked categories: got %v, want it to contain secondary", rl.blockedUntil)
				}

				if want := time.Now().Add(5 * time.Second); !got.Equal(want) {
					t.Errorf("secondary blocked until: got %v, want %v", got, want)
				}
			})
		})

		t.Run("blocks_graphql_on_primary_rate_limit", func(t *testing.T) {
			t.Parallel()

			inner := &testRoundTripper{resp: &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					headerRateRemaining: []string{"0"},
					headerRateReset:     []string{strconv.FormatInt(time.Now().Add(time.Hour).Unix(), 10)},
				},
			}}
			tr := &rateLimitTransport{inner: inner, waitMax: time.Second, limiter: NewLimiter()}

			req := mustCreateRequest(t, nil, "POST", "http://example.com/graphql", nil)

			if _, err := tr.RoundTrip(req); err != nil {
				t.Fatalf("unexpected error from first RoundTrip: %v", err)
			}

			if _, err := tr.RoundTrip(req); !errors.Is(err, ErrRateLimitWaitExceeded) {
				t.Fatalf("got %v, want %v", err, ErrRateLimitWaitExceeded)
			}
		})

		t.Run("looks_up_category_without_base_path", func(t *testing.T) {
			t.Parallel()

			rl := NewLimiter()
			rl.Block("search", time.Now().Add(time.Hour))

			tr := &rateLimitTransport{
				inner:    &testRoundTripper{resp: &http.Response{StatusCode: http.StatusOK}},
				waitMax:  time.Second,
				basePath: "/api/v3",
				limiter:  rl,
			}

			_, err := tr.RoundTrip(mustCreateRequest(t, nil, http.MethodGet, "http://example.com/api/v3/search/issues", nil))
			if !errors.Is(err, ErrRateLimitWaitExceeded) {
				t.Fatalf("got %v, want %v", err, ErrRateLimitWaitExceeded)
			}
		})
	})

	t.Run("trimBasePath", func(t *testing.T) {
		t.Parallel()

		for _, tt := range []struct {
			name     string
			basePath string
			path     string
			want     string
		}{
			{
				name:     "empty_base_path",
				basePath: "",
				path:     "/graphql",
				want:     "/graphql",
			},
			{
				name:     "strips_base_path",
				basePath: "/api/v3",
				path:     "/api/v3/repos/o/r",
				want:     "/repos/o/r",
			},
			{
				name:     "ignores_partial_segment_match",
				basePath: "/api",
				path:     "/apixyz/graphql",
				want:     "/apixyz/graphql",
			},
			{
				name:     "ignores_non_matching_path",
				basePath: "/api/v3",
				path:     "/repos/o/r",
				want:     "/repos/o/r",
			},
		} {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				tr := &rateLimitTransport{basePath: tt.basePath}

				if got := tr.trimBasePath(tt.path); got != tt.want {
					t.Errorf("got %q, want %q", got, tt.want)
				}
			})
		}
	})

	t.Run("waitIfBlocked", func(t *testing.T) {
		t.Parallel()

		for _, tt := range []struct {
			name string
			wait time.Duration
			want time.Duration
		}{
			{
				name: "returns_immediately_if_not_blocked",
			},
			{
				name: "waits_for_until_to_pass",
				wait: 2 * time.Second,
				want: 2 * time.Second,
			},
		} {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				synctest.Test(t, func(t *testing.T) {
					tr := &rateLimitTransport{
						waitMax: 5 * time.Second,
						limiter: &dummyLimiter{waitUntil: time.Now().Add(tt.wait)},
					}

					start := time.Now()
					err := tr.waitIfBlocked(t.Context(), "test")
					elapsed := time.Since(start)

					if err != nil {
						t.Fatalf("unexpected error: %v", err)
					}

					if elapsed != tt.want {
						t.Fatalf("wait: got %v, want %v", elapsed, tt.want)
					}
				})
			})
		}

		t.Run("errors_if_context_is_cancelled", func(t *testing.T) {
			t.Parallel()

			tr := &rateLimitTransport{}

			ctx, cancel := context.WithCancel(t.Context())
			cancel()

			err := tr.waitIfBlocked(ctx, "test")
			if err == nil {
				t.Fatal("got nil, want error")
			}

			if !errors.Is(err, context.Canceled) {
				t.Fatalf("got %v, want %v", err, context.Canceled)
			}
		})

		t.Run("errors_if_wait_exceeds_max", func(t *testing.T) {
			t.Parallel()

			tr := &rateLimitTransport{
				waitMax: 1 * time.Second,
				limiter: &dummyLimiter{waitUntil: time.Now().Add(10 * time.Second)},
			}

			err := tr.waitIfBlocked(t.Context(), "test")
			if err == nil {
				t.Fatal("got nil, want error")
			}

			if !errors.Is(err, ErrRateLimitWaitExceeded) {
				t.Fatalf("got %v, want %v", err, ErrRateLimitWaitExceeded)
			}
		})

		t.Run("errors_with_wait_exceeded_error_for_primary_block", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				until := time.Now().Add(10 * time.Minute)
				rl := NewLimiter()
				rl.Block(rateLimitCategoryCore, until)
				rl.Block(rateLimitCategorySecondary, time.Now().Add(time.Minute))

				tr := &rateLimitTransport{waitMax: 30 * time.Second, limiter: rl}

				err := tr.waitIfBlocked(t.Context(), rateLimitCategoryCore)
				if !errors.Is(err, ErrRateLimitWaitExceeded) {
					t.Fatalf("got %v, want %v", err, ErrRateLimitWaitExceeded)
				}

				waitErr, ok := errors.AsType[*WaitExceededError](err)
				if !ok {
					t.Fatalf("got %T, want *WaitExceededError", err)
				}

				if waitErr.Category != rateLimitCategoryCore {
					t.Errorf("category: got %q, want %q", waitErr.Category, rateLimitCategoryCore)
				}
				if !waitErr.Until.Equal(until) {
					t.Errorf("until: got %v, want %v", waitErr.Until, until)
				}
				if waitErr.MaxWait != 30*time.Second {
					t.Errorf("max wait: got %v, want %v", waitErr.MaxWait, 30*time.Second)
				}

				want := "rate limit wait exceeded: core rate limit is blocked until 2000-01-01T00:10:00Z, which is longer than the maximum wait of 30s"
				if got := err.Error(); got != want {
					t.Errorf("error message: got %q, want %q", got, want)
				}
			})
		})

		t.Run("reports_secondary_category_when_secondary_block_is_later", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				until := time.Now().Add(10 * time.Minute)
				rl := NewLimiter()
				rl.Block(rateLimitCategoryCore, time.Now().Add(time.Minute))
				rl.Block(rateLimitCategorySecondary, until)

				tr := &rateLimitTransport{waitMax: 30 * time.Second, limiter: rl}

				err := tr.waitIfBlocked(t.Context(), rateLimitCategoryCore)
				waitErr, ok := errors.AsType[*WaitExceededError](err)
				if !ok {
					t.Fatalf("got %T, want *WaitExceededError", err)
				}

				if waitErr.Category != rateLimitCategorySecondary {
					t.Errorf("category: got %q, want %q", waitErr.Category, rateLimitCategorySecondary)
				}
				if !waitErr.Until.Equal(until) {
					t.Errorf("until: got %v, want %v", waitErr.Until, until)
				}
			})
		})

		t.Run("reports_secondary_category_when_blocks_end_at_same_time", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				until := time.Now().Add(10 * time.Minute)
				rl := NewLimiter()
				rl.Block(rateLimitCategoryCore, until)
				rl.Block(rateLimitCategorySecondary, until)

				tr := &rateLimitTransport{waitMax: 30 * time.Second, limiter: rl}

				err := tr.waitIfBlocked(t.Context(), rateLimitCategoryCore)
				waitErr, ok := errors.AsType[*WaitExceededError](err)
				if !ok {
					t.Fatalf("got %T, want *WaitExceededError", err)
				}

				if waitErr.Category != rateLimitCategorySecondary {
					t.Errorf("category: got %q, want %q", waitErr.Category, rateLimitCategorySecondary)
				}
				if !waitErr.Until.Equal(until) {
					t.Errorf("until: got %v, want %v", waitErr.Until, until)
				}
			})
		})

		t.Run("reports_context_deadline_as_max_wait_when_earlier_than_wait_max", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				rl := NewLimiter()
				rl.Block(rateLimitCategoryCore, time.Now().Add(10*time.Minute))

				tr := &rateLimitTransport{waitMax: time.Minute, limiter: rl}

				ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
				defer cancel()

				err := tr.waitIfBlocked(ctx, rateLimitCategoryCore)
				waitErr, ok := errors.AsType[*WaitExceededError](err)
				if !ok {
					t.Fatalf("got %T, want *WaitExceededError", err)
				}

				if waitErr.MaxWait != 10*time.Second {
					t.Errorf("max wait: got %v, want %v", waitErr.MaxWait, 10*time.Second)
				}
			})
		})

		t.Run("waits_if_block_clears_before_context_deadline", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				tr := &rateLimitTransport{
					waitMax: time.Minute,
					limiter: &dummyLimiter{waitUntil: time.Now().Add(5 * time.Second)},
				}

				ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
				defer cancel()

				start := time.Now()
				err := tr.waitIfBlocked(ctx, "test")
				elapsed := time.Since(start)

				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}

				if elapsed != 5*time.Second {
					t.Fatalf("wait: got %v, want %v", elapsed, 5*time.Second)
				}
			})
		})

		t.Run("errors_immediately_if_block_outlasts_context_deadline", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				tr := &rateLimitTransport{
					waitMax: time.Minute,
					limiter: &dummyLimiter{waitUntil: time.Now().Add(30 * time.Second)},
				}

				ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
				defer cancel()

				start := time.Now()
				err := tr.waitIfBlocked(ctx, "test")
				elapsed := time.Since(start)

				if !errors.Is(err, ErrRateLimitWaitExceeded) {
					t.Fatalf("got %v, want %v", err, ErrRateLimitWaitExceeded)
				}

				if errors.Is(err, context.DeadlineExceeded) {
					t.Fatalf("got %v, want not %v", err, context.DeadlineExceeded)
				}

				if elapsed != 0 {
					t.Fatalf("wait: got %v, want 0s", elapsed)
				}
			})
		})

		t.Run("errors_if_context_is_canceled_while_waiting", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				tr := &rateLimitTransport{
					waitMax: time.Minute,
					limiter: &dummyLimiter{waitUntil: time.Now().Add(30 * time.Second)},
				}

				ctx, cancel := context.WithCancel(t.Context())
				defer cancel()

				time.AfterFunc(5*time.Second, cancel)

				start := time.Now()
				err := tr.waitIfBlocked(ctx, "test")
				elapsed := time.Since(start)

				if !errors.Is(err, context.Canceled) {
					t.Fatalf("got %v, want %v", err, context.Canceled)
				}

				if elapsed != 5*time.Second {
					t.Fatalf("wait: got %v, want %v", elapsed, 5*time.Second)
				}
			})
		})
	})

	t.Run("normalizeRateLimitHeaders", func(t *testing.T) {
		t.Parallel()

		for _, tt := range []struct {
			name               string
			resp               func(now time.Time) *http.Response
			secondaryWait      time.Duration
			sleep              time.Duration
			wantRetryAfter     *string
			wantRemainingSet   bool
			wantResetSet       bool
			wantSecondaryLimit bool
		}{
			{
				name: "does_not_modify_response_if_not_rate_limited",
				resp: func(time.Time) *http.Response {
					return &http.Response{StatusCode: http.StatusOK, Header: http.Header{}}
				},
			},
			{
				name: "handles_secondary_rate_limit_with_retry_after",
				resp: func(time.Time) *http.Response {
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header:     http.Header{headerRetryAfter: []string{"5"}, headerRateRemaining: []string{"0"}, headerRateReset: []string{strconv.FormatInt(time.Now().Add(30*time.Second).Unix(), 10)}},
					}
				},
				wantRetryAfter:     new("5"),
				wantSecondaryLimit: true,
			},
			{
				name: "handles_secondary_rate_limit_without_retry_after_uses_default_wait",
				resp: func(time.Time) *http.Response {
					return &http.Response{StatusCode: http.StatusTooManyRequests, Header: http.Header{}}
				},
				wantRetryAfter:     new("60"),
				wantSecondaryLimit: true,
			},
			{
				name: "handles_secondary_rate_limit_rounds_up_partial_second_default_wait",
				resp: func(time.Time) *http.Response {
					return &http.Response{StatusCode: http.StatusForbidden, Header: http.Header{headerRetryAfter: []string{"abc"}}}
				},
				secondaryWait:      1500 * time.Millisecond,
				wantRetryAfter:     new("2"),
				wantSecondaryLimit: true,
			},
			{
				name: "handles_primary_rate_limit",
				resp: func(now time.Time) *http.Response {
					return &http.Response{
						StatusCode: http.StatusForbidden,
						Header:     http.Header{headerRateRemaining: []string{"0"}, headerRateReset: []string{strconv.FormatInt(now.Add(5*time.Second).Unix(), 10)}},
					}
				},
				wantRetryAfter:   new("5"),
				wantRemainingSet: true,
				wantResetSet:     true,
			},

			{
				name: "handles_primary_rate_limit_rounds_up_partial_second",
				resp: func(now time.Time) *http.Response {
					return &http.Response{
						StatusCode: http.StatusForbidden,
						Header:     http.Header{headerRateRemaining: []string{"0"}, headerRateReset: []string{strconv.FormatInt(now.Add(5*time.Second).Unix(), 10)}},
					}
				},
				sleep:            500 * time.Millisecond,
				wantRetryAfter:   new("5"),
				wantRemainingSet: true,
				wantResetSet:     true,
			},
			{
				name: "handles_primary_rate_limit_whole_second_remaining",
				resp: func(now time.Time) *http.Response {
					return &http.Response{
						StatusCode: http.StatusForbidden,
						Header:     http.Header{headerRateRemaining: []string{"0"}, headerRateReset: []string{strconv.FormatInt(now.Add(5*time.Second).Unix(), 10)}},
					}
				},
				sleep:            time.Second,
				wantRetryAfter:   new("4"),
				wantRemainingSet: true,
				wantResetSet:     true,
			},
			{
				name: "handles_primary_rate_limit_already_reset_without_blocking",
				resp: func(now time.Time) *http.Response {
					return &http.Response{
						StatusCode: http.StatusForbidden,
						Header:     http.Header{headerRateRemaining: []string{"0"}, headerRateReset: []string{strconv.FormatInt(now.Add(5*time.Second).Unix(), 10)}},
					}
				},
				sleep:            5 * time.Second,
				wantRetryAfter:   new("0"),
				wantRemainingSet: true,
				wantResetSet:     true,
			},
			{
				name: "skips_forbidden_with_no_reset",
				resp: func(time.Time) *http.Response {
					return &http.Response{StatusCode: http.StatusForbidden, Header: http.Header{}}
				},
			},
			{
				name: "skips_primary_rate_limit_with_invalid_reset",
				resp: func(time.Time) *http.Response {
					return &http.Response{
						StatusCode: http.StatusForbidden,
						Header:     http.Header{headerRateRemaining: []string{"0"}, headerRateReset: []string{"abc"}},
					}
				},
				wantRemainingSet: true,
				wantResetSet:     true,
			},
		} {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				synctest.Test(t, func(t *testing.T) {
					rl := &limiter{
						blockedUntil: make(map[string]time.Time),
					}

					resp := tt.resp(time.Now())
					secondaryWait := tt.secondaryWait
					if secondaryWait == 0 {
						secondaryWait = 60 * time.Second
					}

					tr := &rateLimitTransport{defaultSecondaryWait: secondaryWait, waitMax: 60 * time.Second, limiter: rl}

					time.Sleep(tt.sleep)

					tr.normalizeRateLimitHeaders(resp)

					if tt.wantRetryAfter != nil {
						if got := resp.Header.Get(headerRetryAfter); got != *tt.wantRetryAfter {
							t.Errorf("Retry-After header: got %q, want %q", got, *tt.wantRetryAfter)
						}
					}

					if got := resp.Header.Get(headerRateRemaining); (got != "") != tt.wantRemainingSet {
						t.Errorf("X-Ratelimit-Remaining header set: got %v, want %v", got != "", tt.wantRemainingSet)
					}

					if got := resp.Header.Get(headerRateReset); (got != "") != tt.wantResetSet {
						t.Errorf("X-Ratelimit-Reset header set: got %v, want %v", got != "", tt.wantResetSet)
					}

					if tt.wantSecondaryLimit && len(rl.blockedUntil) == 0 {
						t.Error("blocked categories: got empty, want non-empty")
					}

					if !tt.wantSecondaryLimit && len(rl.blockedUntil) != 0 {
						t.Errorf("blocked categories: got %v, want empty", rl.blockedUntil)
					}

					if tt.wantSecondaryLimit {
						if _, ok := rl.blockedUntil[rateLimitCategorySecondary]; !ok || len(rl.blockedUntil) != 1 {
							t.Errorf("blocked categories: got %v, want only secondary", rl.blockedUntil)
						}
					}
				})
			})
		}
	})

	t.Run("blockIfExhausted", func(t *testing.T) {
		t.Parallel()

		for _, tt := range []struct {
			name        string
			reqCategory string
			category    string
			remaining   string
			reset       time.Duration
			resetStr    *string
			wantBlocked []string
		}{
			{
				name:        "blocks_if_exhausted",
				reqCategory: rateLimitCategoryCore,
				category:    rateLimitCategoryCore,
				remaining:   "0",
				reset:       time.Hour,
				wantBlocked: []string{rateLimitCategoryCore},
			},
			{
				name:        "blocks_if_exhausted_without_category",
				reqCategory: rateLimitCategoryGraphQL,
				remaining:   "0",
				reset:       time.Hour,
				wantBlocked: []string{rateLimitCategoryGraphQL},
			},
			{
				name:        "blocks_request_and_response_categories",
				reqCategory: rateLimitCategoryCore,
				category:    rateLimitCategorySearch,
				remaining:   "0",
				reset:       time.Hour,
				wantBlocked: []string{rateLimitCategoryCore, rateLimitCategorySearch},
			},
			{
				name:        "does_not_block_if_quota_remaining",
				reqCategory: rateLimitCategoryCore,
				category:    rateLimitCategoryCore,
				remaining:   "5",
				reset:       time.Hour,
			},
			{
				name:        "does_not_block_if_already_reset",
				reqCategory: rateLimitCategoryCore,
				category:    rateLimitCategoryCore,
				remaining:   "0",
				reset:       -time.Hour,
			},
			{
				name:        "does_not_block_if_reset_invalid",
				reqCategory: rateLimitCategoryCore,
				category:    rateLimitCategoryCore,
				remaining:   "0",
				resetStr:    new("abc"),
			},
		} {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				synctest.Test(t, func(t *testing.T) {
					rl := &limiter{blockedUntil: make(map[string]time.Time)}
					tr := &rateLimitTransport{limiter: rl}

					var reset string
					if tt.resetStr != nil {
						reset = *tt.resetStr
					} else {
						reset = strconv.FormatInt(time.Now().Add(tt.reset).Unix(), 10)
					}

					resp := &http.Response{Header: http.Header{headerRateRemaining: []string{tt.remaining}, headerRateReset: []string{reset}, headerRateResource: []string{tt.category}}}

					tr.blockIfExhausted(resp, tt.reqCategory)

					if len(rl.blockedUntil) != len(tt.wantBlocked) {
						t.Fatalf("blocked categories: got %v, want %v", rl.blockedUntil, tt.wantBlocked)
					}

					want := time.Now().Add(tt.reset)
					for _, category := range tt.wantBlocked {
						blockedUntil, ok := rl.blockedUntil[category]
						if !ok {
							t.Errorf("category %q: got unblocked, want blocked", category)
							continue
						}

						if !blockedUntil.Equal(want) {
							t.Errorf("category %q blocked until: got %v, want %v", category, blockedUntil, want)
						}
					}
				})
			})
		}
	})
}

func Test_ceilSeconds(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		d    time.Duration
		want int64
	}{
		{
			name: "zero",
			d:    0,
			want: 0,
		},
		{
			name: "whole_seconds",
			d:    5 * time.Second,
			want: 5,
		},
		{
			name: "rounds_up_fractional_seconds",
			d:    5*time.Second + time.Millisecond,
			want: 6,
		},
		{
			name: "rounds_up_sub_second",
			d:    time.Nanosecond,
			want: 1,
		},
		{
			name: "rounds_up_negative",
			d:    -1500 * time.Millisecond,
			want: -1,
		},
		{
			name: "max_duration_does_not_overflow",
			d:    time.Duration(1<<63 - 1),
			want: int64(time.Duration(1<<63-1)/time.Second) + 1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := ceilSeconds(tt.d)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
