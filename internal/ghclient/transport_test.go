package ghclient

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/oauth2"

	"github.com/integrations/terraform-provider-github/v6/internal/ghratelimit"
)

func Test_newTransport(t *testing.T) {
	t.Parallel()

	t.Run("sends_authorization_with_token_source", func(t *testing.T) {
		t.Parallel()

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(r.Header.Get("Authorization")))
		}))
		t.Cleanup(ts.Close)

		opts := defaultClientOptions()
		opts.tokenSource = oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "test-token"})

		client := &http.Client{Transport: newTransport(opts)}

		res, err := client.Do(mustCreateRequest(t, http.MethodGet, ts.URL))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		defer res.Body.Close()

		got, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %s", err)
		}

		if want := "Bearer test-token"; string(got) != want {
			t.Fatalf("authorization header: got %q, want %q", got, want)
		}
	})

	t.Run("omits_authorization_without_token_source", func(t *testing.T) {
		t.Parallel()

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(r.Header.Get("Authorization")))
		}))
		t.Cleanup(ts.Close)

		client := &http.Client{Transport: newTransport(defaultClientOptions())}

		res, err := client.Do(mustCreateRequest(t, http.MethodGet, ts.URL))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		defer res.Body.Close()

		got, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %s", err)
		}

		if len(got) != 0 {
			t.Fatalf("authorization header: got %q, want empty", got)
		}
	})

	t.Run("cache_varies_on_authorization", func(t *testing.T) {
		t.Parallel()

		const etag = `"fixed-etag"`

		called := atomic.Int32{}
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			call := int(called.Add(1))

			w.Header().Set("Vary", "Authorization")
			w.Header().Set("Etag", etag)

			if r.Header.Get("If-None-Match") == etag {
				w.WriteHeader(http.StatusNotModified)

				return
			}

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("resp-" + strconv.Itoa(call)))
		}))
		t.Cleanup(ts.Close)

		store, closeStore, err := createCacheStore(t.TempDir())
		if err != nil {
			t.Fatalf("failed to create cache store: %s", err)
		}
		t.Cleanup(func() { _ = closeStore() })

		opts := defaultClientOptions()
		opts.cacheEnabled = true
		opts.cacheStore = store
		opts.tokenSource = &sequentialTokenSource{tokens: []string{"token-A", "token-A", "token-B"}}

		client := &http.Client{Transport: newTransport(opts)}

		// The cache transport must see the real Authorization header injected by the OAuth2 transport, otherwise it can't tell tokens apart.
		res1, err := client.Do(mustCreateRequest(t, http.MethodGet, ts.URL))
		if err != nil {
			t.Fatalf("unexpected error on first request: %s", err)
		}
		defer res1.Body.Close()

		body1, err := io.ReadAll(res1.Body)
		if err != nil {
			t.Fatalf("failed to read first response body: %s", err)
		}

		if got := res1.Header.Get("X-Cache"); got != "MISS" {
			t.Fatalf("first request X-Cache: got %q, want %q", got, "MISS")
		}

		// Same token, so the cached entry is revalidated and served.
		res2, err := client.Do(mustCreateRequest(t, http.MethodGet, ts.URL))
		if err != nil {
			t.Fatalf("unexpected error on second request: %s", err)
		}
		defer res2.Body.Close()

		body2, err := io.ReadAll(res2.Body)
		if err != nil {
			t.Fatalf("failed to read second response body: %s", err)
		}

		if got := res2.Header.Get("X-Cache"); got != "HIT" {
			t.Fatalf("second request X-Cache: got %q, want %q", got, "HIT")
		}

		if string(body2) != string(body1) {
			t.Fatalf("second request body: got %q, want %q", body2, body1)
		}

		// Different token, so the first token's entry must not be reused.
		res3, err := client.Do(mustCreateRequest(t, http.MethodGet, ts.URL))
		if err != nil {
			t.Fatalf("unexpected error on third request: %s", err)
		}
		defer res3.Body.Close()

		body3, err := io.ReadAll(res3.Body)
		if err != nil {
			t.Fatalf("failed to read third response body: %s", err)
		}

		if got := res3.Header.Get("X-Cache"); got != "MISS" {
			t.Fatalf("third request X-Cache: got %q, want %q", got, "MISS")
		}

		if string(body3) == string(body1) {
			t.Fatalf("third request body: got cached body %q, want fresh body", body3)
		}
	})

	t.Run("returns_response_when_cache_write_fails", func(t *testing.T) {
		t.Parallel()

		called := atomic.Int32{}
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			called.Add(1)

			w.Header().Set("Etag", `"etag"`)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("PASS"))
		}))
		t.Cleanup(ts.Close)

		store := &testStorage{putErr: errors.New("put failed")}

		opts := defaultClientOptions()
		opts.cacheEnabled = true
		opts.cacheStore = store

		client := &http.Client{Transport: newTransport(opts)}

		res, err := client.Do(mustCreateRequest(t, http.MethodGet, ts.URL))
		if err != nil {
			t.Fatalf("got %v, want nil", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Fatalf("status code: got %d, want %d", res.StatusCode, http.StatusOK)
		}

		got, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %s", err)
		}

		if want := "PASS"; string(got) != want {
			t.Fatalf("response body: got %q, want %q", got, want)
		}

		// A failed cache write must not fail the request, so it must not be retried either.
		if got := called.Load(); got != 1 {
			t.Fatalf("calls: got %d, want %d", got, 1)
		}

		if diff := cmp.Diff([]string{"PASS"}, store.putBodies); diff != "" {
			t.Fatalf("put bodies (-want +got):\n%s", diff)
		}
	})

	t.Run("treats_cache_read_error_as_miss", func(t *testing.T) {
		t.Parallel()

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("PASS"))
		}))
		t.Cleanup(ts.Close)

		opts := defaultClientOptions()
		opts.cacheEnabled = true
		opts.cacheStore = &testStorage{getErr: errors.New("get failed")}

		client := &http.Client{Transport: newTransport(opts)}

		res, err := client.Do(mustCreateRequest(t, http.MethodGet, ts.URL))
		if err != nil {
			t.Fatalf("got %v, want nil", err)
		}
		defer res.Body.Close()

		if got := res.Header.Get("X-Cache"); got != "MISS" {
			t.Fatalf("X-Cache: got %q, want %q", got, "MISS")
		}

		got, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %s", err)
		}

		if want := "PASS"; string(got) != want {
			t.Fatalf("response body: got %q, want %q", got, want)
		}
	})

	t.Run("cache_enabled_without_store", func(t *testing.T) {
		t.Parallel()

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Etag", `"etag"`)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("PASS"))
		}))
		t.Cleanup(ts.Close)

		opts := defaultClientOptions()
		opts.cacheEnabled = true

		client := &http.Client{Transport: newTransport(opts)}

		res, err := client.Do(mustCreateRequest(t, http.MethodGet, ts.URL))
		if err != nil {
			t.Fatalf("got %v, want nil", err)
		}
		defer res.Body.Close()

		if got := res.Header.Get("X-Cache"); got != "MISS" {
			t.Fatalf("X-Cache: got %q, want %q", got, "MISS")
		}

		got, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %s", err)
		}

		if want := "PASS"; string(got) != want {
			t.Fatalf("response body: got %q, want %q", got, want)
		}
	})

	t.Run("rewinds_body_on_retry", func(t *testing.T) {
		t.Parallel()

		called := atomic.Int32{}
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if called.Add(1) == 1 {
				w.Header().Set("Retry-After", "1")
				w.WriteHeader(http.StatusForbidden)
				_, _ = w.Write([]byte(`{"message": "You have exceeded a secondary rate limit."}`))

				return
			}

			body, _ := io.ReadAll(r.Body)

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(body)
		}))
		t.Cleanup(ts.Close)

		opts := defaultClientOptions()
		opts.retryMax = 1
		opts.retryWaitMin = time.Millisecond
		opts.retryWaitMax = 2 * time.Second
		opts.retryJitter = 0

		client := &http.Client{Transport: newTransport(opts)}

		// A non-seekable body can only be replayed if the rewindable transport buffers it.
		req, err := http.NewRequestWithContext(t.Context(), http.MethodPost, ts.URL, &nonSeekableReader{r: strings.NewReader("PASS")})
		if err != nil {
			t.Fatalf("failed to create request: %s", err)
		}

		res, err := client.Do(req)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Fatalf("status code: got %d, want %d", res.StatusCode, http.StatusOK)
		}

		got, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %s", err)
		}

		if want := "PASS"; string(got) != want {
			t.Fatalf("response body: got %q, want %q", got, want)
		}

		if got := called.Load(); got != 2 {
			t.Fatalf("calls: got %d, want %d", got, 2)
		}
	})

	t.Run("retries_get_after_response_header_timeout", func(t *testing.T) {
		t.Parallel()

		release := make(chan struct{})
		called := atomic.Int32{}
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if called.Add(1) == 1 {
				select {
				case <-release:
				case <-r.Context().Done():
				}

				return
			}

			w.WriteHeader(http.StatusOK)
		}))
		t.Cleanup(ts.Close)
		// Cleanups run in reverse order, so the handler is released before the server is closed.
		t.Cleanup(func() { close(release) })

		opts := defaultClientOptions()
		opts.responseHeaderTimeout = 50 * time.Millisecond
		opts.retryMax = 1
		opts.retryWaitMin = time.Millisecond
		opts.retryWaitMax = time.Millisecond
		opts.retryJitter = 0

		client := &http.Client{Transport: newTransport(opts)}

		res, err := client.Do(mustCreateRequest(t, http.MethodGet, ts.URL))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Fatalf("status code: got %d, want %d", res.StatusCode, http.StatusOK)
		}

		if got := called.Load(); got != 2 {
			t.Fatalf("calls: got %d, want %d", got, 2)
		}
	})

	t.Run("does_not_retry_post_after_response_header_timeout", func(t *testing.T) {
		t.Parallel()

		release := make(chan struct{})
		called := atomic.Int32{}
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			called.Add(1)

			select {
			case <-release:
			case <-r.Context().Done():
			}
		}))
		t.Cleanup(ts.Close)
		// Cleanups run in reverse order, so the handler is released before the server is closed.
		t.Cleanup(func() { close(release) })

		opts := defaultClientOptions()
		opts.responseHeaderTimeout = 50 * time.Millisecond
		opts.retryMax = 1
		opts.retryWaitMin = time.Millisecond
		opts.retryWaitMax = time.Millisecond
		opts.retryJitter = 0

		client := &http.Client{Transport: newTransport(opts)}

		res, err := client.Do(mustCreateRequest(t, http.MethodPost, ts.URL))
		if err == nil {
			res.Body.Close()
			t.Fatal("got nil, want error")
		}

		if !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("got %v, want %v", err, context.DeadlineExceeded)
		}

		if got := called.Load(); got != 1 {
			t.Fatalf("calls: got %d, want %d", got, 1)
		}
	})

	t.Run("does_not_retry_when_rate_limit_wait_exceeded", func(t *testing.T) {
		t.Parallel()

		synctest.Test(t, func(t *testing.T) {
			// The request is blocked before it is sent, so no server is needed. A retry would sleep for the backoff, which the fake clock makes visible.
			rl := ghratelimit.NewLimiter()
			rl.Block("core", time.Now().Add(time.Hour))

			opts := defaultClientOptions()
			opts.rateLimiter = rl
			opts.retryMax = 3
			opts.retryWaitMin = time.Minute
			opts.retryWaitMax = time.Minute
			opts.retryJitter = 0

			client := &http.Client{Transport: newTransport(opts)}

			start := time.Now()

			res, err := client.Do(mustCreateRequest(t, http.MethodGet, "https://api.github.com/repos/o/r"))
			if err == nil {
				res.Body.Close()
				t.Fatal("got nil, want error")
			}

			if !errors.Is(err, ghratelimit.ErrRateLimitWaitExceeded) {
				t.Fatalf("got %v, want %v", err, ghratelimit.ErrRateLimitWaitExceeded)
			}

			if elapsed := time.Since(start); elapsed != 0 {
				t.Fatalf("retry backoff: got %v, want 0", elapsed)
			}
		})
	})
}

func Test_rateLimitBasePath(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		u    *string
		want string
	}{
		{
			name: "nil_url",
			u:    nil,
			want: "",
		},
		{
			name: "empty_url",
			u:    new(""),
			want: "",
		},
		{
			name: "unparseable_url",
			u:    new("://bad"),
			want: "",
		},
		{
			name: "dotcom_rest",
			u:    new("https://api.github.com/"),
			want: "",
		},
		{
			name: "dotcom_graphql",
			u:    new("https://api.github.com/graphql"),
			want: "",
		},
		{
			name: "ghec_rest",
			u:    new("https://api.octocorp.ghe.com/"),
			want: "",
		},
		{
			name: "ghec_graphql",
			u:    new("https://api.octocorp.ghe.com/graphql"),
			want: "",
		},
		{
			name: "ghes_rest",
			u:    new("https://ghes.example.com/api/v3/"),
			want: "/api/v3",
		},
		{
			name: "ghes_graphql",
			u:    new("https://ghes.example.com/api/graphql"),
			want: "/api",
		},
		{
			name: "ghes_rest_with_prefix",
			u:    new("https://example.com/proxy/api/v3/"),
			want: "/proxy/api/v3",
		},
		{
			name: "graphql_trailing_slash",
			u:    new("https://api.github.com/graphql/"),
			want: "",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := rateLimitBasePath(tt.u)
			if got != tt.want {
				t.Fatalf("base path: got %q, want %q", got, tt.want)
			}
		})
	}
}
