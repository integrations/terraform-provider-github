package ghclient

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
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

			cloned := cloneTransport(tt.original)

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
		name                 string
		opts                 Options
		firstStatusCode      int
		retryStatusCode      int
		expectedStatusCode   int
		expectedRequestCount int32
	}{
		{
			name:                 "retry_disabled",
			opts:                 Options{},
			firstStatusCode:      http.StatusInternalServerError,
			retryStatusCode:      http.StatusOK,
			expectedStatusCode:   http.StatusInternalServerError,
			expectedRequestCount: 1,
		},
		{
			name:                 "retry_enabled",
			opts:                 Options{RetryMax: 1, RetryWaitMin: time.Millisecond, RetryWaitMax: time.Millisecond},
			firstStatusCode:      http.StatusInternalServerError,
			retryStatusCode:      http.StatusOK,
			expectedStatusCode:   http.StatusOK,
			expectedRequestCount: 2,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var requestCount atomic.Int32
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				count := requestCount.Add(1)
				if count == 1 {
					w.WriteHeader(tt.firstStatusCode)
					return
				}

				w.WriteHeader(tt.retryStatusCode)
			}))
			defer ts.Close()

			tr, err := newTransport(nil, tt.opts)
			if err != nil {
				t.Fatalf("failed to create transport: %v", err)
			}

			client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
			resp, err := client.Get(ts.URL)
			if err != nil {
				t.Fatalf("failed to perform request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.expectedStatusCode {
				t.Fatalf("expected status code %d, got %d", tt.expectedStatusCode, resp.StatusCode)
			}

			if requestCount.Load() != tt.expectedRequestCount {
				t.Fatalf("expected %d requests, got %d", tt.expectedRequestCount, requestCount.Load())
			}
		})
	}
}
