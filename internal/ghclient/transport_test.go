package ghclient

import (
	"net/http"
	"testing"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/sync/semaphore"
)

func Test_cloneTransport(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name                string
		original            http.RoundTripper
		expectSamePointer   bool
		expectHTTPTransport bool
	}{
		{
			name:                "http_transport",
			original:            &http.Transport{},
			expectSamePointer:   false,
			expectHTTPTransport: true,
		},
		{
			name:                "non_http_transport",
			original:            &staticRoundTripper{},
			expectSamePointer:   true,
			expectHTTPTransport: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			cloned := cloneTransport(tt.original, Options{maxIdleConns: 10, idleConnTimeout: 30 * time.Second})

			if tt.expectSamePointer && cloned != tt.original {
				t.Fatal("expected cloned transport to match original pointer")
			}

			if !tt.expectSamePointer && cloned == tt.original {
				t.Fatal("expected cloned transport to have a different pointer")
			}

			_, ok := cloned.(*http.Transport)
			if ok != tt.expectHTTPTransport {
				t.Fatalf("unexpected transport type: got %T", cloned)
			}
		})
	}
}

func Test_newTransport(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name            string
		withTokenSource bool
		retryMax        int
		withSemaphore   bool
	}{
		{
			name:            "base_transport",
			withTokenSource: false,
			retryMax:        0,
			withSemaphore:   false,
		},
		{
			name:            "with_retry",
			withTokenSource: false,
			retryMax:        1,
			withSemaphore:   false,
		},
		{
			name:            "with_token_source",
			withTokenSource: true,
			retryMax:        0,
			withSemaphore:   false,
		},
		{
			name:            "with_throttler",
			withTokenSource: false,
			retryMax:        0,
			withSemaphore:   true,
		},
		{
			name:            "with_all_wrappers",
			withTokenSource: true,
			retryMax:        1,
			withSemaphore:   true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			testOpts := testOptions(t)
			testOpts.RetryMax = tt.retryMax
			testOpts.RetryWaitMin = time.Millisecond
			testOpts.RetryWaitMax = time.Millisecond
			if tt.withSemaphore {
				testOpts.sema = semaphore.NewWeighted(1)
			}

			var tokenSource oauth2.TokenSource
			if tt.withTokenSource {
				tokenSource = oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "test-token"})
			}

			tr, err := newTransport(tokenSource, testOpts)
			if err != nil {
				t.Fatalf("failed to create transport: %v", err)
			}

			if tr == nil {
				t.Fatal("expected transport to be non-nil")
			}
		})
	}
}
