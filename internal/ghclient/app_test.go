package ghclient

import (
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
)

func TestNewAppSource(t *testing.T) {
	t.Parallel()

	privateKeyData := mustReadTestAppPrivateKey(t)

	source, err := NewAppSource("123456789", privateKeyData, testOptions(t))
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

		firstRESTClient, err := source.OwnerRESTClient(t.Context(), "acme")
		if err != nil {
			t.Fatalf("failed to get first owner rest client: %v", err)
		}

		secondRESTClient, err := source.OwnerRESTClient(t.Context(), "acme")
		if err != nil {
			t.Fatalf("failed to get second owner rest client: %v", err)
		}

		if firstRESTClient != secondRESTClient {
			t.Fatal("expected owner rest client to be cached")
		}

		if orgRequests.Load() != 1 {
			t.Fatalf("expected one organization installation lookup after rest client requests, got %d", orgRequests.Load())
		}

		firstGraphQLClient, err := source.OwnerGraphQLClient(t.Context(), "acme")
		if err != nil {
			t.Fatalf("failed to get first owner graphql client: %v", err)
		}

		secondGraphQLClient, err := source.OwnerGraphQLClient(t.Context(), "acme")
		if err != nil {
			t.Fatalf("failed to get second owner graphql client: %v", err)
		}

		if firstGraphQLClient != secondGraphQLClient {
			t.Fatal("expected owner graphql client to be cached")
		}

		if orgRequests.Load() != 2 {
			t.Fatalf("expected one organization installation lookup per client type, got %d", orgRequests.Load())
		}
	})

	t.Run("GetInstallationID_scenarios", func(t *testing.T) {
		t.Parallel()

		fallbackID := int64(2002)

		t.Run("fallback_to_user", func(t *testing.T) {
			t.Parallel()

			var orgRequests atomic.Int32
			var userRequests atomic.Int32
			source := mustTestAppSource(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
			}))

			installationID, err := source.GetInstallationID(t.Context(), "octocat")
			if err != nil {
				t.Fatalf("expected success, got error: %v", err)
			}

			if installationID == nil || *installationID != fallbackID {
				t.Fatalf("expected installation id %d, got %v", fallbackID, installationID)
			}

			if orgRequests.Load() != 1 {
				t.Fatalf("expected 1 org request, got %d", orgRequests.Load())
			}

			if userRequests.Load() != 1 {
				t.Fatalf("expected 1 user request, got %d", userRequests.Load())
			}
		})

		t.Run("org_lookup_error", func(t *testing.T) {
			t.Parallel()

			var orgRequests atomic.Int32
			var userRequests atomic.Int32
			source := mustTestAppSource(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if strings.HasSuffix(r.URL.Path, "/orgs/acme/installation") {
					orgRequests.Add(1)
					w.WriteHeader(http.StatusInternalServerError)
					_, _ = w.Write([]byte(`{"message": "boom"}`))
					return
				}
				w.WriteHeader(http.StatusNotFound)
			}))

			_, err := source.GetInstallationID(t.Context(), "acme")
			if err == nil {
				t.Fatal("expected error")
			}

			if !strings.Contains(err.Error(), `failed to get installation for owner "acme"`) {
				t.Fatalf("expected installation lookup error, got %v", err)
			}

			if orgRequests.Load() != 1 {
				t.Fatalf("expected 1 org request, got %d", orgRequests.Load())
			}

			if userRequests.Load() != 0 {
				t.Fatalf("expected 0 user requests, got %d", userRequests.Load())
			}
		})

		t.Run("no_installation_id", func(t *testing.T) {
			t.Parallel()

			var orgRequests atomic.Int32
			var userRequests atomic.Int32
			source := mustTestAppSource(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if strings.HasSuffix(r.URL.Path, "/orgs/acme/installation") {
					orgRequests.Add(1)
					w.WriteHeader(http.StatusOK)
					_, _ = w.Write([]byte(`{}`))
					return
				}
				w.WriteHeader(http.StatusNotFound)
			}))

			_, err := source.GetInstallationID(t.Context(), "acme")
			if err == nil {
				t.Fatal("expected error")
			}

			if !strings.Contains(err.Error(), `no installation found for owner "acme"`) {
				t.Fatalf("expected missing installation error, got %v", err)
			}

			if orgRequests.Load() != 1 {
				t.Fatalf("expected 1 org request, got %d", orgRequests.Load())
			}

			if userRequests.Load() != 0 {
				t.Fatalf("expected 0 user requests, got %d", userRequests.Load())
			}
		})
	})
}
