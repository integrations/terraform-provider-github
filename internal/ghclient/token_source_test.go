package ghclient

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func Test_newTokenSource(t *testing.T) {
	t.Parallel()

	t.Run("succeeds", func(t *testing.T) {
		t.Parallel()

		opts := defaultSourceOptions()
		opts.tokenSource = oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "token"})

		got, err := newTokenSource(opts)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if got.restClient == nil {
			t.Fatal("rest client: got nil, want non-nil")
		}

		if got.graphqlClient == nil {
			t.Fatal("graphql client: got nil, want non-nil")
		}
	})

	t.Run("succeeds_without_token_source", func(t *testing.T) {
		t.Parallel()

		got, err := newTokenSource(defaultSourceOptions())
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if got.restClient == nil {
			t.Fatal("rest client: got nil, want non-nil")
		}

		if got.graphqlClient == nil {
			t.Fatal("graphql client: got nil, want non-nil")
		}
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

				got, err := newTokenSource(tt.opts())
				if err == nil {
					t.Fatal("got nil, want error")
				}

				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("got %q, want it to contain %q", err.Error(), tt.wantErr)
				}

				if got != nil {
					t.Fatal("source: got non-nil, want nil")
				}
			})
		}
	})

	// httptest: the token source is only observable through the Authorization header the clients send.
	t.Run("authenticates_both_clients_with_the_token_source", func(t *testing.T) {
		t.Parallel()

		auths := make(chan string, 2)

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auths <- r.Header.Get("Authorization")

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
		opts.retryMax = 0
		opts.tokenSource = oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "token"})

		s, err := newTokenSource(opts)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		req, err := s.restClient.NewRequest(t.Context(), http.MethodGet, "ping", nil)
		if err != nil {
			t.Fatalf("unexpected error creating rest request: %s", err)
		}

		if _, err := s.restClient.Do(req, nil); err != nil {
			t.Fatalf("unexpected rest error: %s", err)
		}

		var q struct {
			Viewer struct {
				Login githubv4.String
			}
		}
		if err := s.graphqlClient.Query(t.Context(), &q, nil); err != nil {
			t.Fatalf("unexpected graphql error: %s", err)
		}

		for _, name := range []string{"rest", "graphql"} {
			if got, want := <-auths, "Bearer token"; got != want {
				t.Fatalf("%s authorization: got %q, want %q", name, got, want)
			}
		}
	})
}

func Test_tokenSource(t *testing.T) {
	t.Parallel()

	owners := []string{"", "org", "user"}

	t.Run("RESTClient", func(t *testing.T) {
		t.Parallel()

		opts := defaultSourceOptions()
		opts.tokenSource = oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "token"})

		s, err := newTokenSource(opts)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		got, err := s.RESTClient(t.Context())
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if got != s.restClient {
			t.Fatal("rest client: got different client, want source client")
		}
	})

	t.Run("OwnerRESTClient", func(t *testing.T) {
		t.Parallel()

		opts := defaultSourceOptions()
		opts.tokenSource = oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "token"})

		s, err := newTokenSource(opts)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		for _, owner := range owners {
			got, err := s.OwnerRESTClient(t.Context(), owner)
			if err != nil {
				t.Fatalf("unexpected error for owner %q: %s", owner, err)
			}

			if got != s.restClient {
				t.Fatalf("rest client for owner %q: got different client, want source client", owner)
			}
		}
	})

	t.Run("GraphQLClient", func(t *testing.T) {
		t.Parallel()

		opts := defaultSourceOptions()
		opts.tokenSource = oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "token"})

		s, err := newTokenSource(opts)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		got, err := s.GraphQLClient(t.Context())
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if got != s.graphqlClient {
			t.Fatal("graphql client: got different client, want source client")
		}
	})

	t.Run("OwnerGraphQLClient", func(t *testing.T) {
		t.Parallel()

		opts := defaultSourceOptions()
		opts.tokenSource = oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "token"})

		s, err := newTokenSource(opts)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		for _, owner := range owners {
			got, err := s.OwnerGraphQLClient(t.Context(), owner)
			if err != nil {
				t.Fatalf("unexpected error for owner %q: %s", owner, err)
			}

			if got != s.graphqlClient {
				t.Fatalf("graphql client for owner %q: got different client, want source client", owner)
			}
		}
	})
}
