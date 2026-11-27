package ghclient

import (
	"crypto"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

var sourceOptionsCmp = cmp.Options{
	cmp.AllowUnexported(sourceOptions{}),
	cmp.Comparer(func(a, b crypto.Signer) bool {
		return a == b
	}),
	tokenSourceComparer,
}

func TestWithSourceBaseURL(t *testing.T) {
	t.Parallel()

	t.Run("succeeds", func(t *testing.T) {
		t.Parallel()

		u := mustParseURL(t, "https://ghes.example.com")

		tests := []struct {
			name    string
			baseURL *url.URL
			isGHES  bool
			want    sourceOptions
		}{
			{
				name:    "ghes_with_url",
				baseURL: u,
				isGHES:  true,
				want:    sourceOptions{baseURL: u, isGHES: true},
			},
			{
				name:    "non_ghes_with_url",
				baseURL: u,
				isGHES:  false,
				want:    sourceOptions{baseURL: u, isGHES: false},
			},
			{
				name:    "non_ghes_nil_url",
				baseURL: nil,
				isGHES:  false,
				want:    sourceOptions{baseURL: nil, isGHES: false},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got sourceOptions
				if err := WithSourceBaseURL(tt.baseURL, tt.isGHES)(&got); err != nil {
					t.Fatalf("unexpected error: %s", err)
				}

				if diff := cmp.Diff(tt.want, got, sourceOptionsCmp); diff != "" {
					t.Fatalf("unexpected options (-want +got):\n%s", diff)
				}
			})
		}
	})

	t.Run("errors", func(t *testing.T) {
		t.Parallel()

		var got sourceOptions

		err := WithSourceBaseURL(nil, true)(&got)
		if err == nil {
			t.Fatal("got nil, want error")
		}

		if want := "base URL cannot be nil when ghes is enabled"; !strings.Contains(err.Error(), want) {
			t.Fatalf("got %q, want it to contain %q", err.Error(), want)
		}

		if !errors.Is(err, ErrInvalidOption) {
			t.Fatalf("got %v, want %v", err, ErrInvalidOption)
		}
	})
}

func TestWithSourceUserAgent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		initial   sourceOptions
		userAgent *string
		want      sourceOptions
	}{
		{
			name:      "sets_user_agent",
			initial:   sourceOptions{},
			userAgent: new("test-agent"),
			want:      sourceOptions{userAgent: new("test-agent")},
		},
		{
			name:      "nil_clears_user_agent",
			initial:   sourceOptions{userAgent: new("test-agent")},
			userAgent: nil,
			want:      sourceOptions{userAgent: nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := tt.initial
			if err := WithSourceUserAgent(tt.userAgent)(&got); err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if diff := cmp.Diff(tt.want, got, sourceOptionsCmp); diff != "" {
				t.Fatalf("unexpected options (-want +got):\n%s", diff)
			}
		})
	}
}

func TestWithSourceThrottleLimit(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		limit int64
		want  sourceOptions
	}{
		{
			name:  "zero",
			limit: 0,
			want:  sourceOptions{throttleLimit: 0},
		},
		{
			name:  "positive",
			limit: 10,
			want:  sourceOptions{throttleLimit: 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var got sourceOptions
			if err := WithSourceThrottleLimit(tt.limit)(&got); err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if diff := cmp.Diff(tt.want, got, sourceOptionsCmp); diff != "" {
				t.Fatalf("unexpected options (-want +got):\n%s", diff)
			}
		})
	}
}

func TestWithSourceCache(t *testing.T) {
	t.Parallel()

	t.Run("succeeds", func(t *testing.T) {
		t.Parallel()

		var got sourceOptions
		if err := WithSourceCache("/tmp/cache")(&got); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		want := sourceOptions{cacheEnabled: true, cacheBasePath: "/tmp/cache"}
		if diff := cmp.Diff(want, got, sourceOptionsCmp); diff != "" {
			t.Fatalf("unexpected options (-want +got):\n%s", diff)
		}
	})

	t.Run("errors_when_base_path_is_empty", func(t *testing.T) {
		t.Parallel()

		var got sourceOptions

		err := WithSourceCache("")(&got)
		if err == nil {
			t.Fatal("got nil, want error")
		}

		if want := "cache base path cannot be empty"; !strings.Contains(err.Error(), want) {
			t.Fatalf("got %q, want it to contain %q", err.Error(), want)
		}

		if !errors.Is(err, ErrInvalidOption) {
			t.Fatalf("got %v, want %v", err, ErrInvalidOption)
		}

		if diff := cmp.Diff(sourceOptions{}, got, sourceOptionsCmp); diff != "" {
			t.Fatalf("unexpected options (-want +got):\n%s", diff)
		}
	})
}

func TestWithSourceRetry(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		maxRetries int
		waitMin    time.Duration
		waitMax    time.Duration
		jitter     time.Duration
		want       sourceOptions
	}{
		{
			name:       "values_are_set",
			maxRetries: 3,
			waitMin:    time.Second,
			waitMax:    time.Minute,
			jitter:     100 * time.Millisecond,
			want: sourceOptions{
				retryMax:     3,
				retryWaitMin: time.Second,
				retryWaitMax: time.Minute,
				retryJitter:  100 * time.Millisecond,
			},
		},
		{
			name:       "values_are_normalized",
			maxRetries: -1,
			waitMin:    2 * time.Second,
			waitMax:    time.Second,
			jitter:     -time.Second,
			want: sourceOptions{
				retryMax:     0,
				retryWaitMin: 2 * time.Second,
				retryWaitMax: 2 * time.Second,
				retryJitter:  0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var got sourceOptions
			if err := WithSourceRetry(tt.maxRetries, tt.waitMin, tt.waitMax, tt.jitter)(&got); err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if diff := cmp.Diff(tt.want, got, sourceOptionsCmp); diff != "" {
				t.Fatalf("unexpected options (-want +got):\n%s", diff)
			}
		})
	}
}

func TestWithSourceAppAuth(t *testing.T) {
	t.Parallel()

	t.Run("succeeds", func(t *testing.T) {
		t.Parallel()

		signer := mustReadAppSigner(t)

		var got sourceOptions
		if err := WithSourceAppAuth("123456789", signer)(&got); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		want := sourceOptions{
			appAuth:     true,
			appClientID: "123456789",
			appSigner:   signer,
		}
		if diff := cmp.Diff(want, got, sourceOptionsCmp); diff != "" {
			t.Fatalf("unexpected options (-want +got):\n%s", diff)
		}
	})

	t.Run("errors", func(t *testing.T) {
		t.Parallel()

		signer := mustReadAppSigner(t)

		tests := []struct {
			name     string
			clientID string
			signer   crypto.Signer
			wantErr  string
		}{
			{
				name:     "empty_client_id",
				clientID: "",
				signer:   signer,
				wantErr:  "client id cannot be empty when app authentication is enabled",
			},
			{
				name:     "nil_signer",
				clientID: "123456789",
				signer:   nil,
				wantErr:  "signer cannot be nil when app authentication is enabled",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var got sourceOptions

				err := WithSourceAppAuth(tt.clientID, tt.signer)(&got)
				if err == nil {
					t.Fatal("got nil, want error")
				}

				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("got %q, want it to contain %q", err.Error(), tt.wantErr)
				}

				if !errors.Is(err, ErrInvalidOption) {
					t.Fatalf("got %v, want %v", err, ErrInvalidOption)
				}

				if diff := cmp.Diff(sourceOptions{}, got, sourceOptionsCmp); diff != "" {
					t.Fatalf("unexpected options (-want +got):\n%s", diff)
				}
			})
		}
	})
}

func TestWithSourceTokenSource(t *testing.T) {
	t.Parallel()

	t.Run("succeeds", func(t *testing.T) {
		t.Parallel()

		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "token"})

		tests := []struct {
			name    string
			initial sourceOptions
			want    sourceOptions
		}{
			{
				name:    "sets_token_source",
				initial: sourceOptions{},
				want:    sourceOptions{appAuth: false, tokenSource: ts},
			},
			{
				name:    "clears_app_auth",
				initial: sourceOptions{appAuth: true},
				want:    sourceOptions{appAuth: false, tokenSource: ts},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got := tt.initial
				if err := WithSourceTokenSource(ts)(&got); err != nil {
					t.Fatalf("unexpected error: %s", err)
				}

				if diff := cmp.Diff(tt.want, got, sourceOptionsCmp); diff != "" {
					t.Fatalf("unexpected options (-want +got):\n%s", diff)
				}
			})
		}
	})

	t.Run("errors_when_token_source_is_nil", func(t *testing.T) {
		t.Parallel()

		var got sourceOptions

		err := WithSourceTokenSource(nil)(&got)
		if err == nil {
			t.Fatal("got nil, want error")
		}

		if want := "token source cannot be nil when token authentication is enabled"; !strings.Contains(err.Error(), want) {
			t.Fatalf("got %q, want it to contain %q", err.Error(), want)
		}

		if !errors.Is(err, ErrInvalidOption) {
			t.Fatalf("got %v, want %v", err, ErrInvalidOption)
		}
	})
}

func TestWithSourceToken(t *testing.T) {
	t.Parallel()

	t.Run("succeeds", func(t *testing.T) {
		t.Parallel()

		want := sourceOptions{
			appAuth:     false,
			tokenSource: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "token"}),
		}

		tests := []struct {
			name    string
			initial sourceOptions
		}{
			{
				name:    "sets_token_source",
				initial: sourceOptions{},
			},
			{
				name:    "clears_app_auth",
				initial: sourceOptions{appAuth: true},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got := tt.initial
				if err := WithSourceToken("token")(&got); err != nil {
					t.Fatalf("unexpected error: %s", err)
				}

				if diff := cmp.Diff(want, got, sourceOptionsCmp); diff != "" {
					t.Fatalf("unexpected options (-want +got):\n%s", diff)
				}
			})
		}
	})

	t.Run("errors_when_token_is_empty", func(t *testing.T) {
		t.Parallel()

		var got sourceOptions

		err := WithSourceToken("")(&got)
		if err == nil {
			t.Fatal("got nil, want error")
		}

		if want := "token cannot be empty when token authentication is enabled"; !strings.Contains(err.Error(), want) {
			t.Fatalf("got %q, want it to contain %q", err.Error(), want)
		}

		if !errors.Is(err, ErrInvalidOption) {
			t.Fatalf("got %v, want %v", err, ErrInvalidOption)
		}
	})
}

func Test_defaultSourceOptions(t *testing.T) {
	t.Parallel()

	want := sourceOptions{
		baseURL:       nil,
		isGHES:        false,
		userAgent:     nil,
		throttleLimit: 0,
		cacheEnabled:  false,
		cacheBasePath: "",
		retryMax:      3,
		retryWaitMin:  time.Second,
		retryWaitMax:  time.Minute,
		retryJitter:   100 * time.Millisecond,
		appAuth:       false,
		appClientID:   "",
		appSigner:     nil,
		tokenSource:   nil,
	}

	if diff := cmp.Diff(want, defaultSourceOptions(), sourceOptionsCmp); diff != "" {
		t.Fatalf("unexpected defaults (-want +got):\n%s", diff)
	}
}

func Test_newSourceClientGroup(t *testing.T) {
	t.Parallel()

	t.Run("returns_clients", func(t *testing.T) {
		t.Parallel()

		got, err := newSourceClientGroup(defaultSourceOptions(), "", oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "token"}))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if got.restClient == nil {
			t.Fatal("rest client: got nil, want non-nil")
		}

		if got.graphqlClient == nil {
			t.Fatal("graphql client: got nil, want non-nil")
		}

		if want := "https://api.github.com/"; got.restClient.BaseURL() != want {
			t.Fatalf("rest base url: got %q, want %q", got.restClient.BaseURL(), want)
		}

		if got.cleanup == nil {
			t.Fatal("cleanup: got nil, want non-nil")
		}
	})

	t.Run("uses_ghes_rest_url", func(t *testing.T) {
		t.Parallel()

		opts := defaultSourceOptions()
		opts.baseURL = mustParseURL(t, "https://ghes.example.com")
		opts.isGHES = true

		got, err := newSourceClientGroup(opts, "", oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "token"}))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if want := "https://ghes.example.com/api/v3/"; got.restClient.BaseURL() != want {
			t.Fatalf("rest base url: got %q, want %q", got.restClient.BaseURL(), want)
		}
	})

	t.Run("cleanup_without_cache_returns_nil", func(t *testing.T) {
		t.Parallel()

		got, err := newSourceClientGroup(defaultSourceOptions(), "", oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "token"}))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if err := got.cleanup(); err != nil {
			t.Fatalf("unexpected cleanup error: %s", err)
		}
	})

	t.Run("creates_cache_database_in_cache_dir", func(t *testing.T) {
		t.Parallel()

		base := t.TempDir()

		opts := defaultSourceOptions()
		opts.cacheEnabled = true
		opts.cacheBasePath = base

		got, err := newSourceClientGroup(opts, "org", oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "token"}))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		t.Cleanup(func() { _ = got.cleanup() })

		if _, err := os.Stat(filepath.Join(base, "org", "cache.db")); err != nil {
			t.Fatalf("stat cache database: got %v, want nil", err)
		}
	})

	t.Run("cleanup_closes_cache_store", func(t *testing.T) {
		t.Parallel()

		base := t.TempDir()

		opts := defaultSourceOptions()
		opts.cacheEnabled = true
		opts.cacheBasePath = base

		got, err := newSourceClientGroup(opts, "org", oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "token"}))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if err := got.cleanup(); err != nil {
			t.Fatalf("unexpected cleanup error: %s", err)
		}

		// Reopening only succeeds if the first store released its file lock.
		_, closeStore, err := createCacheStore(filepath.Join(base, "org"))
		if err != nil {
			t.Fatalf("reopen cache store after cleanup: got %v, want nil", err)
		}
		t.Cleanup(func() { _ = closeStore() })
	})

	t.Run("errors", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name    string
			opts    func() sourceOptions
			wantErr string
		}{
			{
				name: "ghes_without_base_url",
				opts: func() sourceOptions {
					o := defaultSourceOptions()
					o.isGHES = true

					return o
				},
				wantErr: "base URL cannot be nil when ghes is enabled",
			},
			{
				name: "cache_directory_cannot_be_created",
				opts: func() sourceOptions {
					o := defaultSourceOptions()
					o.cacheEnabled = true
					o.cacheBasePath = "\x00c"

					return o
				},
				wantErr: "failed to create cache directory",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := newSourceClientGroup(tt.opts(), "", oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "token"}))
				if err == nil {
					t.Fatal("got nil, want error")
				}

				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("got %q, want it to contain %q", err.Error(), tt.wantErr)
				}

				if got != nil {
					t.Fatal("group: got non-nil, want nil")
				}
			})
		}
	})

	// The remaining subtests need to send requests through the real transport stack, so they use httptest.

	t.Run("routes_requests_through_base_url", func(t *testing.T) {
		t.Parallel()

		type seen struct {
			method    string
			path      string
			userAgent string
			auth      string
		}

		var (
			mu  sync.Mutex
			got []seen
		)

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mu.Lock()
			got = append(got, seen{
				method:    r.Method,
				path:      r.URL.Path,
				userAgent: r.UserAgent(),
				auth:      r.Header.Get("Authorization"),
			})
			mu.Unlock()

			w.Header().Set("Content-Type", "application/json")

			if r.URL.Path == "/api/graphql" {
				_, _ = io.WriteString(w, `{"data":{"viewer":{"login":"octocat"}}}`)

				return
			}

			_, _ = io.WriteString(w, `{}`)
		}))
		t.Cleanup(ts.Close)

		opts := defaultSourceOptions()
		opts.baseURL = mustParseURL(t, ts.URL)
		opts.isGHES = true
		opts.userAgent = new("test-agent")
		opts.retryMax = 0

		g, err := newSourceClientGroup(opts, "", oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "token"}))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		req, err := g.restClient.NewRequest(t.Context(), http.MethodGet, "ping", nil)
		if err != nil {
			t.Fatalf("unexpected error creating rest request: %s", err)
		}

		if _, err := g.restClient.Do(req, nil); err != nil {
			t.Fatalf("unexpected rest error: %s", err)
		}

		var q struct {
			Viewer struct {
				Login githubv4.String
			}
		}
		if err := g.graphqlClient.Query(t.Context(), &q, nil); err != nil {
			t.Fatalf("unexpected graphql error: %s", err)
		}

		mu.Lock()
		defer mu.Unlock()

		if len(got) != 2 {
			t.Fatalf("request count: got %d, want 2", len(got))
		}

		wantREST := seen{method: http.MethodGet, path: "/api/v3/ping", userAgent: "test-agent", auth: "Bearer token"}
		if diff := cmp.Diff(wantREST, got[0], cmp.AllowUnexported(seen{})); diff != "" {
			t.Fatalf("unexpected rest request (-want +got):\n%s", diff)
		}

		// GraphQL clients don't set a user agent, so only the method, path and auth are asserted exactly.
		if got[1].method != http.MethodPost {
			t.Fatalf("graphql method: got %q, want %q", got[1].method, http.MethodPost)
		}

		if want := "/api/graphql"; got[1].path != want {
			t.Fatalf("graphql path: got %q, want %q", got[1].path, want)
		}

		if got[1].userAgent == "test-agent" {
			t.Fatal("graphql user agent: got rest user agent, want different")
		}

		if want := "Bearer token"; got[1].auth != want {
			t.Fatalf("graphql authorization: got %q, want %q", got[1].auth, want)
		}
	})

	t.Run("shares_throttle_between_rest_and_graphql", func(t *testing.T) {
		t.Parallel()

		var inflight, maxInflight atomic.Int32

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			n := inflight.Add(1)
			defer inflight.Add(-1)

			for {
				m := maxInflight.Load()
				if n <= m || maxInflight.CompareAndSwap(m, n) {
					break
				}
			}

			time.Sleep(50 * time.Millisecond)

			w.Header().Set("Content-Type", "application/json")

			if r.URL.Path == "/api/graphql" {
				_, _ = io.WriteString(w, `{"data":{"viewer":{"login":"octocat"}}}`)

				return
			}

			_, _ = io.WriteString(w, `{}`)
		}))
		t.Cleanup(ts.Close)

		opts := defaultSourceOptions()
		opts.baseURL = mustParseURL(t, ts.URL)
		opts.isGHES = true
		opts.throttleLimit = 1
		opts.retryMax = 0

		g, err := newSourceClientGroup(opts, "", oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "token"}))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		var wg sync.WaitGroup

		for range 4 {
			wg.Go(func() {
				req, err := g.restClient.NewRequest(t.Context(), http.MethodGet, "ping", nil)
				if err != nil {
					t.Errorf("unexpected error creating rest request: %s", err)

					return
				}

				if _, err := g.restClient.Do(req, nil); err != nil {
					t.Errorf("unexpected rest error: %s", err)
				}
			})

			wg.Go(func() {
				var q struct {
					Viewer struct {
						Login githubv4.String
					}
				}
				if err := g.graphqlClient.Query(t.Context(), &q, nil); err != nil {
					t.Errorf("unexpected graphql error: %s", err)
				}
			})
		}

		wg.Wait()

		if got := maxInflight.Load(); got != 1 {
			t.Fatalf("max in-flight requests: got %d, want 1", got)
		}
	})

	t.Run("caches_rest_requests_only", func(t *testing.T) {
		t.Parallel()

		var restConditional, graphqlConditional atomic.Int32

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			if r.URL.Path == "/api/graphql" {
				if r.Header.Get("If-None-Match") != "" {
					graphqlConditional.Add(1)
				}

				_, _ = io.WriteString(w, `{"data":{"viewer":{"login":"octocat"}}}`)

				return
			}

			w.Header().Set("ETag", `"v1"`)
			w.Header().Set("Vary", "Authorization")

			if r.Header.Get("If-None-Match") == `"v1"` {
				restConditional.Add(1)
				w.WriteHeader(http.StatusNotModified)

				return
			}

			_, _ = io.WriteString(w, `{}`)
		}))
		t.Cleanup(ts.Close)

		opts := defaultSourceOptions()
		opts.baseURL = mustParseURL(t, ts.URL)
		opts.isGHES = true
		opts.cacheEnabled = true
		opts.cacheBasePath = t.TempDir()
		opts.retryMax = 0

		g, err := newSourceClientGroup(opts, "", oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "token"}))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		t.Cleanup(func() { _ = g.cleanup() })

		for range 2 {
			req, err := g.restClient.NewRequest(t.Context(), http.MethodGet, "ping", nil)
			if err != nil {
				t.Fatalf("unexpected error creating rest request: %s", err)
			}

			resp, err := g.restClient.Do(req, nil)
			if err != nil {
				t.Fatalf("unexpected rest error: %s", err)
			}

			if resp.StatusCode != http.StatusOK {
				t.Fatalf("rest status: got %d, want %d", resp.StatusCode, http.StatusOK)
			}
		}

		for range 2 {
			var q struct {
				Viewer struct {
					Login githubv4.String
				}
			}
			if err := g.graphqlClient.Query(t.Context(), &q, nil); err != nil {
				t.Fatalf("unexpected graphql error: %s", err)
			}
		}

		if got := restConditional.Load(); got != 1 {
			t.Fatalf("conditional rest requests: got %d, want 1", got)
		}

		if got := graphqlConditional.Load(); got != 0 {
			t.Fatalf("conditional graphql requests: got %d, want 0", got)
		}
	})
}

func TestNewSource(t *testing.T) {
	t.Parallel()

	t.Run("returns_source_for_auth_mode", func(t *testing.T) {
		t.Parallel()

		signer := mustReadAppSigner(t)
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "token"})

		tests := []struct {
			name    string
			opts    []SourceOption
			wantApp bool
		}{
			{
				name:    "no_auth",
				opts:    nil,
				wantApp: false,
			},
			{
				name:    "token",
				opts:    []SourceOption{WithSourceToken("token")},
				wantApp: false,
			},
			{
				name:    "token_source",
				opts:    []SourceOption{WithSourceTokenSource(ts)},
				wantApp: false,
			},
			{
				name:    "app_auth",
				opts:    []SourceOption{WithSourceAppAuth("123456789", signer)},
				wantApp: true,
			},
			{
				name: "token_after_app_auth_overrides_app_auth",
				opts: []SourceOption{
					WithSourceAppAuth("123456789", signer),
					WithSourceToken("token"),
				},
				wantApp: false,
			},
			{
				name: "app_auth_after_token_overrides_token",
				opts: []SourceOption{
					WithSourceToken("token"),
					WithSourceAppAuth("123456789", signer),
				},
				wantApp: true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := NewSource(tt.opts...)
				if err != nil {
					t.Fatalf("unexpected error: %s", err)
				}

				switch got.(type) {
				case *appSource:
					if !tt.wantApp {
						t.Fatal("got *appSource, want *tokenSource")
					}
				case *tokenSource:
					if tt.wantApp {
						t.Fatal("got *tokenSource, want *appSource")
					}
				default:
					t.Fatalf("got %T, want *appSource or *tokenSource", got)
				}
			})
		}
	})

	t.Run("errors", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name              string
			opts              []SourceOption
			wantErr           string
			wantInvalidOption bool
		}{
			{
				name:              "option_error_is_wrapped",
				opts:              []SourceOption{WithSourceToken("")},
				wantErr:           "failed to apply source option: invalid option: token cannot be empty when token authentication is enabled",
				wantInvalidOption: true,
			},
			{
				name:              "ghes_without_base_url",
				opts:              []SourceOption{WithSourceBaseURL(nil, true)},
				wantErr:           "failed to apply source option: invalid option: base URL cannot be nil when ghes is enabled",
				wantInvalidOption: true,
			},
			{
				name:    "client_group_error_is_returned",
				opts:    []SourceOption{WithSourceToken("token"), WithSourceCache("\x00c")},
				wantErr: "failed to create cache directory",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := NewSource(tt.opts...)
				if err == nil {
					t.Fatal("got nil, want error")
				}

				if got != nil {
					t.Fatalf("source: got %T, want nil", got)
				}

				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("got %q, want it to contain %q", err.Error(), tt.wantErr)
				}

				if got := errors.Is(err, ErrInvalidOption); got != tt.wantInvalidOption {
					t.Fatalf("errors.Is(%v, ErrInvalidOption): got %v, want %v", err, got, tt.wantInvalidOption)
				}
			})
		}

		// Waits for the cache lock timeout, so takes about a second.
		t.Run("cache_path_shared_with_another_source", func(t *testing.T) {
			t.Parallel()

			dir := t.TempDir()

			first, err := NewSource(WithSourceToken("token"), WithSourceCache(dir))
			if err != nil {
				t.Fatalf("unexpected error creating first source: %s", err)
			}

			got, err := NewSource(WithSourceToken("token"), WithSourceCache(dir))
			runtime.KeepAlive(first)
			if err == nil {
				t.Fatal("got nil, want error")
			}

			if got != nil {
				t.Fatalf("source: got %T, want nil", got)
			}

			if !errors.Is(err, ErrCacheInUse) {
				t.Fatalf("got %v, want %v", err, ErrCacheInUse)
			}
		})
	})
}
