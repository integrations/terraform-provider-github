package ghclient

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/google/go-github/v92/github"
	"github.com/shurcooL/githubv4"
)

func Test_newAppSource(t *testing.T) {
	t.Parallel()

	t.Run("succeeds", func(t *testing.T) {
		t.Parallel()

		signer := mustReadAppSigner(t)

		opts := defaultSourceOptions()
		opts.appClientID = "123456789"
		opts.appSigner = signer

		got, err := newAppSource(opts)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if want := "https://api.github.com/"; got.restURL != want {
			t.Fatalf("rest url: got %q, want %q", got.restURL, want)
		}

		if got.opts.appClientID != opts.appClientID {
			t.Fatalf("client id: got %q, want %q", got.opts.appClientID, opts.appClientID)
		}

		if got.opts.appSigner != signer {
			t.Fatal("signer: got different signer, want same signer")
		}

		if got.groups == nil {
			t.Fatal("groups: got nil, want non-nil")
		}
	})

	t.Run("errors_when_ghes_without_base_url", func(t *testing.T) {
		t.Parallel()

		opts := defaultSourceOptions()
		opts.isGHES = true

		got, err := newAppSource(opts)
		if err == nil {
			t.Fatal("got nil, want error")
		}

		if want := "failed to create rest url"; !strings.Contains(err.Error(), want) {
			t.Fatalf("got %q, want it to contain %q", err.Error(), want)
		}

		if got != nil {
			t.Fatal("source: got non-nil, want nil")
		}
	})
}

func Test_appSource(t *testing.T) {
	t.Parallel()

	t.Run("getGroup", func(t *testing.T) {
		t.Parallel()

		t.Run("returns_false_when_missing", func(t *testing.T) {
			t.Parallel()

			opts := defaultSourceOptions()
			opts.appClientID = "123456789"
			opts.appSigner = mustReadAppSigner(t)

			s, err := newAppSource(opts)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			got, ok := s.getGroup("org")
			if ok {
				t.Fatal("group found: got true, want false")
			}

			if got != nil {
				t.Fatal("group: got non-nil, want nil")
			}
		})

		t.Run("returns_stored_group", func(t *testing.T) {
			t.Parallel()

			opts := defaultSourceOptions()
			opts.appClientID = "123456789"
			opts.appSigner = mustReadAppSigner(t)

			s, err := newAppSource(opts)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			want := &sourceClientGroup{}
			s.groups["org"] = want

			got, ok := s.getGroup("org")
			if !ok {
				t.Fatal("group found: got false, want true")
			}

			if got != want {
				t.Fatal("group: got different group, want same group")
			}
		})
	})

	t.Run("group", func(t *testing.T) {
		t.Parallel()

		t.Run("returns_existing_group_ignoring_owner_case", func(t *testing.T) {
			t.Parallel()

			opts := defaultSourceOptions()
			opts.appClientID = "123456789"
			opts.appSigner = mustReadAppSigner(t)

			s, err := newAppSource(opts)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			want := &sourceClientGroup{}
			s.groups["org"] = want

			got, err := s.group(t.Context(), new("ORG"))
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if got != want {
				t.Fatal("group: got different group, want existing group")
			}
		})

		t.Run("creates_app_group_for_nil_owner", func(t *testing.T) {
			t.Parallel()

			opts := defaultSourceOptions()
			opts.appClientID = "123456789"
			opts.appSigner = mustReadAppSigner(t)
			opts.cacheEnabled = true
			opts.cacheBasePath = t.TempDir()

			s, err := newAppSource(opts)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			got, err := s.group(t.Context(), nil)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			t.Cleanup(func() { _ = got.cleanup() })

			if got.restClient == nil {
				t.Fatal("rest client: got nil, want non-nil")
			}

			if got.graphqlClient == nil {
				t.Fatal("graphql client: got nil, want non-nil")
			}

			again, err := s.group(t.Context(), nil)
			if err != nil {
				t.Fatalf("unexpected error on second call: %s", err)
			}

			if again != got {
				t.Fatal("group on repeated call: got different group, want same group")
			}

			if stored, ok := s.getGroup("_"); !ok || stored != got {
				t.Fatal("group stored under app key: got false, want true")
			}

			if _, err := os.Stat(filepath.Join(opts.cacheBasePath, "app-123456789", "cache.db")); err != nil {
				t.Fatalf("stat app cache database: got %v, want nil", err)
			}
		})

		t.Run("errors_when_token_source_cannot_be_created", func(t *testing.T) {
			t.Parallel()

			key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
			if err != nil {
				t.Fatalf("failed to generate key: %s", err)
			}

			opts := defaultSourceOptions()
			opts.appClientID = "123456789"
			opts.appSigner = key

			s, err := newAppSource(opts)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			got, err := s.group(t.Context(), nil)
			if err == nil {
				t.Fatal("got nil, want error")
			}

			if want := "failed to create app token source"; !strings.Contains(err.Error(), want) {
				t.Fatalf("got %q, want it to contain %q", err.Error(), want)
			}

			if got != nil {
				t.Fatal("group: got non-nil, want nil")
			}
		})

		t.Run("errors_when_client_group_cannot_be_created", func(t *testing.T) {
			t.Parallel()

			opts := defaultSourceOptions()
			opts.appClientID = "123456789"
			opts.appSigner = mustReadAppSigner(t)
			opts.cacheEnabled = true
			opts.cacheBasePath = "\x00c"

			s, err := newAppSource(opts)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			got, err := s.group(t.Context(), nil)
			if err == nil {
				t.Fatal("got nil, want error")
			}

			if want := "failed to create app client group"; !strings.Contains(err.Error(), want) {
				t.Fatalf("got %q, want it to contain %q", err.Error(), want)
			}

			if got != nil {
				t.Fatal("group: got non-nil, want nil")
			}

			if _, ok := s.getGroup("_"); ok {
				t.Fatal("failed group stored: got true, want false")
			}
		})

		// httptest: the owner lookup goes through the real transport to a URL built from the base URL.
		t.Run("creates_owner_group_with_lowercase_key_and_cache_dir", func(t *testing.T) {
			t.Parallel()

			gotOwner := atomic.Pointer[string]{}

			mux := http.NewServeMux()
			mux.HandleFunc("/api/v3/orgs/{org}/installation", func(w http.ResponseWriter, r *http.Request) {
				owner := r.PathValue("org")
				gotOwner.Store(&owner)

				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"id": 42}`))
			})

			ts := httptest.NewServer(mux)
			t.Cleanup(ts.Close)

			opts := defaultSourceOptions()
			opts.baseURL = mustParseURL(t, ts.URL)
			opts.isGHES = true
			opts.retryMax = 0
			opts.appClientID = "123456789"
			opts.appSigner = mustReadAppSigner(t)
			opts.cacheEnabled = true
			opts.cacheBasePath = t.TempDir()

			s, err := newAppSource(opts)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			got, err := s.group(t.Context(), new("Org"))
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			t.Cleanup(func() { _ = got.cleanup() })

			if ag, ok := s.getGroup("_"); ok {
				t.Cleanup(func() { _ = ag.cleanup() })
			}

			if owner := gotOwner.Load(); owner == nil || *owner != "Org" {
				t.Fatalf("lookup owner: got %v, want Org", owner)
			}

			if stored, ok := s.getGroup("org"); !ok || stored != got {
				t.Fatal("group stored under lowercase owner key: got false, want true")
			}

			if _, err := os.Stat(filepath.Join(opts.cacheBasePath, "app-123456789-org", "cache.db")); err != nil {
				t.Fatalf("stat lowercase owner cache database: got %v, want nil", err)
			}
		})

		// httptest: coalescing only matters across real in-flight lookups.
		t.Run("coalesces_concurrent_callers", func(t *testing.T) {
			t.Parallel()

			const callers = 8

			calls := atomic.Int32{}
			release := make(chan struct{})

			mux := http.NewServeMux()
			mux.HandleFunc("/api/v3/orgs/{org}/installation", func(w http.ResponseWriter, _ *http.Request) {
				calls.Add(1)
				<-release

				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"id": 42}`))
			})

			ts := httptest.NewServer(mux)
			t.Cleanup(ts.Close)

			opts := defaultSourceOptions()
			opts.baseURL = mustParseURL(t, ts.URL)
			opts.isGHES = true
			opts.retryMax = 0
			opts.appClientID = "123456789"
			opts.appSigner = mustReadAppSigner(t)

			s, err := newAppSource(opts)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			started := sync.WaitGroup{}
			done := sync.WaitGroup{}
			groups := make([]*sourceClientGroup, callers)

			for i := range callers {
				owner := "org"
				if i%2 == 1 {
					owner = "ORG"
				}

				started.Add(1)

				done.Go(func() {
					started.Done()

					g, err := s.group(t.Context(), &owner)
					if err != nil {
						t.Errorf("unexpected error: %s", err)

						return
					}

					groups[i] = g
				})
			}

			started.Wait()
			close(release)
			done.Wait()

			if got := calls.Load(); got != 1 {
				t.Fatalf("installation lookups: got %d, want %d", got, 1)
			}

			for i, g := range groups {
				if g == nil || g != groups[0] {
					t.Fatalf("caller %d group: got different group, want shared group", i)
				}
			}
		})

		// httptest: needs a lookup that is genuinely in flight when the caller gives up.
		t.Run("cancelled_caller_does_not_fail_other_callers", func(t *testing.T) {
			t.Parallel()

			entered := make(chan struct{}, 1)
			release := make(chan struct{})

			mux := http.NewServeMux()
			mux.HandleFunc("/api/v3/orgs/{org}/installation", func(w http.ResponseWriter, _ *http.Request) {
				entered <- struct{}{}
				<-release

				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"id": 42}`))
			})

			ts := httptest.NewServer(mux)
			t.Cleanup(ts.Close)

			opts := defaultSourceOptions()
			opts.baseURL = mustParseURL(t, ts.URL)
			opts.isGHES = true
			opts.retryMax = 0
			opts.appClientID = "123456789"
			opts.appSigner = mustReadAppSigner(t)

			s, err := newAppSource(opts)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			cancelCtx, cancel := context.WithCancel(t.Context())
			defer cancel()

			cancelledErr := make(chan error, 1)
			go func() {
				_, err := s.group(cancelCtx, new("org"))
				cancelledErr <- err
			}()

			otherGroup := make(chan *sourceClientGroup, 1)
			otherErr := make(chan error, 1)
			go func() {
				g, err := s.group(t.Context(), new("org"))
				otherGroup <- g
				otherErr <- err
			}()

			<-entered
			cancel()

			if err := <-cancelledErr; !errors.Is(err, context.Canceled) {
				t.Fatalf("cancelled caller: got %v, want %v", err, context.Canceled)
			}

			close(release)

			if err := <-otherErr; err != nil {
				t.Fatalf("unexpected error for other caller: %s", err)
			}

			if g := <-otherGroup; g == nil {
				t.Fatal("other caller group: got nil, want non-nil")
			}
		})

		// httptest: needs a lookup that fails and then succeeds against the same source.
		t.Run("does_not_cache_failures", func(t *testing.T) {
			t.Parallel()

			installed := atomic.Bool{}

			mux := http.NewServeMux()
			mux.HandleFunc("/api/v3/orgs/{org}/installation", func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Type", "application/json")

				if !installed.Load() {
					w.WriteHeader(http.StatusNotFound)
					_, _ = w.Write([]byte(`{"message": "Not Found"}`))

					return
				}

				_, _ = w.Write([]byte(`{"id": 42}`))
			})
			mux.HandleFunc("/api/v3/users/{user}/installation", func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write([]byte(`{"message": "Not Found"}`))
			})

			ts := httptest.NewServer(mux)
			t.Cleanup(ts.Close)

			opts := defaultSourceOptions()
			opts.baseURL = mustParseURL(t, ts.URL)
			opts.isGHES = true
			opts.retryMax = 0
			opts.appClientID = "123456789"
			opts.appSigner = mustReadAppSigner(t)

			s, err := newAppSource(opts)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			_, err = s.group(t.Context(), new("org"))
			if err == nil {
				t.Fatal("got nil, want error")
			}

			if want := `failed to get installation id: app installation not found for owner "org"`; !strings.Contains(err.Error(), want) {
				t.Fatalf("got %q, want it to contain %q", err.Error(), want)
			}

			if !errors.Is(err, ErrInstallationNotFound) {
				t.Fatalf("got %v, want %v", err, ErrInstallationNotFound)
			}

			if _, ok := s.getGroup("org"); ok {
				t.Fatal("failed group stored: got true, want false")
			}

			installed.Store(true)

			got, err := s.group(t.Context(), new("org"))
			if err != nil {
				t.Fatalf("unexpected error on second call: %s", err)
			}

			if got == nil {
				t.Fatal("group on second call: got nil, want non-nil")
			}
		})
	})

	t.Run("RESTClient", func(t *testing.T) {
		t.Parallel()

		t.Run("succeeds", func(t *testing.T) {
			t.Parallel()

			opts := defaultSourceOptions()
			opts.appClientID = "123456789"
			opts.appSigner = mustReadAppSigner(t)

			s, err := newAppSource(opts)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			got, err := s.RESTClient(t.Context())
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			sg, ok := s.getGroup("_")
			if !ok {
				t.Fatal("app group stored: got false, want true")
			}

			if got != sg.restClient {
				t.Fatal("rest client: got different client, want app group client")
			}
		})

		t.Run("errors_if_client_group_cannot_be_created", func(t *testing.T) {
			t.Parallel()

			opts := defaultSourceOptions()
			opts.appClientID = "123456789"
			opts.appSigner = mustReadAppSigner(t)
			opts.cacheEnabled = true
			opts.cacheBasePath = "\x00c"

			s, err := newAppSource(opts)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			_, err = s.RESTClient(t.Context())
			if err == nil {
				t.Fatal("got nil, want error")
			}

			if want := "failed to create app client group"; !strings.Contains(err.Error(), want) {
				t.Fatalf("got %q, want it to contain %q", err.Error(), want)
			}
		})
	})

	t.Run("OwnerRESTClient", func(t *testing.T) {
		t.Parallel()

		t.Run("succeeds", func(t *testing.T) {
			t.Parallel()

			opts := defaultSourceOptions()
			opts.appClientID = "123456789"
			opts.appSigner = mustReadAppSigner(t)

			s, err := newAppSource(opts)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			want := new(github.Client)
			s.groups["org"] = &sourceClientGroup{restClient: want}

			got, err := s.OwnerRESTClient(t.Context(), "Org")
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if got != want {
				t.Fatal("rest client: got different client, want owner group client")
			}
		})

		t.Run("errors_if_client_group_cannot_be_created", func(t *testing.T) {
			t.Parallel()

			opts := defaultSourceOptions()
			opts.appClientID = "123456789"
			opts.appSigner = mustReadAppSigner(t)
			opts.cacheEnabled = true
			opts.cacheBasePath = "\x00c"

			s, err := newAppSource(opts)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			_, err = s.OwnerRESTClient(t.Context(), "org")
			if err == nil {
				t.Fatal("got nil, want error")
			}

			if want := "failed to get installation id"; !strings.Contains(err.Error(), want) {
				t.Fatalf("got %q, want it to contain %q", err.Error(), want)
			}
		})
	})

	t.Run("GraphQLClient", func(t *testing.T) {
		t.Parallel()

		t.Run("succeeds", func(t *testing.T) {
			t.Parallel()

			opts := defaultSourceOptions()
			opts.appClientID = "123456789"
			opts.appSigner = mustReadAppSigner(t)

			s, err := newAppSource(opts)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			got, err := s.GraphQLClient(t.Context())
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			sg, ok := s.getGroup("_")
			if !ok {
				t.Fatal("app group stored: got false, want true")
			}

			if got != sg.graphqlClient {
				t.Fatal("graphql client: got different client, want app group client")
			}
		})

		t.Run("errors_if_client_group_cannot_be_created", func(t *testing.T) {
			t.Parallel()

			opts := defaultSourceOptions()
			opts.appClientID = "123456789"
			opts.appSigner = mustReadAppSigner(t)
			opts.cacheEnabled = true
			opts.cacheBasePath = "\x00c"

			s, err := newAppSource(opts)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			_, err = s.GraphQLClient(t.Context())
			if err == nil {
				t.Fatal("got nil, want error")
			}

			if want := "failed to create app client group"; !strings.Contains(err.Error(), want) {
				t.Fatalf("got %q, want it to contain %q", err.Error(), want)
			}
		})
	})

	t.Run("OwnerGraphQLClient", func(t *testing.T) {
		t.Parallel()

		t.Run("succeeds", func(t *testing.T) {
			t.Parallel()

			opts := defaultSourceOptions()
			opts.appClientID = "123456789"
			opts.appSigner = mustReadAppSigner(t)

			s, err := newAppSource(opts)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			want := new(githubv4.Client)
			s.groups["org"] = &sourceClientGroup{graphqlClient: want}

			got, err := s.OwnerGraphQLClient(t.Context(), "Org")
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if got != want {
				t.Fatal("graphql client: got different client, want owner group client")
			}
		})

		t.Run("errors_if_client_group_cannot_be_created", func(t *testing.T) {
			t.Parallel()

			opts := defaultSourceOptions()
			opts.appClientID = "123456789"
			opts.appSigner = mustReadAppSigner(t)
			opts.cacheEnabled = true
			opts.cacheBasePath = "\x00c"

			s, err := newAppSource(opts)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			_, err = s.OwnerGraphQLClient(t.Context(), "org")
			if err == nil {
				t.Fatal("got nil, want error")
			}

			if want := "failed to get installation id"; !strings.Contains(err.Error(), want) {
				t.Fatalf("got %q, want it to contain %q", err.Error(), want)
			}
		})
	})

	t.Run("installationID", func(t *testing.T) {
		t.Parallel()

		t.Run("succeeds", func(t *testing.T) {
			t.Parallel()

			for _, tt := range []struct {
				name          string
				orgStatus     int
				orgBody       string
				userStatus    int
				userBody      string
				want          int64
				wantUserCalls int32
			}{
				{
					name:          "org_installation",
					orgStatus:     http.StatusOK,
					orgBody:       `{"id": 42}`,
					userStatus:    http.StatusInternalServerError,
					userBody:      "",
					want:          42,
					wantUserCalls: 0,
				},
				{
					name:          "user_installation_when_org_not_found",
					orgStatus:     http.StatusNotFound,
					orgBody:       `{"message": "Not Found"}`,
					userStatus:    http.StatusOK,
					userBody:      `{"id": 43}`,
					want:          43,
					wantUserCalls: 1,
				},
			} {
				t.Run(tt.name, func(t *testing.T) {
					t.Parallel()

					userCalls := atomic.Int32{}

					mux := http.NewServeMux()
					mux.HandleFunc("/api/v3/orgs/{org}/installation", func(w http.ResponseWriter, _ *http.Request) {
						w.Header().Set("Content-Type", "application/json")
						w.WriteHeader(tt.orgStatus)
						_, _ = w.Write([]byte(tt.orgBody))
					})
					mux.HandleFunc("/api/v3/users/{user}/installation", func(w http.ResponseWriter, _ *http.Request) {
						userCalls.Add(1)

						w.Header().Set("Content-Type", "application/json")
						w.WriteHeader(tt.userStatus)
						_, _ = w.Write([]byte(tt.userBody))
					})

					ts := httptest.NewServer(mux)
					t.Cleanup(ts.Close)

					opts := defaultSourceOptions()
					opts.baseURL = mustParseURL(t, ts.URL)
					opts.isGHES = true
					opts.retryMax = 0
					opts.appClientID = "123456789"
					opts.appSigner = mustReadAppSigner(t)

					s, err := newAppSource(opts)
					if err != nil {
						t.Fatalf("unexpected error: %s", err)
					}

					got, err := s.installationID(t.Context(), "org")
					if err != nil {
						t.Fatalf("unexpected error: %s", err)
					}

					if got == 0 || got != tt.want {
						t.Fatalf("installation id: got %v, want %d", got, tt.want)
					}

					if got := userCalls.Load(); got != tt.wantUserCalls {
						t.Fatalf("user lookups: got %d, want %d", got, tt.wantUserCalls)
					}
				})
			}
		})

		t.Run("errors", func(t *testing.T) {
			t.Parallel()

			for _, tt := range []struct {
				name          string
				orgConnClose  bool
				orgStatus     int
				orgBody       string
				userStatus    int
				userBody      string
				wantUserCalls int32
				wantErr       string
				wantNotFound  bool
			}{
				{
					name:          "not_installed",
					orgStatus:     http.StatusNotFound,
					orgBody:       `{"message": "Not Found"}`,
					userStatus:    http.StatusNotFound,
					userBody:      `{"message": "Not Found"}`,
					wantUserCalls: 1,
					wantErr:       `app installation not found for owner "org"`,
					wantNotFound:  true,
				},
				{
					name:          "user_lookup_error_after_org_not_found",
					orgStatus:     http.StatusNotFound,
					orgBody:       `{"message": "Not Found"}`,
					userStatus:    http.StatusInternalServerError,
					userBody:      `{"message": "Server Error"}`,
					wantUserCalls: 1,
					wantErr:       `failed to get installation for owner "org"`,
				},
				{
					name:          "org_lookup_error_skips_user_lookup",
					orgStatus:     http.StatusUnauthorized,
					orgBody:       `{"message": "Bad credentials"}`,
					userStatus:    http.StatusOK,
					userBody:      `{"id": 43}`,
					wantUserCalls: 0,
					wantErr:       `failed to get installation for owner "org"`,
				},
				{
					name:          "org_lookup_network_error_skips_user_lookup",
					orgConnClose:  true,
					userStatus:    http.StatusOK,
					userBody:      `{"id": 43}`,
					wantUserCalls: 0,
					wantErr:       `failed to get installation for owner "org"`,
				},
				{
					name:          "org_installation_without_id",
					orgStatus:     http.StatusOK,
					orgBody:       `{}`,
					userStatus:    http.StatusInternalServerError,
					userBody:      "",
					wantUserCalls: 0,
					wantErr:       `installation for owner "org" has no id`,
				},
				{
					name:          "user_installation_without_id",
					orgStatus:     http.StatusNotFound,
					orgBody:       `{"message": "Not Found"}`,
					userStatus:    http.StatusOK,
					userBody:      `{}`,
					wantUserCalls: 1,
					wantErr:       `installation for owner "org" has no id`,
				},
			} {
				t.Run(tt.name, func(t *testing.T) {
					t.Parallel()

					userCalls := atomic.Int32{}

					mux := http.NewServeMux()
					mux.HandleFunc("/api/v3/orgs/{org}/installation", func(w http.ResponseWriter, _ *http.Request) {
						if tt.orgConnClose {
							hj, ok := w.(http.Hijacker)
							if !ok {
								t.Error("response writer: got non-hijacker, want hijacker")
								return
							}

							conn, _, err := hj.Hijack()
							if err != nil {
								t.Errorf("unexpected error: %s", err)
								return
							}

							_ = conn.Close()
							return
						}

						w.Header().Set("Content-Type", "application/json")
						w.WriteHeader(tt.orgStatus)
						_, _ = w.Write([]byte(tt.orgBody))
					})
					mux.HandleFunc("/api/v3/users/{user}/installation", func(w http.ResponseWriter, _ *http.Request) {
						userCalls.Add(1)

						w.Header().Set("Content-Type", "application/json")
						w.WriteHeader(tt.userStatus)
						_, _ = w.Write([]byte(tt.userBody))
					})

					ts := httptest.NewServer(mux)
					t.Cleanup(ts.Close)

					opts := defaultSourceOptions()
					opts.baseURL = mustParseURL(t, ts.URL)
					opts.isGHES = true
					opts.retryMax = 0
					opts.appClientID = "123456789"
					opts.appSigner = mustReadAppSigner(t)

					s, err := newAppSource(opts)
					if err != nil {
						t.Fatalf("unexpected error: %s", err)
					}

					got, err := s.installationID(t.Context(), "org")
					if err == nil {
						t.Fatal("got nil, want error")
					}

					if !strings.Contains(err.Error(), tt.wantErr) {
						t.Fatalf("got %q, want it to contain %q", err.Error(), tt.wantErr)
					}

					if got := errors.Is(err, ErrInstallationNotFound); got != tt.wantNotFound {
						t.Fatalf("errors.Is(%v, ErrInstallationNotFound): got %v, want %v", err, got, tt.wantNotFound)
					}

					if got != 0 {
						t.Fatalf("installation id: got %d, want nil", got)
					}

					if got := userCalls.Load(); got != tt.wantUserCalls {
						t.Fatalf("user lookups: got %d, want %d", got, tt.wantUserCalls)
					}
				})
			}
		})

		t.Run("errors_if_app_client_cannot_be_created", func(t *testing.T) {
			t.Parallel()

			opts := defaultSourceOptions()
			opts.appClientID = "123456789"
			opts.appSigner = mustReadAppSigner(t)
			opts.cacheEnabled = true
			opts.cacheBasePath = "\x00c"

			s, err := newAppSource(opts)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			got, err := s.installationID(t.Context(), "org")
			if err == nil {
				t.Fatal("got nil, want error")
			}

			if want := "failed to get app client"; !strings.Contains(err.Error(), want) {
				t.Fatalf("got %q, want it to contain %q", err.Error(), want)
			}

			if got != 0 {
				t.Fatalf("installation id: got %d, want nil", got)
			}
		})
	})
}

func Test_appTokenSource(t *testing.T) {
	t.Parallel()

	t.Run("succeeds", func(t *testing.T) {
		t.Parallel()

		for _, tt := range []struct {
			name           string
			installationID *int64
		}{
			{
				name:           "app",
				installationID: nil,
			},
			{
				name:           "installation",
				installationID: new(int64(42)),
			},
		} {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				got, err := appTokenSource("https://api.github.com/", "123456789", mustReadAppSigner(t), tt.installationID)
				if err != nil {
					t.Fatalf("unexpected error: %s", err)
				}

				if got == nil {
					t.Fatal("token source: got nil, want non-nil")
				}
			})
		}
	})

	t.Run("errors_with_non_rsa_signer", func(t *testing.T) {
		t.Parallel()

		key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			t.Fatalf("failed to generate key: %s", err)
		}

		got, err := appTokenSource("https://api.github.com/", "123456789", key, nil)
		if err == nil {
			t.Fatal("got nil, want error")
		}

		if want := "failed to create application token source"; !strings.Contains(err.Error(), want) {
			t.Fatalf("got %q, want it to contain %q", err.Error(), want)
		}

		if got != nil {
			t.Fatal("token source: got non-nil, want nil")
		}
	})
}
