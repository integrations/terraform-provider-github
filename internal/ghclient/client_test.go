package ghclient

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	ghctbbolt "github.com/bored-engineer/github-conditional-http-transport/bbolt"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"golang.org/x/oauth2"
	"golang.org/x/sync/semaphore"

	"github.com/integrations/terraform-provider-github/v6/internal/httpx"
)

// clientOptionsCmp ignores the limiter and cache store fields, which are asserted separately by type.
var clientOptionsCmp = cmp.Options{
	cmp.AllowUnexported(clientOptions{}),
	cmpopts.IgnoreFields(clientOptions{}, "throttleLimiter", "rateLimiter", "cacheStore"),
	tokenSourceComparer,
}

func TestWithClientURL(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		initial clientOptions
		url     *string
		want    clientOptions
	}{
		{
			name:    "sets_url",
			initial: clientOptions{},
			url:     new("https://ghes.example.com/api/v3/"),
			want:    clientOptions{url: new("https://ghes.example.com/api/v3/")},
		},
		{
			name:    "nil_clears_url",
			initial: clientOptions{url: new("https://ghes.example.com/api/v3/")},
			url:     nil,
			want:    clientOptions{url: nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := tt.initial
			if err := WithClientURL(tt.url)(&got); err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if diff := cmp.Diff(tt.want, got, clientOptionsCmp); diff != "" {
				t.Fatalf("unexpected options (-want +got):\n%s", diff)
			}
		})
	}
}

func TestWithClientUserAgent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		initial   clientOptions
		userAgent *string
		want      clientOptions
	}{
		{
			name:      "sets_user_agent",
			initial:   clientOptions{},
			userAgent: new("test-agent"),
			want:      clientOptions{userAgent: new("test-agent")},
		},
		{
			name:      "nil_clears_user_agent",
			initial:   clientOptions{userAgent: new("test-agent")},
			userAgent: nil,
			want:      clientOptions{userAgent: nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := tt.initial
			if err := WithClientUserAgent(tt.userAgent)(&got); err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if diff := cmp.Diff(tt.want, got, clientOptionsCmp); diff != "" {
				t.Fatalf("unexpected options (-want +got):\n%s", diff)
			}
		})
	}
}

func TestWithClientIdleOptions(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		maxIdleConns    int
		idleConnTimeout time.Duration
		want            clientOptions
	}{
		{
			name:            "valid_values_are_unchanged",
			maxIdleConns:    50,
			idleConnTimeout: time.Minute,
			want:            clientOptions{maxIdleConns: 50, idleConnTimeout: time.Minute},
		},
		{
			name:            "zero_values_are_unchanged",
			maxIdleConns:    0,
			idleConnTimeout: 0,
			want:            clientOptions{maxIdleConns: 0, idleConnTimeout: 0},
		},
		{
			name:            "negative_max_idle_conns_clamped_to_zero",
			maxIdleConns:    -1,
			idleConnTimeout: time.Minute,
			want:            clientOptions{maxIdleConns: 0, idleConnTimeout: time.Minute},
		},
		{
			name:            "negative_idle_conn_timeout_clamped_to_zero",
			maxIdleConns:    50,
			idleConnTimeout: -time.Minute,
			want:            clientOptions{maxIdleConns: 50, idleConnTimeout: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var got clientOptions
			if err := WithClientIdleOptions(tt.maxIdleConns, tt.idleConnTimeout)(&got); err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if diff := cmp.Diff(tt.want, got, clientOptionsCmp); diff != "" {
				t.Fatalf("unexpected options (-want +got):\n%s", diff)
			}
		})
	}
}

func TestWithClientTimeout(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		timeout time.Duration
		want    clientOptions
	}{
		{
			name:    "positive",
			timeout: time.Minute,
			want:    clientOptions{clientTimeout: time.Minute},
		},
		{
			name:    "zero_means_no_timeout",
			timeout: 0,
			want:    clientOptions{clientTimeout: 0},
		},
		{
			name:    "negative_clamped_to_zero",
			timeout: -time.Minute,
			want:    clientOptions{clientTimeout: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var got clientOptions
			if err := WithClientTimeout(tt.timeout)(&got); err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if diff := cmp.Diff(tt.want, got, clientOptionsCmp); diff != "" {
				t.Fatalf("unexpected options (-want +got):\n%s", diff)
			}
		})
	}
}

func TestWithClientThrottleLimit(t *testing.T) {
	t.Parallel()

	t.Run("positive_sets_limiter", func(t *testing.T) {
		t.Parallel()

		var got clientOptions
		if err := WithClientThrottleLimit(5)(&got); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if got.throttleLimiter == nil {
			t.Fatal("throttle limiter: got nil, want non-nil")
		}

		if _, ok := got.throttleLimiter.(*semaphore.Weighted); !ok {
			t.Fatalf("limiter type: got %T, want *semaphore.Weighted", got.throttleLimiter)
		}
	})

	t.Run("non_positive_leaves_limiter_unset", func(t *testing.T) {
		t.Parallel()

		for _, m := range []int64{0, -1} {
			var got clientOptions
			if err := WithClientThrottleLimit(m)(&got); err != nil {
				t.Fatalf("unexpected error for %d: %s", m, err)
			}

			if got.throttleLimiter != nil {
				t.Fatalf("throttle limiter for %d: got %T, want nil", m, got.throttleLimiter)
			}
		}
	})

	t.Run("non_positive_clears_limiter", func(t *testing.T) {
		t.Parallel()

		for _, m := range []int64{0, -1} {
			got := clientOptions{throttleLimiter: httpx.NewThrottleLimiter(5)}
			if err := WithClientThrottleLimit(m)(&got); err != nil {
				t.Fatalf("unexpected error for %d: %s", m, err)
			}

			if got.throttleLimiter != nil {
				t.Fatalf("throttle limiter for %d: got %T, want nil", m, got.throttleLimiter)
			}
		}
	})
}

func TestWithClientCacheOptions(t *testing.T) {
	t.Parallel()

	t.Run("enabled_creates_store", func(t *testing.T) {
		t.Parallel()

		dir := filepath.Join(t.TempDir(), "cache")

		var got clientOptions
		if err := WithClientCacheOptions(true, dir)(&got); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		// The close func is discarded by the option, so the store is closed through the database.
		store, ok := got.cacheStore.(*ghctbbolt.Storage)
		if !ok {
			t.Fatalf("store type: got %T, want *ghctbbolt.Storage", got.cacheStore)
		}
		t.Cleanup(func() { _ = store.DB.Close() })

		want := clientOptions{cacheEnabled: true}
		if diff := cmp.Diff(want, got, clientOptionsCmp); diff != "" {
			t.Fatalf("unexpected options (-want +got):\n%s", diff)
		}

		if _, err := os.Stat(filepath.Join(dir, "cache.db")); err != nil {
			t.Fatalf("stat cache database: got %v, want nil", err)
		}
	})

	t.Run("disabled_does_not_create_store", func(t *testing.T) {
		t.Parallel()

		dir := filepath.Join(t.TempDir(), "cache")

		var got clientOptions
		if err := WithClientCacheOptions(false, dir)(&got); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if got.cacheStore != nil {
			t.Fatalf("store: got %T, want nil", got.cacheStore)
		}

		if diff := cmp.Diff(clientOptions{cacheEnabled: false}, got, clientOptionsCmp); diff != "" {
			t.Fatalf("unexpected options (-want +got):\n%s", diff)
		}

		if _, err := os.Stat(dir); !os.IsNotExist(err) {
			t.Fatalf("stat cache directory: got %v, want not exist", err)
		}
	})

	t.Run("errors_when_store_cannot_be_created", func(t *testing.T) {
		t.Parallel()

		var got clientOptions

		err := WithClientCacheOptions(true, "\x00c")(&got)
		if err == nil {
			t.Fatal("got nil, want error")
		}

		if want := "failed to create cache directory"; !strings.Contains(err.Error(), want) {
			t.Fatalf("got %q, want it to contain %q", err.Error(), want)
		}

		if got.cacheStore != nil {
			t.Fatalf("store: got %T, want nil", got.cacheStore)
		}
	})
}

func TestWithClientRetryOptions(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		maxRetries int
		waitMin    time.Duration
		waitMax    time.Duration
		jitter     time.Duration
		want       clientOptions
	}{
		{
			name:       "values_are_set",
			maxRetries: 3,
			waitMin:    time.Second,
			waitMax:    time.Minute,
			jitter:     100 * time.Millisecond,
			want: clientOptions{
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
			want: clientOptions{
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

			var got clientOptions
			if err := WithClientRetryOptions(tt.maxRetries, tt.waitMin, tt.waitMax, tt.jitter)(&got); err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if diff := cmp.Diff(tt.want, got, clientOptionsCmp); diff != "" {
				t.Fatalf("unexpected options (-want +got):\n%s", diff)
			}
		})
	}
}

func TestWithClientTokenSource(t *testing.T) {
	t.Parallel()

	t.Run("succeeds", func(t *testing.T) {
		t.Parallel()

		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "token"})

		tests := []struct {
			name    string
			initial clientOptions
			want    clientOptions
		}{
			{
				name:    "sets_token_source",
				initial: clientOptions{},
				want:    clientOptions{tokenSource: ts},
			},
			{
				name:    "replaces_token_source",
				initial: clientOptions{tokenSource: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "other"})},
				want:    clientOptions{tokenSource: ts},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got := tt.initial
				if err := WithClientTokenSource(ts)(&got); err != nil {
					t.Fatalf("unexpected error: %s", err)
				}

				if diff := cmp.Diff(tt.want, got, clientOptionsCmp); diff != "" {
					t.Fatalf("unexpected options (-want +got):\n%s", diff)
				}
			})
		}
	})

	t.Run("errors_when_token_source_is_nil", func(t *testing.T) {
		t.Parallel()

		var got clientOptions

		err := WithClientTokenSource(nil)(&got)
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

func TestWithClientToken(t *testing.T) {
	t.Parallel()

	t.Run("succeeds", func(t *testing.T) {
		t.Parallel()

		var got clientOptions
		if err := WithClientToken("token")(&got); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		want := clientOptions{tokenSource: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "token"})}
		if diff := cmp.Diff(want, got, clientOptionsCmp); diff != "" {
			t.Fatalf("unexpected options (-want +got):\n%s", diff)
		}
	})

	t.Run("errors_when_token_is_empty", func(t *testing.T) {
		t.Parallel()

		var got clientOptions

		err := WithClientToken("")(&got)
		if err == nil {
			t.Fatal("got nil, want error")
		}

		if want := "token cannot be empty when token authentication is enabled"; !strings.Contains(err.Error(), want) {
			t.Fatalf("got %q, want it to contain %q", err.Error(), want)
		}

		if !errors.Is(err, ErrInvalidOption) {
			t.Fatalf("got %v, want %v", err, ErrInvalidOption)
		}

		if got.tokenSource != nil {
			t.Fatal("token source: got non-nil, want nil")
		}
	})
}

func Test_defaultClientOptions(t *testing.T) {
	t.Parallel()

	got := defaultClientOptions()

	want := clientOptions{
		url:                   nil,
		userAgent:             nil,
		maxIdleConns:          100,
		idleConnTimeout:       90 * time.Second,
		responseHeaderTimeout: 60 * time.Second,
		clientTimeout:         10 * time.Minute,
		cacheEnabled:          false,
		retryMax:              3,
		retryWaitMin:          time.Second,
		retryWaitMax:          time.Minute,
		retryJitter:           100 * time.Millisecond,
		tokenSource:           nil,
	}
	if diff := cmp.Diff(want, got, clientOptionsCmp); diff != "" {
		t.Fatalf("unexpected defaults (-want +got):\n%s", diff)
	}

	// The ignored fields must be unset by default.
	if got.throttleLimiter != nil {
		t.Fatalf("throttle limiter: got %T, want nil", got.throttleLimiter)
	}

	if got.rateLimiter != nil {
		t.Fatalf("rate limiter: got %T, want nil", got.rateLimiter)
	}

	if got.cacheStore != nil {
		t.Fatalf("cache store: got %T, want nil", got.cacheStore)
	}
}

func Test_normalizeRetryOptions(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		maxRetries     int
		waitMin        time.Duration
		waitMax        time.Duration
		jitter         time.Duration
		wantMaxRetries int
		wantWaitMin    time.Duration
		wantWaitMax    time.Duration
		wantJitter     time.Duration
	}{
		{
			name:           "valid_values_are_unchanged",
			maxRetries:     3,
			waitMin:        time.Second,
			waitMax:        time.Minute,
			jitter:         100 * time.Millisecond,
			wantMaxRetries: 3,
			wantWaitMin:    time.Second,
			wantWaitMax:    time.Minute,
			wantJitter:     100 * time.Millisecond,
		},
		{
			name:           "zero_values_are_unchanged",
			maxRetries:     0,
			waitMin:        0,
			waitMax:        0,
			jitter:         0,
			wantMaxRetries: 0,
			wantWaitMin:    0,
			wantWaitMax:    0,
			wantJitter:     0,
		},
		{
			name:           "equal_waits_are_unchanged",
			maxRetries:     3,
			waitMin:        time.Second,
			waitMax:        time.Second,
			jitter:         0,
			wantMaxRetries: 3,
			wantWaitMin:    time.Second,
			wantWaitMax:    time.Second,
			wantJitter:     0,
		},
		{
			name:           "negative_max_retries_clamped_to_zero",
			maxRetries:     -1,
			waitMin:        time.Second,
			waitMax:        time.Minute,
			jitter:         0,
			wantMaxRetries: 0,
			wantWaitMin:    time.Second,
			wantWaitMax:    time.Minute,
			wantJitter:     0,
		},
		{
			name:           "negative_wait_min_clamped_to_zero",
			maxRetries:     3,
			waitMin:        -time.Second,
			waitMax:        time.Minute,
			jitter:         0,
			wantMaxRetries: 3,
			wantWaitMin:    0,
			wantWaitMax:    time.Minute,
			wantJitter:     0,
		},
		{
			name:           "negative_wait_max_raised_to_wait_min",
			maxRetries:     3,
			waitMin:        time.Second,
			waitMax:        -time.Minute,
			jitter:         0,
			wantMaxRetries: 3,
			wantWaitMin:    time.Second,
			wantWaitMax:    time.Second,
			wantJitter:     0,
		},
		{
			name:           "negative_waits_clamped_to_zero",
			maxRetries:     3,
			waitMin:        -time.Second,
			waitMax:        -time.Minute,
			jitter:         0,
			wantMaxRetries: 3,
			wantWaitMin:    0,
			wantWaitMax:    0,
			wantJitter:     0,
		},
		{
			name:           "wait_max_raised_to_wait_min",
			maxRetries:     3,
			waitMin:        2 * time.Second,
			waitMax:        time.Second,
			jitter:         0,
			wantMaxRetries: 3,
			wantWaitMin:    2 * time.Second,
			wantWaitMax:    2 * time.Second,
			wantJitter:     0,
		},
		{
			name:           "negative_jitter_clamped_to_zero",
			maxRetries:     3,
			waitMin:        time.Second,
			waitMax:        time.Minute,
			jitter:         -time.Second,
			wantMaxRetries: 3,
			wantWaitMin:    time.Second,
			wantWaitMax:    time.Minute,
			wantJitter:     0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotMaxRetries, gotWaitMin, gotWaitMax, gotJitter := normalizeRetryOptions(tt.maxRetries, tt.waitMin, tt.waitMax, tt.jitter)

			if gotMaxRetries != tt.wantMaxRetries {
				t.Errorf("max retries: got %d, want %d", gotMaxRetries, tt.wantMaxRetries)
			}

			if gotWaitMin != tt.wantWaitMin {
				t.Errorf("wait min: got %v, want %v", gotWaitMin, tt.wantWaitMin)
			}

			if gotWaitMax != tt.wantWaitMax {
				t.Errorf("wait max: got %v, want %v", gotWaitMax, tt.wantWaitMax)
			}

			if gotJitter != tt.wantJitter {
				t.Errorf("jitter: got %v, want %v", gotJitter, tt.wantJitter)
			}
		})
	}
}
