package ghclient

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func TestNewAppSource(t *testing.T) {
	t.Parallel()

	privateKeyData := mustReadAppPrivateKey(t)

	cacheBasePath := mustMkdirTemp(t, "", "*")
	t.Cleanup(func() {
		_ = os.RemoveAll(cacheBasePath)
	})

	for _, tt := range []struct {
		name string
		opts Options
	}{
		{
			name: "default",
			opts: Options{
				RESTAPIURL: "https://api.github.com/",
				GraphQLURL: "https://api.github.com/graphql",
			},
		},
		{
			name: "with_cache_path",
			opts: Options{
				RESTAPIURL: "https://api.github.com/",
				GraphQLURL: "https://api.github.com/graphql",
				CachePath:  mustMkdirTemp(t, cacheBasePath, "*"),
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			source, err := NewAppSource("123456789", privateKeyData, tt.opts)
			if err != nil {
				t.Fatalf("failed to create app source: %v", err)
			}

			if source == nil {
				t.Fatal("expected app source to be non-nil")
			}
		})
	}
}

func Test_appSource(t *testing.T) {
	t.Parallel()

	owner1 := "octocat"
	owner2 := "acme"

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, fmt.Sprintf("/orgs/%s/installation", owner1)) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"id": 1000}`))
			return
		}

		if strings.HasSuffix(r.URL.Path, fmt.Sprintf("/orgs/%s/installation", owner2)) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"id": 1001}`))
			return
		}

		w.WriteHeader(http.StatusNotFound)
	})

	ts := httptest.NewServer(handler)
	t.Cleanup(ts.Close)

	cacheBasePath := mustMkdirTemp(t, "", "*")
	t.Cleanup(func() {
		_ = os.RemoveAll(cacheBasePath)
	})

	opts := Options{
		RESTAPIURL: ts.URL,
		GraphQLURL: ts.URL,
		CachePath:  mustMkdirTemp(t, cacheBasePath, "*"),
	}

	source, err := NewAppSource("123456789", mustReadAppPrivateKey(t), opts)
	if err != nil {
		t.Fatalf("failed to create app source: %v", err)
	}

	restClient, err := source.RESTClient()
	if err != nil {
		t.Fatalf("failed to get rest client: %v", err)
	}

	if restClient == nil {
		t.Fatal("expected rest client to be non-nil")
	}

	ownerRESTClientFirst, err := source.OwnerRESTClient(t.Context(), owner1)
	if err != nil {
		t.Fatalf("failed to get first owner rest client: %v", err)
	}

	if ownerRESTClientFirst == nil {
		t.Fatal("expected first owner rest client to be non-nil")
	}

	ownerRESTClientFirstAgain, err := source.OwnerRESTClient(t.Context(), owner1)
	if err != nil {
		t.Fatalf("failed to get first owner rest client again: %v", err)
	}

	if ownerRESTClientFirstAgain != ownerRESTClientFirst {
		t.Fatal("expected first owner rest client to be cached and reused")
	}

	ownerRESTClientSecond, err := source.OwnerRESTClient(t.Context(), owner2)
	if err != nil {
		t.Fatalf("failed to get second owner rest client: %v", err)
	}

	if ownerRESTClientSecond == nil {
		t.Fatal("expected second owner rest client to be non-nil")
	}

	if ownerRESTClientFirst == ownerRESTClientSecond {
		t.Fatal("expected different owner rest clients for different owners")
	}

	graphQLClient, err := source.GraphQLClient()
	if err != nil {
		t.Fatalf("failed to get graphql client: %v", err)
	}

	if graphQLClient == nil {
		t.Fatal("expected graphql client to be non-nil")
	}

	ownerGraphQLClientFirst, err := source.OwnerGraphQLClient(t.Context(), owner1)
	if err != nil {
		t.Fatalf("failed to get first owner graphql client: %v", err)
	}

	if ownerGraphQLClientFirst == nil {
		t.Fatal("expected first owner graphql client to be non-nil")
	}

	ownerGraphQLClientFirstAgain, err := source.OwnerGraphQLClient(t.Context(), owner1)
	if err != nil {
		t.Fatalf("failed to get first owner graphql client again: %v", err)
	}

	if ownerGraphQLClientFirstAgain != ownerGraphQLClientFirst {
		t.Fatal("expected first owner graphql client to be cached and reused")
	}

	ownerGraphQLClientSecond, err := source.OwnerGraphQLClient(t.Context(), owner2)
	if err != nil {
		t.Fatalf("failed to get second owner graphql client: %v", err)
	}

	if ownerGraphQLClientSecond == nil {
		t.Fatal("expected second owner graphql client to be non-nil")
	}

	if ownerGraphQLClientFirst == ownerGraphQLClientSecond {
		t.Fatal("expected different owner graphql clients for different owners")
	}
}

func Test_appSource_installationTokenRefresh(t *testing.T) {
	t.Parallel()

	const owner = "acme"
	const installationID = int64(1001)

	t.Run("reuses a valid token", func(t *testing.T) {
		t.Parallel()

		var tokenRequests atomic.Int32
		source := mustAppSourceForTokenTest(t, appSourceTokenHandler(&tokenRequests, owner, installationID, func(count int32) (string, time.Time) {
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
		source := mustAppSourceForTokenTest(t, appSourceTokenHandler(&tokenRequests, owner, installationID, func(count int32) (string, time.Time) {
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

func mustAppSourceForTokenTest(t *testing.T, handler http.Handler) *appSource {
	t.Helper()

	ts := httptest.NewServer(handler)
	t.Cleanup(ts.Close)

	source, err := NewAppSource("123456789", mustReadAppPrivateKey(t), Options{
		RESTAPIURL: ts.URL,
		GraphQLURL: ts.URL,
	})
	if err != nil {
		t.Fatalf("failed to create app source: %v", err)
	}

	return source
}
