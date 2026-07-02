package ghclient

import (
	"os"
	"testing"

	"github.com/google/go-github/v88/github"
	"github.com/shurcooL/githubv4"
)

func TestNewAnonymousSource(t *testing.T) {
	t.Parallel()

	cacheBasePath := mustMkdirTemp(t, "", "*")
	t.Cleanup(func() {
		_ = os.RemoveAll(cacheBasePath)
	})

	for _, tt := range []struct {
		name string
		opts SourceOptions
	}{
		{
			name: "default",
			opts: SourceOptions{},
		},
		{
			name: "with_cache_base_path",
			opts: SourceOptions{
				Cache:         true,
				CacheBasePath: mustMkdirTemp(t, cacheBasePath, "*"),
			},
		},
		{
			name: "with_cache_no_base_path",
			opts: SourceOptions{
				Cache: true,
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			source, err := NewAnonymousSource(tt.opts)
			if err != nil {
				t.Fatalf("failed to create anonymous source: %v", err)
			}

			if source == nil {
				t.Fatal("expected anonymous source to be non-nil")
			}
		})
	}
}

func Test_anonymousSource(t *testing.T) {
	t.Parallel()

	source := &anonymousSource{
		restClient:    &github.Client{},
		graphQLClient: &githubv4.Client{},
	}

	restClient, err := source.RESTClient()
	if err != nil {
		t.Fatalf("failed to get rest client: %v", err)
	}

	if restClient == nil {
		t.Fatal("expected rest client to be non-nil")
	}

	ownerRESTClient, err := source.OwnerRESTClient(t.Context(), "octocat")
	if err != nil {
		t.Fatalf("failed to get owner rest client: %v", err)
	}

	if ownerRESTClient != restClient {
		t.Fatal("expected owner rest client to be the same client as default rest client")
	}

	graphQLClient, err := source.GraphQLClient()
	if err != nil {
		t.Fatalf("failed to get graphql client: %v", err)
	}

	if graphQLClient == nil {
		t.Fatal("expected graphql client to be non-nil")
	}

	ownerGraphQLClient, err := source.OwnerGraphQLClient(t.Context(), "octocat")
	if err != nil {
		t.Fatalf("failed to get owner graphql client: %v", err)
	}

	if ownerGraphQLClient != graphQLClient {
		t.Fatal("expected owner graphql client to be the same client as default graphql client")
	}
}
