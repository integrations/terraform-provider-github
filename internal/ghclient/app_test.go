package ghclient

import (
	"context"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
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

		testCases := []struct {
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
		}

		for _, tc := range testCases {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
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

				firstClient, err := tc.callClient(context.Background(), source, "acme")
				if err != nil {
					t.Fatalf("failed to get first owner client: %v", err)
				}

				secondClient, err := tc.callClient(context.Background(), source, "acme")
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

		testCases := []struct {
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
		}

		for _, tc := range testCases {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				var orgRequests atomic.Int32
				var userRequests atomic.Int32
				source := mustTestAppSource(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					tc.handleRequest(w, r, &orgRequests, &userRequests)
				}))

				installationID, err := source.GetInstallationID(context.Background(), tc.owner)
				if tc.expectError {
					if err == nil {
						t.Fatal("expected error")
					}
					if tc.errorContains != "" && !strings.Contains(err.Error(), tc.errorContains) {
						t.Fatalf("expected error containing %q, got %v", tc.errorContains, err)
					}
				} else {
					if err != nil {
						t.Fatalf("expected success, got error: %v", err)
					}
					if tc.expectedID != nil && (installationID == nil || *installationID != *tc.expectedID) {
						t.Fatalf("expected installation id %d, got %v", *tc.expectedID, installationID)
					}
				}

				if orgRequests.Load() != tc.expectedOrgReq {
					t.Fatalf("expected %d org requests, got %d", tc.expectedOrgReq, orgRequests.Load())
				}
				if userRequests.Load() != tc.expectedUserReq {
					t.Fatalf("expected %d user requests, got %d", tc.expectedUserReq, userRequests.Load())
				}
			})
		}
	})
}
