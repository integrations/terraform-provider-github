package ghclient

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func TestNewGraphQLClient(t *testing.T) {
	t.Parallel()

	t.Run("succeeds", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name string
			opts []ClientOption
		}{
			{
				name: "no_options",
				opts: nil,
			},
			{
				name: "token",
				opts: []ClientOption{WithClientToken("token")},
			},
			{
				name: "url",
				opts: []ClientOption{WithClientURL(new("https://ghes.example.com/api/graphql"))},
			},
			{
				name: "retry_disabled",
				opts: []ClientOption{WithClientRetryOptions(0, 0, 0, 0)},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := NewGraphQLClient(tt.opts...)
				if err != nil {
					t.Fatalf("unexpected error: %s", err)
				}

				if got == nil {
					t.Fatal("client: got nil, want non-nil")
				}
			})
		}
	})

	t.Run("errors", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name    string
			opts    []ClientOption
			wantErr string
		}{
			{
				name:    "option_error_is_wrapped",
				opts:    []ClientOption{WithClientToken("")},
				wantErr: "failed to apply client option: invalid option: token cannot be empty when token authentication is enabled",
			},
			{
				name:    "user_agent_is_not_supported",
				opts:    []ClientOption{WithClientUserAgent(new("test-agent"))},
				wantErr: "user agent is not supported for graphql clients",
			},
			{
				// Sets the flag directly so no cache store is opened (and leaked) by the test.
				name: "cache_is_not_supported",
				opts: []ClientOption{
					func(o *clientOptions) error {
						o.cacheEnabled = true

						return nil
					},
				},
				wantErr: "caching is not supported for graphql clients",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := NewGraphQLClient(tt.opts...)
				if err == nil {
					t.Fatal("got nil, want error")
				}

				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("got %q, want it to contain %q", err.Error(), tt.wantErr)
				}

				if !errors.Is(err, ErrInvalidOption) {
					t.Fatalf("got %v, want %v", err, ErrInvalidOption)
				}

				if got != nil {
					t.Fatal("client: got non-nil, want nil")
				}
			})
		}
	})
}

func Test_newGraphQLClient(t *testing.T) {
	t.Parallel()

	t.Run("returns_client_without_url", func(t *testing.T) {
		t.Parallel()

		got, err := newGraphQLClient(defaultClientOptions())
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if got == nil {
			t.Fatal("client: got nil, want non-nil")
		}
	})

	// httptest: the client has no getters, so the URL and authentication can only be observed from a request.
	t.Run("sends_requests_to_url_with_token", func(t *testing.T) {
		t.Parallel()

		type seen struct {
			method string
			path   string
			auth   string
		}

		reqs := make(chan seen, 1)

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqs <- seen{method: r.Method, path: r.URL.Path, auth: r.Header.Get("Authorization")}

			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"data":{"viewer":{"login":"octocat"}}}`)
		}))
		t.Cleanup(ts.Close)

		opts := defaultClientOptions()
		opts.url = new(ts.URL + "/api/graphql")
		opts.tokenSource = oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "token"})
		opts.retryMax = 0

		client, err := newGraphQLClient(opts)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		var q struct {
			Viewer struct {
				Login githubv4.String
			}
		}
		if err := client.Query(t.Context(), &q, nil); err != nil {
			t.Fatalf("unexpected query error: %s", err)
		}

		if want := "octocat"; string(q.Viewer.Login) != want {
			t.Fatalf("login: got %q, want %q", q.Viewer.Login, want)
		}

		got := <-reqs

		if got.method != http.MethodPost {
			t.Fatalf("method: got %q, want %q", got.method, http.MethodPost)
		}

		if want := "/api/graphql"; got.path != want {
			t.Fatalf("path: got %q, want %q", got.path, want)
		}

		if want := "Bearer token"; got.auth != want {
			t.Fatalf("authorization: got %q, want %q", got.auth, want)
		}
	})
}

func Test_graphqlURL(t *testing.T) {
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
			want:   "https://api.github.com/graphql",
		},
		{
			name:   "ghes",
			base:   "https://ghes.example.com",
			isGHES: true,
			want:   "https://ghes.example.com/api/graphql",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var base *url.URL
			if tt.base != "" {
				base = mustParseURL(t, tt.base)
			}

			got, err := graphqlURL(base, tt.isGHES)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if got.String() != tt.want {
				t.Fatalf("url: got %q, want %q", got.String(), tt.want)
			}
		})
	}
}
