package ghclient

import (
	"testing"
)

func Test_newRESTClient(t *testing.T) {
	t.Parallel()

	client, err := newRESTClient(nil, Options{})
	if err != nil {
		t.Fatalf("failed to create rest client: %v", err)
	}

	if client == nil {
		t.Fatal("expected rest client to be non-nil")
	}
}

func TestNewAnonymousRESTClient(t *testing.T) {
	t.Parallel()

	client, err := NewAnonymousRESTClient(Options{})
	if err != nil {
		t.Fatalf("failed to create anonymous rest client: %v", err)
	}

	if client == nil {
		t.Fatal("expected anonymous rest client to be non-nil")
	}
}

func TestNewAppRESTClient(t *testing.T) {
	t.Parallel()

	t.Run("invalid_private_key", func(t *testing.T) {
		t.Parallel()

		_, err := NewAppRESTClient("123456789", []byte("invalid-private-key"), nil, Options{})
		if err == nil {
			t.Fatal("expected app rest client creation to fail for invalid private key")
		}
	})

	t.Run("valid_private_key", func(t *testing.T) {
		t.Parallel()

		privateKeyData := mustReadTestAppPrivateKey(t)

		client, err := NewAppRESTClient("123456789", privateKeyData, nil, Options{})
		if err != nil {
			t.Fatalf("failed to create app rest client: %v", err)
		}

		if client == nil {
			t.Fatal("expected app rest client to be non-nil")
		}
	})
}
