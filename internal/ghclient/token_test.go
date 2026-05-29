package ghclient

import (
	"context"
	"testing"
)

func TestNewTokenSource(t *testing.T) {
	t.Parallel()

	source, err := NewTokenSource("test-token", Options{})
	if err != nil {
		t.Fatalf("failed to create token source: %v", err)
	}

	if source == nil {
		t.Fatal("expected token source to be non-nil")
	}
}

func TestNewTokenRESTClient(t *testing.T) {
	t.Parallel()

	client, err := NewTokenRESTClient("test-token", Options{})
	if err != nil {
		t.Fatalf("failed to create token rest client: %v", err)
	}

	if client == nil {
		t.Fatal("expected token rest client to be non-nil")
	}
}

func TestNewTokenGraphQLClient(t *testing.T) {
	t.Parallel()

	client, err := NewTokenGraphQLClient("test-token", Options{})
	if err != nil {
		t.Fatalf("failed to create token graphql client: %v", err)
	}

	if client == nil {
		t.Fatal("expected token graphql client to be non-nil")
	}
}

func Test_tokenSource(t *testing.T) {
	t.Parallel()

	t.Run("RESTClient", func(t *testing.T) {
		t.Parallel()

		source, err := NewTokenSource("test-token", Options{})
		if err != nil {
			t.Fatalf("failed to create token source: %v", err)
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

		source, err := NewTokenSource("test-token", Options{})
		if err != nil {
			t.Fatalf("failed to create token source: %v", err)
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

		source, err := NewTokenSource("test-token", Options{})
		if err != nil {
			t.Fatalf("failed to create token source: %v", err)
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

		source, err := NewTokenSource("test-token", Options{})
		if err != nil {
			t.Fatalf("failed to create token source: %v", err)
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
