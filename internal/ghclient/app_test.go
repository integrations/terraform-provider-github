package ghclient

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func TestNewAppSource(t *testing.T) {
	t.Parallel()

	privateKeyData := mustReadTestAppPrivateKey(t)

	source, err := NewAppSource("123456789", privateKeyData, Options{})
	if err != nil {
		t.Fatalf("failed to create app source: %v", err)
	}

	if source == nil {
		t.Fatal("expected app source to be non-nil")
	}
}

func Test_appSource(t *testing.T) {
	t.Parallel()

	t.Run("RESTClient_cache", func(t *testing.T) {
		t.Parallel()

		var requestCount atomic.Int32
		source := mustTestAppSource(t, http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			requestCount.Add(1)
			w.WriteHeader(http.StatusInternalServerError)
		}))

		firstClient, err := source.RESTClient()
		if err != nil {
			t.Fatalf("failed to get first rest client: %v", err)
		}

		secondClient, err := source.RESTClient()
		if err != nil {
			t.Fatalf("failed to get second rest client: %v", err)
		}

		if firstClient != secondClient {
			t.Fatal("expected rest client to be cached")
		}

		if requestCount.Load() != 0 {
			t.Fatalf("expected no HTTP requests when building cached rest client, got %d", requestCount.Load())
		}
	})

	t.Run("OwnerClient_cache_by_owner", func(t *testing.T) {
		t.Parallel()

		for _, tt := range []struct {
			name       string
			callClient func(context.Context, *appSource, string) (any, error)
		}{
			{
				name: "rest",
				callClient: func(ctx context.Context, source *appSource, owner string) (any, error) {
					return source.OwnerRESTClient(ctx, owner)
				},
			},
			{
				name: "graphql",
				callClient: func(ctx context.Context, source *appSource, owner string) (any, error) {
					return source.OwnerGraphQLClient(ctx, owner)
				},
			},
		} {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var orgRequests atomic.Int32
				source := mustTestAppSource(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					if strings.HasSuffix(r.URL.Path, "/orgs/acme/installation") {
						orgRequests.Add(1)
						w.WriteHeader(http.StatusOK)
						_, _ = w.Write([]byte(`{"id": 1001}`))
						return
					}

					w.WriteHeader(http.StatusNotFound)
				}))

				firstClient, err := tt.callClient(t.Context(), source, "acme")
				if err != nil {
					t.Fatalf("failed to get first owner client: %v", err)
				}

				secondClient, err := tt.callClient(t.Context(), source, "acme")
				if err != nil {
					t.Fatalf("failed to get second owner client: %v", err)
				}

				if firstClient != secondClient {
					t.Fatal("expected owner client to be cached")
				}

				if orgRequests.Load() != 1 {
					t.Fatalf("expected one organization installation lookup, got %d", orgRequests.Load())
				}
			})
		}
	})

	t.Run("GetInstallationID_scenarios", func(t *testing.T) {
		t.Parallel()

		fallbackID := int64(2002)

		for _, tt := range []struct {
			name            string
			owner           string
			handleRequest   func(w http.ResponseWriter, r *http.Request, orgRequests, userRequests *atomic.Int32)
			expectError     bool
			errorContains   string
			expectedID      *int64
			expectedOrgReq  int32
			expectedUserReq int32
		}{
			{
				name:  "fallback_to_user",
				owner: "octocat",
				handleRequest: func(w http.ResponseWriter, r *http.Request, orgRequests, userRequests *atomic.Int32) {
					if strings.HasSuffix(r.URL.Path, "/orgs/octocat/installation") {
						orgRequests.Add(1)
						w.WriteHeader(http.StatusNotFound)
						return
					}
					if strings.HasSuffix(r.URL.Path, "/users/octocat/installation") {
						userRequests.Add(1)
						w.WriteHeader(http.StatusOK)
						_, _ = w.Write([]byte(`{"id": 2002}`))
						return
					}
					w.WriteHeader(http.StatusNotFound)
				},
				expectedID:      &fallbackID,
				expectedOrgReq:  1,
				expectedUserReq: 1,
			},
			{
				name:  "org_lookup_error",
				owner: "acme",
				handleRequest: func(w http.ResponseWriter, r *http.Request, orgRequests, _ *atomic.Int32) {
					if strings.HasSuffix(r.URL.Path, "/orgs/acme/installation") {
						orgRequests.Add(1)
						w.WriteHeader(http.StatusInternalServerError)
						_, _ = w.Write([]byte(`{"message": "boom"}`))
						return
					}
					w.WriteHeader(http.StatusNotFound)
				},
				expectError:    true,
				errorContains:  `failed to get installation for owner "acme"`,
				expectedOrgReq: 1,
			},
			{
				name:  "no_installation_id",
				owner: "acme",
				handleRequest: func(w http.ResponseWriter, r *http.Request, orgRequests, _ *atomic.Int32) {
					if strings.HasSuffix(r.URL.Path, "/orgs/acme/installation") {
						orgRequests.Add(1)
						w.WriteHeader(http.StatusOK)
						_, _ = w.Write([]byte(`{}`))
						return
					}
					w.WriteHeader(http.StatusNotFound)
				},
				expectError:    true,
				errorContains:  `no installation found for owner "acme"`,
				expectedOrgReq: 1,
			},
		} {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var orgRequests atomic.Int32
				var userRequests atomic.Int32
				source := mustTestAppSource(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					tt.handleRequest(w, r, &orgRequests, &userRequests)
				}))

				installationID, err := source.GetInstallationID(t.Context(), tt.owner)
				if tt.expectError {
					if err == nil {
						t.Fatal("expected error")
					}
					if tt.errorContains != "" && !strings.Contains(err.Error(), tt.errorContains) {
						t.Fatalf("expected error containing %q, got %v", tt.errorContains, err)
					}
				} else {
					if err != nil {
						t.Fatalf("expected success, got error: %v", err)
					}
					if tt.expectedID != nil && (installationID == nil || *installationID != *tt.expectedID) {
						t.Fatalf("expected installation id %d, got %v", *tt.expectedID, installationID)
					}
				}

				if orgRequests.Load() != tt.expectedOrgReq {
					t.Fatalf("expected %d org requests, got %d", tt.expectedOrgReq, orgRequests.Load())
				}
				if userRequests.Load() != tt.expectedUserReq {
					t.Fatalf("expected %d user requests, got %d", tt.expectedUserReq, userRequests.Load())
				}
			})
		}
	})
}

func Test_appSource_installationTokenRefresh(t *testing.T) {
	t.Parallel()

	const owner = "acme"
	const installationID = int64(1001)

	t.Run("reuses a valid token", func(t *testing.T) {
		t.Parallel()

		var tokenRequests atomic.Int32
		source := mustTestAppSource(t, appSourceTokenHandler(&tokenRequests, owner, installationID, func(count int32) (string, time.Time) {
			return "valid-token", time.Now().Add(time.Hour)
		}))

		client, err := source.OwnerRESTClient(t.Context(), owner)
		if err != nil {
			t.Fatalf("failed to get owner rest client: %v", err)
		}

		for i := range 2 {
			_, _, err := client.Organizations.Get(t.Context(), owner)
			if err != nil {
				t.Fatalf("request %d failed: %v", i+1, err)
			}
		}

		if got := tokenRequests.Load(); got != 1 {
			t.Fatalf("expected 1 installation token request, got %d", got)
		}
	})

	t.Run("refreshes an expired token", func(t *testing.T) {
		t.Parallel()

		var tokenRequests atomic.Int32
		source := mustTestAppSource(t, appSourceTokenHandler(&tokenRequests, owner, installationID, func(count int32) (string, time.Time) {
			if count == 1 {
				return "expired-token", time.Now().Add(-time.Hour)
			}
			return "refreshed-token", time.Now().Add(time.Hour)
		}))

		client, err := source.OwnerRESTClient(t.Context(), owner)
		if err != nil {
			t.Fatalf("failed to get owner rest client: %v", err)
		}

		for i := range 2 {
			_, _, err := client.Organizations.Get(t.Context(), owner)
			if err != nil {
				t.Fatalf("request %d failed: %v", i+1, err)
			}
		}

		if got := tokenRequests.Load(); got != 2 {
			t.Fatalf("expected 2 installation token requests, got %d", got)
		}
	})
}

func appSourceTokenHandler(tokenRequests *atomic.Int32, owner string, installationID int64, tokenFn func(count int32) (string, time.Time)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, fmt.Sprintf("/orgs/%s/installation", owner)) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = io.WriteString(w, fmt.Sprintf(`{"id": %d}`, installationID))
			return
		}

		if r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, fmt.Sprintf("/app/installations/%d/access_tokens", installationID)) {
			count := tokenRequests.Add(1)
			token, expiresAt := tokenFn(count)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_, _ = io.WriteString(w, fmt.Sprintf(`{"token":%q,"expires_at":%q}`, token, expiresAt.UTC().Format(time.RFC3339)))
			return
		}

		if r.Method == http.MethodGet && r.URL.Path == fmt.Sprintf("/orgs/%s", owner) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = io.WriteString(w, `{"id": 1, "login": "`+owner+`"}`)
			return
		}

		http.NotFound(w, r)
	}
}
