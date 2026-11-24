package ghclient

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"golang.org/x/oauth2"
)

func TestNewRESTClient(t *testing.T) {
	t.Parallel()

	t.Run("succeeds", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name        string
			opts        []ClientOption
			wantBaseURL string
		}{
			{
				name:        "no_options",
				opts:        nil,
				wantBaseURL: "https://api.github.com/",
			},
			{
				name:        "token",
				opts:        []ClientOption{WithClientToken("token")},
				wantBaseURL: "https://api.github.com/",
			},
			{
				name:        "url",
				opts:        []ClientOption{WithClientURL(new("https://ghes.example.com/api/v3/"))},
				wantBaseURL: "https://ghes.example.com/api/v3/",
			},
			{
				name:        "retry_disabled",
				opts:        []ClientOption{WithClientRetryOptions(0, 0, 0, 0)},
				wantBaseURL: "https://api.github.com/",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := NewRESTClient(tt.opts...)
				if err != nil {
					t.Fatalf("unexpected error: %s", err)
				}

				if got == nil {
					t.Fatal("client: got nil, want non-nil")
				}

				if got.BaseURL() != tt.wantBaseURL {
					t.Fatalf("base url: got %q, want %q", got.BaseURL(), tt.wantBaseURL)
				}
			})
		}
	})

	t.Run("applies_user_agent_and_timeout_options", func(t *testing.T) {
		t.Parallel()

		got, err := NewRESTClient(
			WithClientUserAgent(new("test-agent")),
			WithClientTimeout(time.Minute),
		)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if want := "test-agent"; got.UserAgent() != want {
			t.Fatalf("user agent: got %q, want %q", got.UserAgent(), want)
		}

		if want := time.Minute; got.Client().Timeout != want {
			t.Fatalf("timeout: got %s, want %s", got.Client().Timeout, want)
		}
	})

	t.Run("errors_when_option_fails", func(t *testing.T) {
		t.Parallel()

		got, err := NewRESTClient(WithClientToken(""))
		if err == nil {
			t.Fatal("got nil, want error")
		}

		if want := "failed to apply client option: invalid option: token cannot be empty when token authentication is enabled"; !strings.Contains(err.Error(), want) {
			t.Fatalf("got %q, want it to contain %q", err.Error(), want)
		}

		if !errors.Is(err, ErrInvalidOption) {
			t.Fatalf("got %v, want %v", err, ErrInvalidOption)
		}

		if got != nil {
			t.Fatal("client: got non-nil, want nil")
		}
	})
}

func Test_newRESTClient(t *testing.T) {
	t.Parallel()

	t.Run("uses_default_url_when_unset", func(t *testing.T) {
		t.Parallel()

		got, err := newRESTClient(defaultClientOptions())
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if want := "https://api.github.com/"; got.BaseURL() != want {
			t.Fatalf("base url: got %q, want %q", got.BaseURL(), want)
		}
	})

	t.Run("uses_url", func(t *testing.T) {
		t.Parallel()

		opts := defaultClientOptions()
		opts.url = new("https://ghes.example.com/api/v3/")

		got, err := newRESTClient(opts)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if want := "https://ghes.example.com/api/v3/"; got.BaseURL() != want {
			t.Fatalf("base url: got %q, want %q", got.BaseURL(), want)
		}
	})

	t.Run("uses_user_agent", func(t *testing.T) {
		t.Parallel()

		opts := defaultClientOptions()
		opts.userAgent = new("test-agent")

		got, err := newRESTClient(opts)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if want := "test-agent"; got.UserAgent() != want {
			t.Fatalf("user agent: got %q, want %q", got.UserAgent(), want)
		}
	})

	t.Run("uses_client_timeout", func(t *testing.T) {
		t.Parallel()

		opts := defaultClientOptions()
		opts.clientTimeout = time.Minute

		got, err := newRESTClient(opts)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if want := time.Minute; got.Client().Timeout != want {
			t.Fatalf("timeout: got %s, want %s", got.Client().Timeout, want)
		}
	})

	t.Run("errors_when_url_is_invalid", func(t *testing.T) {
		t.Parallel()

		opts := defaultClientOptions()
		opts.url = new("http://[::1")

		got, err := newRESTClient(opts)
		if err == nil {
			t.Fatal("got nil, want error")
		}

		if got != nil {
			t.Fatal("client: got non-nil, want nil")
		}
	})

	// httptest: the transport stack is private to the client, so the token and user agent can only be observed from a request.
	t.Run("sends_requests_with_token_and_user_agent", func(t *testing.T) {
		t.Parallel()

		type seen struct {
			method    string
			path      string
			userAgent string
			auth      string
		}

		reqs := make(chan seen, 1)

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqs <- seen{
				method:    r.Method,
				path:      r.URL.Path,
				userAgent: r.UserAgent(),
				auth:      r.Header.Get("Authorization"),
			}

			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{}`)
		}))
		t.Cleanup(ts.Close)

		opts := defaultClientOptions()
		opts.url = new(ts.URL + "/api/v3/")
		opts.userAgent = new("test-agent")
		opts.tokenSource = oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "token"})
		opts.retryMax = 0

		client, err := newRESTClient(opts)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		req, err := client.NewRequest(t.Context(), http.MethodGet, "ping", nil)
		if err != nil {
			t.Fatalf("unexpected error creating request: %s", err)
		}

		if _, err := client.Do(req, nil); err != nil {
			t.Fatalf("unexpected request error: %s", err)
		}

		got := <-reqs

		if got.method != http.MethodGet {
			t.Fatalf("method: got %q, want %q", got.method, http.MethodGet)
		}

		if want := "/api/v3/ping"; got.path != want {
			t.Fatalf("path: got %q, want %q", got.path, want)
		}

		if want := "test-agent"; got.userAgent != want {
			t.Fatalf("user agent: got %q, want %q", got.userAgent, want)
		}

		if want := "Bearer token"; got.auth != want {
			t.Fatalf("authorization: got %q, want %q", got.auth, want)
		}
	})
}

func Test_restURL(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		base   string
		isGHES bool
		want   string
	}{
		{
			name:   "dotcom",
			base:   "",
			isGHES: false,
			want:   "https://api.github.com/",
		},
		{
			name:   "ghes",
			base:   "https://ghes.example.com",
			isGHES: true,
			want:   "https://ghes.example.com/api/v3/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var base *url.URL
			if tt.base != "" {
				base = mustParseURL(t, tt.base)
			}

			got, err := restURL(base, tt.isGHES)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if got.String() != tt.want {
				t.Fatalf("url: got %q, want %q", got.String(), tt.want)
			}
		})
	}
}

func Test_apiURL(t *testing.T) {
	t.Parallel()

	t.Run("succeeds", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name   string
			base   string
			isGHES bool
			want   string
		}{
			{
				name:   "nil_url_uses_dotcom",
				base:   "",
				isGHES: false,
				want:   "https://api.github.com/path",
			},
			{
				name:   "non_ghes_url",
				base:   "https://api.example.ghe.com",
				isGHES: false,
				want:   "https://api.example.ghe.com/path",
			},
			{
				name:   "non_ghes_url_with_path_prefix",
				base:   "https://api.example.ghe.com/prefix",
				isGHES: false,
				want:   "https://api.example.ghe.com/prefix/path",
			},
			{
				name:   "ghes_url",
				base:   "https://ghes.example.com",
				isGHES: true,
				want:   "https://ghes.example.com/ghes/path",
			},
			{
				name:   "ghes_url_with_trailing_slash",
				base:   "https://ghes.example.com/",
				isGHES: true,
				want:   "https://ghes.example.com/ghes/path",
			},
			{
				name:   "ghes_url_with_path_prefix",
				base:   "https://ghes.example.com/prefix",
				isGHES: true,
				want:   "https://ghes.example.com/prefix/ghes/path",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var base *url.URL
				if tt.base != "" {
					base = mustParseURL(t, tt.base)
				}

				got, err := apiURL(base, tt.isGHES, "path", "ghes/path")
				if err != nil {
					t.Fatalf("unexpected error: %s", err)
				}

				if got.String() != tt.want {
					t.Fatalf("url: got %q, want %q", got.String(), tt.want)
				}
			})
		}
	})

	t.Run("does_not_modify_input", func(t *testing.T) {
		t.Parallel()

		base := mustParseURL(t, "https://ghes.example.com/prefix")

		if _, err := apiURL(base, true, "path", "ghes/path"); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if want := "https://ghes.example.com/prefix"; base.String() != want {
			t.Fatalf("input url: got %q, want %q", base.String(), want)
		}
	})

	t.Run("does_not_modify_default_url", func(t *testing.T) {
		t.Parallel()

		if _, err := apiURL(nil, false, "path", "ghes/path"); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if got := dotComAPIURL.String(); got != DotComAPIURLStr {
			t.Fatalf("default url: got %q, want %q", got, DotComAPIURLStr)
		}
	})

	t.Run("errors_when_ghes_without_url", func(t *testing.T) {
		t.Parallel()

		got, err := apiURL(nil, true, "path", "ghes/path")
		if err == nil {
			t.Fatal("got nil, want error")
		}

		if want := "base URL cannot be nil when ghes is enabled"; !strings.Contains(err.Error(), want) {
			t.Fatalf("got %q, want it to contain %q", err.Error(), want)
		}

		if !errors.Is(err, ErrInvalidOption) {
			t.Fatalf("got %v, want %v", err, ErrInvalidOption)
		}

		if got != nil {
			t.Fatal("url: got non-nil, want nil")
		}
	})
}
