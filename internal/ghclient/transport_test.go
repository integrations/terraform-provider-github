package ghclient

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strconv"
	"sync/atomic"
	"testing"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/sync/semaphore"
)

func Test_cloneTransport(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name          string
		source        http.RoundTripper
		httpTransport bool
	}{
		{
			name:          "http_transport",
			source:        &http.Transport{},
			httpTransport: true,
		},
		{
			name:          "non_http_transport",
			source:        &testRoundTripper{err: errors.New("not used")},
			httpTransport: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			opts := ClientOptions{MaxIdleConns: 10, IdleConnTimeout: 30 * time.Second}
			cloned := cloneTransport(tt.source, opts)

			if !tt.httpTransport && cloned != tt.source {
				t.Fatal("expected cloned transport to match original pointer")
			}

			if tt.httpTransport && cloned == tt.source {
				t.Fatal("expected cloned transport to have a different pointer")
			}

			htr, ok := cloned.(*http.Transport)

			if !tt.httpTransport && ok {
				t.Fatalf("expected cloned transport to not be an *http.Transport")
			}

			if tt.httpTransport && !ok {
				t.Fatalf("expected cloned transport to be an *http.Transport, got %T", cloned)
			}

			if !tt.httpTransport {
				return
			}

			if htr.ForceAttemptHTTP2 != true {
				t.Fatal("expected ForceAttemptHTTP2 to be true")
			}

			if htr.MaxIdleConns != opts.MaxIdleConns {
				t.Fatalf("expected MaxIdleConns to be %d, got %d", opts.MaxIdleConns, htr.MaxIdleConns)
			}

			if htr.MaxIdleConnsPerHost != opts.MaxIdleConns {
				t.Fatalf("expected MaxIdleConnsPerHost to be %d, got %d", opts.MaxIdleConns, htr.MaxIdleConnsPerHost)
			}

			if htr.IdleConnTimeout != opts.IdleConnTimeout {
				t.Fatalf("expected IdleConnTimeout to be %v, got %v", opts.IdleConnTimeout, htr.IdleConnTimeout)
			}
		})
	}
}

func Test_newTransport(t *testing.T) {
	t.Parallel()

	cacheBasePath := mustMkdirTemp(t, "", "*")
	t.Cleanup(func() {
		_ = os.RemoveAll(cacheBasePath)
	})

	for _, tt := range []struct {
		name        string
		tokenSource oauth2.TokenSource
		opts        ClientOptions
		wantErr     string
	}{
		{
			name:        "succeeds_with_empty_options",
			tokenSource: nil,
			opts:        ClientOptions{},
		},
		{
			name:        "succeeds_with_token",
			tokenSource: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "test-token"}),
			opts:        ClientOptions{},
		},
		{
			name:        "succeeds_with_retry",
			tokenSource: nil,
			opts:        ClientOptions{RetryMax: 1, RetryWaitMin: time.Millisecond, RetryWaitMax: time.Millisecond},
		},
		{
			name:        "succeeds_with_throttler",
			tokenSource: nil,
			opts:        ClientOptions{Sema: semaphore.NewWeighted(1)},
		},
		{
			name:        "succeeds_with_cache",
			tokenSource: nil,
			opts:        ClientOptions{Cache: true, CachePath: mustMkdirTemp(t, cacheBasePath, "*")},
		},
		{
			name:        "succeeds_with_all_options",
			tokenSource: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "test-token"}),
			opts:        ClientOptions{RetryMax: 1, RetryWaitMin: time.Millisecond, RetryWaitMax: time.Millisecond, Sema: semaphore.NewWeighted(1), Cache: true, CachePath: mustMkdirTemp(t, cacheBasePath, "*")},
		},
		{
			name:        "errors_with_invalid_cache_path",
			tokenSource: nil,
			opts:        ClientOptions{Cache: true, CachePath: "\x00c"},
			wantErr:     "failed to create cache store",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tr, err := newTransport(tt.tokenSource, tt.opts)
			if err != nil {
				if tt.wantErr == "" {
					t.Fatalf("failed to create transport: %v", err)
				}

				if !regexp.MustCompile(regexp.QuoteMeta(tt.wantErr)).MatchString(err.Error()) {
					t.Fatalf("expected error to match %q, got %v", tt.wantErr, err)
				}

				return
			}

			if tt.wantErr != "" {
				t.Fatalf("expected error %q, got nil", tt.wantErr)
			}

			if tr == nil {
				t.Fatal("expected transport to be non-nil")
			}
		})
	}

	t.Run("transport_cache_varies_on_authorization", func(t *testing.T) {
		t.Parallel()

		// The cache transport must observe the real per-request Authorization header (injected
		// by the OAuth2 transport) rather than an empty one. Otherwise its cache-key validation
		// can never distinguish between requests using different tokens, and repeated requests
		// using the *same* token will never be served from cache either, since the stored Vary
		// metadata for Authorization is never populated correctly.
		const etag = `"fixed-etag"`

		called := atomic.Int32{}
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			call := int(called.Add(1))

			// A real GitHub response varies on the Authorization header.
			w.Header().Set("Vary", "Authorization")

			if r.Header.Get("If-None-Match") == etag {
				w.Header().Set("Etag", etag)
				w.WriteHeader(http.StatusNotModified)
				return
			}

			w.Header().Set("Etag", etag)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("resp-" + strconv.Itoa(call)))
		}))
		defer ts.Close()

		opts := ClientOptions{Cache: true, CachePath: mustMkdirTemp(t, cacheBasePath, "*")}
		tr, err := newTransport(&sequentialTokenSource{tokens: []string{"token-A", "token-A", "token-B"}}, opts)
		if err != nil {
			t.Fatalf("failed to create transport: %v", err)
		}

		client := &http.Client{Transport: tr}

		// Request 1: token-A, nothing cached yet -> MISS.
		res1, err := client.Get(ts.URL)
		if err != nil {
			t.Fatalf("failed to make first request: %v", err)
		}
		b1, err := io.ReadAll(res1.Body)
		if err != nil {
			t.Fatalf("failed to read first response body: %v", err)
		}
		res1.Body.Close()
		if xCache := res1.Header.Get("X-Cache"); xCache != "MISS" {
			t.Fatalf("expected first request to be a MISS, got %q", xCache)
		}

		// Request 2: same token-A as request 1 -> the cached entry's stored Authorization Vary
		// metadata must match, so this must be a HIT with an identical cached body.
		res2, err := client.Get(ts.URL)
		if err != nil {
			t.Fatalf("failed to make second request: %v", err)
		}
		b2, err := io.ReadAll(res2.Body)
		if err != nil {
			t.Fatalf("failed to read second response body: %v", err)
		}
		res2.Body.Close()
		if xCache := res2.Header.Get("X-Cache"); xCache != "HIT" {
			t.Fatalf("expected second request (same token) to be a HIT, got %q", xCache)
		}
		if string(b2) != string(b1) {
			t.Fatalf("expected second request (same token) to reuse the first cached response, got %q and %q", string(b2), string(b1))
		}

		// Request 3: token-B, different token -> must not reuse token-A's cached response, so
		// this must be a MISS with fresh content from the origin.
		res3, err := client.Get(ts.URL)
		if err != nil {
			t.Fatalf("failed to make third request: %v", err)
		}
		b3, err := io.ReadAll(res3.Body)
		if err != nil {
			t.Fatalf("failed to read third response body: %v", err)
		}
		res3.Body.Close()
		if xCache := res3.Header.Get("X-Cache"); xCache != "MISS" {
			t.Fatalf("expected third request (different token) to be a MISS, got %q", xCache)
		}
		if string(b3) == string(b1) {
			t.Fatalf("expected third request (different token) to fetch fresh content, got cached body %q", string(b3))
		}
	})

	t.Run("transport_retries_requests", func(t *testing.T) {
		t.Parallel()

		for _, tt := range []struct {
			name               string
			retryMax           int
			failures           int
			failStatusCode     int
			wantFailStatusCode bool
			wantError          string
		}{
			{
				name:               "no_retries",
				retryMax:           0,
				failures:           1,
				failStatusCode:     http.StatusInternalServerError,
				wantFailStatusCode: true,
			},
			{
				name:           "retries_until_success",
				retryMax:       3,
				failures:       2,
				failStatusCode: http.StatusInternalServerError,
			},
			{
				name:           "retries_until_failure",
				retryMax:       2,
				failures:       3,
				failStatusCode: http.StatusInternalServerError,
				wantError:      "giving up after",
			},
			{
				name:               "does_not_retry_on_4xx",
				retryMax:           3,
				failures:           1,
				failStatusCode:     http.StatusBadRequest,
				wantFailStatusCode: true,
			},
		} {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				called := atomic.Int32{}
				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					call := int(called.Add(1))

					if call <= tt.failures {
						w.WriteHeader(tt.failStatusCode)
						return
					}

					w.WriteHeader(http.StatusOK)
					_, _ = w.Write([]byte("PASS"))
				}))
				defer ts.Close()

				opts := ClientOptions{RetryMax: tt.retryMax, RetryWaitMin: time.Millisecond, RetryWaitMax: time.Millisecond}
				tr, err := newTransport(nil, opts)
				if err != nil {
					t.Fatalf("failed to create transport: %v", err)
				}

				if tr == nil {
					t.Fatal("expected transport to be non-nil")
				}

				client := &http.Client{Transport: tr}

				res, err := client.Get(ts.URL)
				if err != nil {
					if tt.wantError != "" {
						if !regexp.MustCompile(regexp.QuoteMeta(tt.wantError)).MatchString(err.Error()) {
							t.Fatalf("expected error to match %q, got %v", tt.wantError, err)
						}

						return
					}

					t.Fatalf("failed to make request: %v", err)
				}
				defer res.Body.Close()

				if tt.wantError != "" {
					t.Fatalf("expected error %q, got nil", tt.wantError)
				}

				if tt.wantFailStatusCode {
					if res.StatusCode != tt.failStatusCode {
						t.Fatalf("expected status code %d, got %d", tt.failStatusCode, res.StatusCode)
					}

					return
				}

				if res.StatusCode != http.StatusOK {
					t.Fatalf("expected status code %d, got %d", http.StatusOK, res.StatusCode)
				}

				body, err := io.ReadAll(res.Body)
				if err != nil {
					t.Fatalf("failed to read response body: %v", err)
				}

				if string(body) != "PASS" {
					t.Fatalf("expected response body to be %q, got %q", "PASS", string(body))
				}
			})
		}
	})

	t.Run("transport_throttles_requests", func(t *testing.T) {
		t.Parallel()

		result := "FAIL"

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(result))
		}))
		defer ts.Close()

		opts := ClientOptions{Sema: semaphore.NewWeighted(1)}
		tr, err := newTransport(nil, opts)
		if err != nil {
			t.Fatalf("failed to create transport: %v", err)
		}

		if tr == nil {
			t.Fatal("expected transport to be non-nil")
		}

		client := &http.Client{Transport: tr}

		if err := opts.Sema.Acquire(t.Context(), 1); err != nil {
			t.Fatalf("failed to acquire semaphore: %v", err)
		}

		go func() {
			time.Sleep(1 * time.Second)
			result = "PASS"
			opts.Sema.Release(1)
		}()

		res, err := client.Get(ts.URL)
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %v", err)
		}

		if string(body) != "PASS" {
			t.Fatalf("expected response body to be %q, got %q", "PASS", string(body))
		}
	})
}
