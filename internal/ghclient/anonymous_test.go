package ghclient

import (
	"context"
	"testing"
)

func TestNewAnonymousSource(t *testing.T) {
	t.Parallel()

	source, err := NewAnonymousSource(Options{})
	if err != nil {
		t.Fatalf("failed to create anonymous source: %v", err)
	}

	if source == nil {
		t.Fatal("expected anonymous source to be non-nil")
	}
}

func Test_anonymousSource(t *testing.T) {
	t.Parallel()

	t.Run("RESTClient", func(t *testing.T) {
		t.Parallel()

		source, err := NewAnonymousSource(Options{})
		if err != nil {
			t.Fatalf("failed to create anonymous source: %v", err)
		}

		client, err := source.RESTClient()
		if err != nil {
			t.Fatalf("failed to get rest client: %v", err)
		}

		if client == nil {
			t.Fatal("expected rest client to be non-nil")
		}
	})

	t.Run("OwnerRESTClient", func(t *testing.T) {
		t.Parallel()

		source, err := NewAnonymousSource(Options{})
		if err != nil {
			t.Fatalf("failed to create anonymous source: %v", err)
		}

		defaultClient, err := source.RESTClient()
		if err != nil {
			t.Fatalf("failed to get default rest client: %v", err)
		}

		ownerClient, err := source.OwnerRESTClient(context.Background(), "octocat")
		if err != nil {
			t.Fatalf("failed to get owner rest client: %v", err)
		}

		if ownerClient != defaultClient {
			t.Fatal("expected owner rest client to be the same client as default rest client")
		}
	})

	t.Run("GraphQLClient", func(t *testing.T) {
		t.Parallel()

		source, err := NewAnonymousSource(Options{})
		if err != nil {
			t.Fatalf("failed to create anonymous source: %v", err)
		}

		client, err := source.GraphQLClient()
		if err != nil {
			t.Fatalf("failed to get graphql client: %v", err)
		}

		if client == nil {
			t.Fatal("expected graphql client to be non-nil")
		}
	})

	t.Run("OwnerGraphQLClient", func(t *testing.T) {
		t.Parallel()

		source, err := NewAnonymousSource(Options{})
		if err != nil {
			t.Fatalf("failed to create anonymous source: %v", err)
		}

		defaultClient, err := source.GraphQLClient()
		if err != nil {
			t.Fatalf("failed to get default graphql client: %v", err)
		}

		ownerClient, err := source.OwnerGraphQLClient(context.Background(), "octocat")
		if err != nil {
			t.Fatalf("failed to get owner graphql client: %v", err)
		}

		if ownerClient != defaultClient {
			t.Fatal("expected owner graphql client to be the same client as default graphql client")
		}
	})
}
