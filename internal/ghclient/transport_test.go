package ghclient

import (
	"errors"
	"net/http"
	"os"
	"regexp"
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

			opts := Options{maxIdleConns: 10, idleConnTimeout: 30 * time.Second}
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

			if htr.MaxIdleConns != opts.maxIdleConns {
				t.Fatalf("expected MaxIdleConns to be %d, got %d", opts.maxIdleConns, htr.MaxIdleConns)
			}

			if htr.MaxIdleConnsPerHost != opts.maxIdleConns {
				t.Fatalf("expected MaxIdleConnsPerHost to be %d, got %d", opts.maxIdleConns, htr.MaxIdleConnsPerHost)
			}

			if htr.IdleConnTimeout != opts.idleConnTimeout {
				t.Fatalf("expected IdleConnTimeout to be %v, got %v", opts.idleConnTimeout, htr.IdleConnTimeout)
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
		opts        Options
		wantErr     string
	}{
		{
			name:        "empty_options",
			tokenSource: nil,
			opts:        Options{},
		},
		{
			name:        "with_token",
			tokenSource: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "test-token"}),
			opts:        Options{},
		},
		{
			name:        "with_retry",
			tokenSource: nil,
			opts:        Options{RetryMax: 1, RetryWaitMin: time.Millisecond, RetryWaitMax: time.Millisecond},
		},
		{
			name:        "with_throttler",
			tokenSource: nil,
			opts:        Options{sema: semaphore.NewWeighted(1)},
		},
		{
			name:        "with_cache",
			tokenSource: nil,
			opts:        Options{CachePath: mustMkdirTemp(t, cacheBasePath, "*")},
		},
		{
			name:        "all_options",
			tokenSource: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "test-token"}),
			opts:        Options{RetryMax: 1, RetryWaitMin: time.Millisecond, RetryWaitMax: time.Millisecond, sema: semaphore.NewWeighted(1), CachePath: mustMkdirTemp(t, cacheBasePath, "*")},
		},
		{
			name:        "errors_with_invalid_cache_path",
			tokenSource: nil,
			opts:        Options{CachePath: "\x00c"},
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
}
