package ghclient

import (
	"regexp"
	"testing"

	"golang.org/x/oauth2"
)

func TestNewAnonymousGraphQLClient(t *testing.T) {
	t.Parallel()

	opts := ClientOptions{}

	client, err := NewAnonymousGraphQLClient(opts)
	if err != nil {
		t.Fatalf("failed to create anonymous graphql client: %v", err)
	}

	if client == nil {
		t.Fatal("expected anonymous graphql client to be non-nil")
	}
}

func TestNewAppGraphQLClient(t *testing.T) {
	t.Parallel()

	privateKeyData := mustReadAppPrivateKey(t)

	for _, tt := range []struct {
		name           string
		privateKey     []byte
		installationID *int64
		opts           ClientOptions
		wantErr        string
	}{
		{
			name:           "app_client",
			privateKey:     privateKeyData,
			installationID: nil,
			opts:           ClientOptions{},
		},
		{
			name:           "installation_client",
			privateKey:     privateKeyData,
			installationID: new(int64(8888)),
			opts:           ClientOptions{},
		},
		{
			name:           "handles_token_source_error",
			privateKey:     []byte("invalid-private-key"),
			installationID: nil,
			opts:           ClientOptions{},
			wantErr:        "failed to create app token source",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			client, err := NewAppGraphQLClient("123456789", tt.privateKey, tt.installationID, tt.opts)
			if err != nil {
				if tt.wantErr == "" {
					t.Fatalf("failed to create app graphql client: %v", err)
				}

				if !regexp.MustCompile(regexp.QuoteMeta(tt.wantErr)).MatchString(err.Error()) {
					t.Fatalf("expected error to match %q, got %v", tt.wantErr, err)
				}

				return
			}

			if tt.wantErr != "" {
				t.Fatalf("expected error %q, got nil", tt.wantErr)
			}

			if client == nil {
				t.Fatal("expected app graphql client to be non-nil")
			}
		})
	}
}

func TestNewTokenGraphQLClient(t *testing.T) {
	t.Parallel()

	opts := ClientOptions{}

	client, err := NewTokenGraphQLClient("test-token", opts)
	if err != nil {
		t.Fatalf("failed to create token graphql client: %v", err)
	}

	if client == nil {
		t.Fatal("expected token graphql client to be non-nil")
	}
}

func Test_newGraphQLClient(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name        string
		tokenSource oauth2.TokenSource
		opts        ClientOptions
		wantErr     string
	}{
		{
			name:        "minimal",
			tokenSource: nil,
			opts:        ClientOptions{},
		},
		{
			name:        "with_token_source",
			tokenSource: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "test-token"}),
			opts:        ClientOptions{},
		},
		{
			name:        "with_url",
			tokenSource: nil,
			opts:        ClientOptions{BaseURL: "https://api.github.com/"},
		},
		{
			name:        "errors_if_transport_cannot_be_created",
			tokenSource: nil,
			opts:        ClientOptions{Cache: true, CachePath: "\x00c"},
			wantErr:     "failed to create transport",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := newGraphQLClient(tt.tokenSource, tt.opts)
			if err != nil {
				if tt.wantErr == "" {
					t.Fatalf("failed to create graphql client: %v", err)
				}

				if !regexp.MustCompile(regexp.QuoteMeta(tt.wantErr)).MatchString(err.Error()) {
					t.Fatalf("expected error to match %q, got %v", tt.wantErr, err)
				}

				return
			}

			if tt.wantErr != "" {
				t.Fatalf("expected error %q, got nil", tt.wantErr)
			}

			if got == nil {
				t.Fatal("expected graphql client to be non-nil")
			}
		})
	}
}
