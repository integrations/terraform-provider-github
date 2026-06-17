package ghclient

import (
	"regexp"
	"testing"

	"golang.org/x/oauth2"
)

func TestNewAnonymousRESTClient(t *testing.T) {
	t.Parallel()

	opts := Options{
		RESTAPIURL: "https://api.github.com/",
	}

	client, err := NewAnonymousRESTClient(opts)
	if err != nil {
		t.Fatalf("failed to create anonymous rest client: %v", err)
	}

	if client == nil {
		t.Fatal("expected anonymous rest client to be non-nil")
	}
}

func TestNewAppRESTClient(t *testing.T) {
	t.Parallel()

	privateKeyData := mustReadAppPrivateKey(t)

	for _, tt := range []struct {
		name           string
		privateKey     []byte
		installationID *int64
		opts           Options
		wantErr        string
	}{
		{
			name:           "app_client",
			privateKey:     privateKeyData,
			installationID: nil,
			opts:           Options{RESTAPIURL: "https://api.github.com/"},
		},
		{
			name:           "installation_client",
			privateKey:     privateKeyData,
			installationID: new(int64(8888)),
			opts:           Options{RESTAPIURL: "https://api.github.com/"},
		},
		{
			name:           "handles_token_source_error",
			privateKey:     []byte("invalid-private-key"),
			installationID: nil,
			opts:           Options{RESTAPIURL: "https://api.github.com/"},
			wantErr:        "failed to create app token source",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			client, err := NewAppRESTClient("123456789", tt.privateKey, tt.installationID, tt.opts)
			if err != nil {
				if tt.wantErr == "" {
					t.Fatalf("failed to create app rest client: %v", err)
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
				t.Fatal("expected app rest client to be non-nil")
			}
		})
	}
}

func TestNewTokenRESTClient(t *testing.T) {
	t.Parallel()

	opts := Options{
		RESTAPIURL: "https://api.github.com/",
	}

	client, err := NewTokenRESTClient("test-token", opts)
	if err != nil {
		t.Fatalf("failed to create token rest client: %v", err)
	}

	if client == nil {
		t.Fatal("expected token rest client to be non-nil")
	}
}

func Test_newRESTClient(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name        string
		tokenSource oauth2.TokenSource
		opts        Options
		wantErr     string
	}{
		{
			name:        "with_token_source",
			tokenSource: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "test-token"}),
			opts:        Options{RESTAPIURL: "https://api.github.com/"},
		},
		{
			name:        "without_token_source",
			tokenSource: nil,
			opts:        Options{RESTAPIURL: "https://api.github.com/"},
		},
		{
			name:        "errors_if_transport_cannot_be_created",
			tokenSource: nil,
			opts:        Options{RESTAPIURL: "https://api.github.com/", CachePath: "\x00c"},
			wantErr:     "failed to create transport",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := newRESTClient(tt.tokenSource, tt.opts)
			if err != nil {
				if tt.wantErr == "" {
					t.Fatalf("failed to create rest client: %v", err)
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
				t.Fatal("expected rest client to be non-nil")
			}
		})
	}
}
