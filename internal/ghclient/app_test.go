package ghclient

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
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
