package ghclient

import (
	"testing"
)

func Test_newRESTClient(t *testing.T) {
	t.Parallel()

	client, err := newRESTClient(nil, testOptions(t))
	if err != nil {
		t.Fatalf("failed to create rest client: %v", err)
	}

	if client == nil {
		t.Fatal("expected rest client to be non-nil")
	}
}

func TestNewAnonymousRESTClient(t *testing.T) {
	t.Parallel()

	client, err := NewAnonymousRESTClient(testOptions(t))
	if err != nil {
		t.Fatalf("failed to create anonymous rest client: %v", err)
	}

	if client == nil {
		t.Fatal("expected anonymous rest client to be non-nil")
	}
}

func TestNewAppRESTClient(t *testing.T) {
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

			client, err := NewAppRESTClient("123456789", tt.privateKey, nil, testOptions(t))
			if tt.expectErr {
				if err == nil {
					t.Fatal("expected app rest client creation to fail for invalid private key")
				}

				return
			}

			if err != nil {
				t.Fatalf("failed to create app rest client: %v", err)
			}

			if client == nil {
				t.Fatal("expected app rest client to be non-nil")
			}
		})
	}
}
