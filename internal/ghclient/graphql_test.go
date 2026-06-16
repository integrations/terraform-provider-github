package ghclient

import "testing"

func Test_newGraphQLClient(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name       string
		graphQLURL string
	}{
		{
			name:       "default_URL",
			graphQLURL: "https://api.github.com/graphql",
		},
		{
			name:       "enterprise_URL",
			graphQLURL: "https://ghe.example.com/api/graphql",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			opts := testOptions(t)
			opts.GraphQLURL = tt.graphQLURL

			client, err := newGraphQLClient(nil, opts)
			if err != nil {
				t.Fatalf("failed to create graphql client: %v", err)
			}

			if client == nil {
				t.Fatal("expected graphql client to be non-nil")
			}
		})
	}
}

func TestNewAnonymousGraphQLClient(t *testing.T) {
	t.Parallel()

	client, err := NewAnonymousGraphQLClient(testOptions(t))
	if err != nil {
		t.Fatalf("failed to create anonymous graphql client: %v", err)
	}

	if client == nil {
		t.Fatal("expected anonymous graphql client to be non-nil")
	}
}

func TestNewAppGraphQLClient(t *testing.T) {
	t.Parallel()
	privateKeyData := mustReadTestAppPrivateKey(t)

	for _, tt := range []struct {
		name       string
		privateKey []byte
		expectErr  bool
	}{
		{
			name:       "invalid_private_key",
			privateKey: []byte("invalid-private-key"),
			expectErr:  true,
		},
		{
			name:       "valid_private_key",
			privateKey: privateKeyData,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			client, err := NewAppGraphQLClient("123456789", tt.privateKey, nil, testOptions(t))
			if tt.expectErr {
				if err == nil {
					t.Fatal("expected app graphql client creation to fail for invalid private key")
				}

				return
			}

			if err != nil {
				t.Fatalf("failed to create app graphql client: %v", err)
			}

			if client == nil {
				t.Fatal("expected app graphql client to be non-nil")
			}
		})
	}
}
