package github

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func TestAuthenticatedHTTPClientAppAuthRefresh(t *testing.T) {
	t.Parallel()

	appID := testGitHubAppID
	installationID := testGitHubAppInstallationID

	t.Run("reuses a valid token", func(t *testing.T) {
		t.Parallel()

		var tokenRequests atomic.Int32
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, fmt.Sprintf("/app/installations/%s/access_tokens", installationID)) {
				tokenRequests.Add(1)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				_, _ = io.WriteString(w, fmt.Sprintf(`{"token":"valid-token","expires_at":%q}`, time.Now().Add(time.Hour).UTC().Format(time.RFC3339)))
				return
			}
			if r.Method == http.MethodGet && r.URL.Path == "/test" {
				w.WriteHeader(http.StatusOK)
				return
			}
			http.NotFound(w, r)
		}))
		t.Cleanup(ts.Close)

		serverURL, err := url.Parse(ts.URL + "/")
		if err != nil {
			t.Fatalf("failed to parse test server URL: %v", err)
		}

		cfg := &Config{
			AppID:             &appID,
			AppInstallationID: &installationID,
			AppPEM:            testGitHubAppPrivateKeyPemData,
			BaseURL:           serverURL,
		}

		client, err := cfg.AuthenticatedHTTPClient()
		if err != nil {
			t.Fatalf("failed to create authenticated HTTP client: %v", err)
		}

		for i := range 2 {
			resp, err := client.Get(ts.URL + "/test")
			if err != nil {
				t.Fatalf("request %d failed: %v", i+1, err)
			}
			_ = resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				t.Fatalf("request %d: expected status 200, got %d", i+1, resp.StatusCode)
			}
		}

		if got := tokenRequests.Load(); got != 1 {
			t.Fatalf("expected 1 installation token request, got %d", got)
		}
	})

	t.Run("refreshes an expired token", func(t *testing.T) {
		t.Parallel()

		var tokenRequests atomic.Int32
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, fmt.Sprintf("/app/installations/%s/access_tokens", installationID)) {
				count := tokenRequests.Add(1)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				var expiresAt time.Time
				var token string
				if count == 1 {
					expiresAt = time.Now().Add(-time.Hour)
					token = "expired-token"
				} else {
					expiresAt = time.Now().Add(time.Hour)
					token = "refreshed-token"
				}
				_, _ = io.WriteString(w, fmt.Sprintf(`{"token":%q,"expires_at":%q}`, token, expiresAt.UTC().Format(time.RFC3339)))
				return
			}
			if r.Method == http.MethodGet && r.URL.Path == "/test" {
				w.WriteHeader(http.StatusOK)
				return
			}
			http.NotFound(w, r)
		}))
		t.Cleanup(ts.Close)

		serverURL, err := url.Parse(ts.URL + "/")
		if err != nil {
			t.Fatalf("failed to parse test server URL: %v", err)
		}

		cfg := &Config{
			AppID:             &appID,
			AppInstallationID: &installationID,
			AppPEM:            testGitHubAppPrivateKeyPemData,
			BaseURL:           serverURL,
		}

		client, err := cfg.AuthenticatedHTTPClient()
		if err != nil {
			t.Fatalf("failed to create authenticated HTTP client: %v", err)
		}

		for i := range 2 {
			resp, err := client.Get(ts.URL + "/test")
			if err != nil {
				t.Fatalf("request %d failed: %v", i+1, err)
			}
			_ = resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				t.Fatalf("request %d: expected status 200, got %d", i+1, resp.StatusCode)
			}
		}

		if got := tokenRequests.Load(); got != 2 {
			t.Fatalf("expected 2 installation token requests, got %d", got)
		}
	})
}

func Test_getBaseURL(t *testing.T) {
	testCases := []struct {
		name        string
		url         string
		expectedURL string
		isGHES      bool
		errors      bool
	}{
		{
			name:        "dotcom",
			url:         "https://api.github.com/",
			expectedURL: "https://api.github.com/",
			isGHES:      false,
			errors:      false,
		},
		{
			name:        "dotcom no trailing slash",
			url:         "https://api.github.com",
			expectedURL: "https://api.github.com/",
			isGHES:      false,
			errors:      false,
		},
		{
			name:        "dotcom ui",
			url:         "https://github.com/",
			expectedURL: "https://api.github.com/",
			isGHES:      false,
			errors:      false,
		},
		{
			name:        "dotcom http errors",
			url:         "http://api.github.com/",
			expectedURL: "",
			isGHES:      false,
			errors:      true,
		},
		{
			name:        "dotcom with path errors",
			url:         "https://api.github.com/xxx/",
			expectedURL: "",
			isGHES:      false,
			errors:      true,
		},
		{
			name:        "ghec",
			url:         "https://api.customer.ghe.com/",
			expectedURL: "https://api.customer.ghe.com/",
			isGHES:      false,
			errors:      false,
		},
		{
			name:        "ghec no trailing slash",
			url:         "https://api.customer.ghe.com",
			expectedURL: "https://api.customer.ghe.com/",
			isGHES:      false,
			errors:      false,
		},
		{
			name:        "ghec ui",
			url:         "https://customer.ghe.com/",
			expectedURL: "https://api.customer.ghe.com/",
			isGHES:      false,
			errors:      false,
		},
		{
			name:        "ghec http errors",
			url:         "http://api.customer.ghe.com/",
			expectedURL: "",
			isGHES:      false,
			errors:      true,
		},
		{
			name:        "ghec with path errors",
			url:         "https://api.customer.ghe.com/xxx/",
			expectedURL: "",
			isGHES:      false,
			errors:      true,
		},
		{
			name:        "ghes",
			url:         "https://example.com/",
			expectedURL: "https://example.com/",
			isGHES:      true,
			errors:      false,
		},
		{
			name:        "ghes no trailing slash",
			url:         "https://example.com",
			expectedURL: "https://example.com/",
			isGHES:      true,
			errors:      false,
		},
		{
			name:        "ghes with path prefix",
			url:         "https://example.com/test/",
			expectedURL: "https://example.com/test/",
			isGHES:      true,
			errors:      false,
		},
		{
			name:        "empty url errors",
			url:         "",
			expectedURL: "",
			isGHES:      false,
			errors:      true,
		},
		{
			name:        "not absolute url errors",
			url:         "example.com/",
			expectedURL: "",
			isGHES:      false,
			errors:      true,
		},
		{
			name:        "invalid url errors",
			url:         "xxx",
			expectedURL: "",
			isGHES:      false,
			errors:      true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			baseURL, isGHES, err := getBaseURL(tc.url)
			if err != nil {
				if tc.errors {
					return
				}
				t.Fatalf("expected no error, got: %v", err)
			}

			if tc.errors {
				t.Fatalf("expected error, got none")
			}

			if baseURL.String() != tc.expectedURL {
				t.Errorf("expected base URL %q, got %q", tc.expectedURL, baseURL.String())
			}

			if isGHES != tc.isGHES {
				t.Errorf("expected isGHES to be %v, got %v", tc.isGHES, isGHES)
			}
		})
	}
}

func TestPreviewHeaderInjectorTransport_RoundTrip(t *testing.T) {
	tests := []struct {
		name                string
		previewHeaders      map[string]string
		existingHeaders     map[string]string
		expectedHeaders     map[string]string
		expectRoundTripCall bool
	}{
		{
			name:                "empty preview headers",
			previewHeaders:      map[string]string{},
			existingHeaders:     map[string]string{"User-Agent": "test"},
			expectedHeaders:     map[string]string{"User-Agent": "test"},
			expectRoundTripCall: true,
		},
		{
			name: "add new preview header",
			previewHeaders: map[string]string{
				"Accept": "application/vnd.github.v3+json",
			},
			existingHeaders: map[string]string{},
			expectedHeaders: map[string]string{
				"Accept": "application/vnd.github.v3+json",
			},
			expectRoundTripCall: true,
		},
		{
			name: "append to existing header",
			previewHeaders: map[string]string{
				"Accept": "application/vnd.github.preview+json",
			},
			existingHeaders: map[string]string{
				"Accept": "application/json",
			},
			expectedHeaders: map[string]string{
				"Accept": "application/json,application/vnd.github.preview+json",
			},
			expectRoundTripCall: true,
		},
		{
			name: "preserve existing Accept application/octet-stream",
			previewHeaders: map[string]string{
				"Accept": "application/vnd.github.preview+json",
			},
			existingHeaders: map[string]string{
				"Accept": "application/octet-stream",
			},
			expectedHeaders: map[string]string{
				"Accept": "application/octet-stream",
			},
			expectRoundTripCall: true,
		},
		{
			name: "preserve existing accept application/octet-stream (lowercase)",
			previewHeaders: map[string]string{
				"accept": "application/vnd.github.preview+json",
			},
			existingHeaders: map[string]string{
				"accept": "application/octet-stream",
			},
			expectedHeaders: map[string]string{
				"Accept": "application/octet-stream",
			},
			expectRoundTripCall: true,
		},
		{
			name: "preserve existing Accept application/octet-stream (mixed case)",
			previewHeaders: map[string]string{
				"AcCePt": "application/vnd.github.preview+json",
			},
			existingHeaders: map[string]string{
				"Accept": "application/octet-stream",
			},
			expectedHeaders: map[string]string{
				"Accept": "application/octet-stream",
			},
			expectRoundTripCall: true,
		},
		{
			name: "multiple preview headers",
			previewHeaders: map[string]string{
				"Accept":               "application/vnd.github.v3+json",
				"X-GitHub-Api-Version": "2022-11-28",
			},
			existingHeaders: map[string]string{},
			expectedHeaders: map[string]string{
				"Accept":               "application/vnd.github.v3+json",
				"X-Github-Api-Version": "2022-11-28",
			},
			expectRoundTripCall: true,
		},
		{
			name: "append multiple preview headers to existing",
			previewHeaders: map[string]string{
				"Accept":               "application/vnd.github.v3+json",
				"X-GitHub-Api-Version": "2022-11-28",
			},
			existingHeaders: map[string]string{
				"Accept":               "application/json",
				"X-GitHub-Api-Version": "2021-01-01",
			},
			expectedHeaders: map[string]string{
				"Accept":               "application/json,application/vnd.github.v3+json",
				"X-Github-Api-Version": "2021-01-01,2022-11-28",
			},
			expectRoundTripCall: true,
		},
		{
			name: "non-accept headers always append",
			previewHeaders: map[string]string{
				"X-Custom-Header": "preview-value",
			},
			existingHeaders: map[string]string{
				"X-Custom-Header": "application/octet-stream",
			},
			expectedHeaders: map[string]string{
				"X-Custom-Header": "application/octet-stream,preview-value",
			},
			expectRoundTripCall: true,
		},
		{
			name: "accept header with different value appends",
			previewHeaders: map[string]string{
				"Accept": "application/vnd.github.preview+json",
			},
			existingHeaders: map[string]string{
				"Accept": "application/json",
			},
			expectedHeaders: map[string]string{
				"Accept": "application/json,application/vnd.github.preview+json",
			},
			expectRoundTripCall: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock RoundTripper that records the request
			var capturedRequest *http.Request
			mockRT := &mockRoundTripper{
				roundTripFunc: func(req *http.Request) (*http.Response, error) {
					capturedRequest = req
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       http.NoBody,
					}, nil
				},
			}

			injector := &previewHeaderInjectorTransport{
				rt:             mockRT,
				previewHeaders: tt.previewHeaders,
			}

			// Create a test request with existing headers
			req := httptest.NewRequest(http.MethodGet, "https://api.github.com/test", nil)
			for name, value := range tt.existingHeaders {
				req.Header.Set(name, value)
			}

			// Execute RoundTrip
			resp, err := injector.RoundTrip(req)
			// Verify no error
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Verify response
			if resp == nil {
				t.Fatal("expected non-nil response")
			}

			// Verify RoundTrip was called on the underlying transport
			if tt.expectRoundTripCall && capturedRequest == nil {
				t.Fatal("expected RoundTrip to be called on underlying transport")
			}

			// Verify headers in the captured request
			if capturedRequest != nil {
				for name, expectedValue := range tt.expectedHeaders {
					actualValue := capturedRequest.Header.Get(name)
					if actualValue != expectedValue {
						t.Errorf("header %q: expected %q, got %q", name, expectedValue, actualValue)
					}
				}

				// Verify no unexpected headers were added
				for name := range capturedRequest.Header {
					if _, exists := tt.expectedHeaders[name]; !exists {
						// Allow headers that were in existingHeaders but not in expectedHeaders
						if _, wasExisting := tt.existingHeaders[name]; !wasExisting {
							t.Errorf("unexpected header %q: %q", name, capturedRequest.Header.Get(name))
						}
					}
				}
			}
		})
	}
}

// mockRoundTripper is a mock implementation of http.RoundTripper for testing.
type mockRoundTripper struct {
	roundTripFunc func(*http.Request) (*http.Response, error)
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if m.roundTripFunc != nil {
		return m.roundTripFunc(req)
	}
	return &http.Response{StatusCode: http.StatusOK, Body: http.NoBody}, nil
}
