package ghclient

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/sync/semaphore"
)

func Test_getHTTPTransport(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name   string
		source http.RoundTripper
	}{
		{
			name:   "sourcehttp_transport",
			source: &http.Transport{},
		},
		{
			name:   "source_non_http_transport",
			source: &testRoundTripper{err: errors.New("not used")},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			opts := ClientOptions{MaxIdleConns: 10, IdleConnTimeout: 30 * time.Second}
			got := getHTTPTransport(tt.source, opts)

			if got == nil {
				t.Fatal("expected http transport to not be nil")
			}

			if got == tt.source {
				t.Fatal("expected http transport to have been cloned")
			}

			if got.ForceAttemptHTTP2 != true {
				t.Fatal("expected ForceAttemptHTTP2 to be true")
			}

			if got.MaxIdleConns != opts.MaxIdleConns {
				t.Fatalf("expected MaxIdleConns to be %d, got %d", opts.MaxIdleConns, got.MaxIdleConns)
			}

			if got.MaxIdleConnsPerHost != opts.MaxIdleConns {
				t.Fatalf("expected MaxIdleConnsPerHost to be %d, got %d", opts.MaxIdleConns, got.MaxIdleConnsPerHost)
			}

			if got.IdleConnTimeout != opts.IdleConnTimeout {
				t.Fatalf("expected IdleConnTimeout to be %v, got %v", opts.IdleConnTimeout, got.IdleConnTimeout)
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
		wantErr     *string
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
			opts:        ClientOptions{Retry: RetryOptions{Max: 1, WaitMin: time.Millisecond, WaitMax: time.Millisecond}},
		},
		{
			name:        "succeeds_with_throttler",
			tokenSource: nil,
			opts:        ClientOptions{Concurrency: ConcurrencyOptions{Max: 1}},
		},
		{
			name:        "succeeds_with_cache",
			tokenSource: nil,
			opts:        ClientOptions{Cache: CacheOptions{Enabled: true, BasePath: mustMkdirTemp(t, cacheBasePath, "*")}},
		},
		{
			name:        "succeeds_with_all_options",
			tokenSource: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "test-token"}),
			opts:        ClientOptions{Retry: RetryOptions{Max: 1, WaitMin: time.Millisecond, WaitMax: time.Millisecond}, Concurrency: ConcurrencyOptions{Max: 1}, Cache: CacheOptions{Enabled: true, BasePath: mustMkdirTemp(t, cacheBasePath, "*")}},
		},
		{
			name:        "errors_with_invalid_cache_path",
			tokenSource: nil,
			opts:        ClientOptions{Cache: CacheOptions{Enabled: true, BasePath: "\x00c"}},
			wantErr:     new("failed to create cache store"),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tr, err := newTransport(tt.tokenSource, tt.opts)
			if err != nil {
				if tt.wantErr == nil {
					t.Fatalf("expected no error, got %v", err)
				}

				if !regexp.MustCompile(regexp.QuoteMeta(*tt.wantErr)).MatchString(err.Error()) {
					t.Fatalf("expected error %q, got %q", *tt.wantErr, err.Error())
				}

				return
			}

			if tt.wantErr != nil {
				t.Fatalf("expected error %q, got nil", *tt.wantErr)
			}

			if tr == nil {
				t.Fatal("expected transport to be non-nil")
			}
		})
	}
}

func Test_transport_cache_varies_on_authorization(t *testing.T) {
	t.Parallel()

	cacheBasePath := mustMkdirTemp(t, "", "*")
	t.Cleanup(func() {
		_ = os.RemoveAll(cacheBasePath)
	})

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

	opts := ClientOptions{Cache: CacheOptions{Enabled: true, BasePath: mustMkdirTemp(t, cacheBasePath, "*")}}
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
}

func Test_transport_retries(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name            string
		nonSeekableBody bool
		retryMax        int
		failures        []fakeHTTPResponse
		wantStatusCode  int
		wantErr         *string
	}{
		{
			name:           "no_failures",
			retryMax:       3,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "no_failures_no_retries",
			retryMax:       0,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "can_error_with_no_retries",
			retryMax:       0,
			failures:       []fakeHTTPResponse{{statusCode: http.StatusInternalServerError}},
			wantStatusCode: http.StatusInternalServerError,
		},
		{
			name:           "retries_until_success",
			retryMax:       3,
			failures:       []fakeHTTPResponse{{statusCode: http.StatusInternalServerError}, {statusCode: http.StatusInternalServerError}},
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "retries_until_failure",
			retryMax:       1,
			failures:       []fakeHTTPResponse{{statusCode: http.StatusInternalServerError}, {statusCode: http.StatusInternalServerError}},
			wantStatusCode: http.StatusInternalServerError,
		},
		{
			name:           "does_not_retry_on_4xx",
			retryMax:       3,
			failures:       []fakeHTTPResponse{{statusCode: http.StatusBadRequest}},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:            "handles_non_seekable_body",
			nonSeekableBody: true,
			retryMax:        3,
			failures:        []fakeHTTPResponse{{statusCode: http.StatusInternalServerError}},
			wantStatusCode:  http.StatusOK,
		},
		{
			name:           "handles_secondary_rate_limit",
			retryMax:       3,
			failures:       []fakeHTTPResponse{{statusCode: http.StatusForbidden, headers: map[string]string{"retry-after": "1"}, body: []byte(`{"message": "You have exceeded a secondary rate limit."}`)}},
			wantStatusCode: http.StatusOK,
		},
		{
			name:            "handles_secondary_rate_limit_with_non_seekable_body",
			nonSeekableBody: true,
			retryMax:        3,
			failures:        []fakeHTTPResponse{{statusCode: http.StatusForbidden, headers: map[string]string{"retry-after": "1"}, body: []byte(`{"message": "You have exceeded a secondary rate limit."}`)}},
			wantStatusCode:  http.StatusOK,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			wantBody := "PASS"

			called := atomic.Int32{}
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				call := int(called.Add(1))

				if call <= len(tt.failures) {
					f := tt.failures[call-1]

					for k, v := range f.headers {
						w.Header().Set(k, v)
					}

					w.WriteHeader(f.statusCode)

					_, _ = w.Write(f.body)

					return
				}

				by, _ := io.ReadAll(r.Body)

				w.WriteHeader(http.StatusOK)
				_, _ = w.Write(by)
			}))
			defer ts.Close()

			opts := ClientOptions{Retry: RetryOptions{Max: tt.retryMax, WaitMin: time.Millisecond, WaitMax: time.Millisecond}}
			tr, err := newTransport(nil, opts)
			if err != nil {
				t.Fatalf("failed to create transport: %v", err)
			}

			if tr == nil {
				t.Fatal("expected transport to be non-nil")
			}

			client := &http.Client{Transport: tr}

			var reqBody io.Reader
			if tt.nonSeekableBody {
				reqBody = &nonSeekableReader{r: strings.NewReader(wantBody)}
			} else {
				reqBody = strings.NewReader(wantBody)
			}

			req, err := http.NewRequestWithContext(t.Context(), "POST", ts.URL, reqBody)
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}

			res, err := client.Do(req)
			if err != nil {
				if tt.wantErr == nil {
					t.Fatalf("expected no error, got %v", err)
				}

				if !regexp.MustCompile(regexp.QuoteMeta(*tt.wantErr)).MatchString(err.Error()) {
					t.Fatalf("expected error %q, got %q", *tt.wantErr, err.Error())
				}

				return
			}
			defer res.Body.Close()

			if tt.wantErr != nil {
				t.Fatalf("expected error %q, got nil", *tt.wantErr)
				return
			}

			if res.StatusCode != tt.wantStatusCode {
				t.Fatalf("expected status code %d, got %d", tt.wantStatusCode, res.StatusCode)
			}

			if res.StatusCode != http.StatusOK {
				return
			}

			by, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("failed to read response body: %v", err)
			}
			body := string(by)

			if body != wantBody {
				t.Fatalf("expected response body to be %q, got %q", wantBody, body)
			}
		})
	}
}

func Test_transport_throttles(t *testing.T) {
	t.Parallel()

	result := "FAIL"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(result))
	}))
	defer ts.Close()

	opts := ClientOptions{Concurrency: ConcurrencyOptions{Max: 1, sema: semaphore.NewWeighted(1)}}
	tr, err := newTransport(nil, opts)
	if err != nil {
		t.Fatalf("failed to create transport: %v", err)
	}

	if tr == nil {
		t.Fatal("expected transport to be non-nil")
	}

	client := &http.Client{Transport: tr}

	if err := opts.Concurrency.sema.Acquire(t.Context(), 1); err != nil {
		t.Fatalf("failed to acquire semaphore: %v", err)
	}

	go func() {
		time.Sleep(1 * time.Second)
		result = "PASS"
		opts.Concurrency.sema.Release(1)
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
}
