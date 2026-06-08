package ghclient

import "testing"

func Test_newGraphQLClient(t *testing.T) {
	t.Parallel()

	t.Run("default_URL", func(t *testing.T) {
		t.Parallel()

		client, err := newGraphQLClient(nil, Options{})
		if err != nil {
			t.Fatalf("failed to create graphql client with default url: %v", err)
		}

		if client == nil {
			t.Fatal("expected graphql client to be non-nil")
		}
	})

	t.Run("enterprise_URL", func(t *testing.T) {
		t.Parallel()

		graphQLURL := "https://ghe.example.com/api/graphql"
		client, err := newGraphQLClient(nil, Options{GraphQLURL: &graphQLURL})
		if err != nil {
			t.Fatalf("failed to create graphql client with enterprise url: %v", err)
		}

		if client == nil {
			t.Fatal("expected graphql client to be non-nil")
		}
	})
}

func TestNewAnonymousGraphQLClient(t *testing.T) {
	t.Parallel()

	client, err := NewAnonymousGraphQLClient(Options{})
	if err != nil {
		t.Fatalf("failed to create anonymous graphql client: %v", err)
	}

	if client == nil {
		t.Fatal("expected anonymous graphql client to be non-nil")
	}
}

func TestNewAppGraphQLClient(t *testing.T) {
	t.Parallel()

	t.Run("invalid_private_key", func(t *testing.T) {
		t.Parallel()

		_, err := NewAppGraphQLClient("123456789", []byte("invalid-private-key"), nil, Options{})
		if err == nil {
			t.Fatal("expected app graphql client creation to fail for invalid private key")
		}
	})

	t.Run("valid_private_key", func(t *testing.T) {
		t.Parallel()

		privateKeyData := mustReadTestAppPrivateKey(t)

		client, err := NewAppGraphQLClient("123456789", privateKeyData, nil, Options{})
		if err != nil {
			t.Fatalf("failed to create app graphql client: %v", err)
		}

		if client == nil {
			t.Fatal("expected app graphql client to be non-nil")
		}
	})
}
