package ghclient

import (
	"regexp"
	"testing"
)

func TestClientOptions_getRESTURL(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name    string
		opts    ClientOptions
		wantURL *string
		wantErr string
	}{
		{
			name:    "default",
			opts:    ClientOptions{},
			wantURL: new("https://api.github.com/"),
		},
		{
			name: "dotcom",
			opts: ClientOptions{
				BaseURL: "https://api.github.com/",
			},
			wantURL: new("https://api.github.com/"),
		},
		{
			name: "ghec",
			opts: ClientOptions{
				BaseURL: "https://api.my-enterprise.ghe.com/",
			},
			wantURL: new("https://api.my-enterprise.ghe.com/"),
		},
		{
			name: "ghes",
			opts: ClientOptions{
				BaseURL: "https://github.example.com/",
				IsGHES:  true,
			},
			wantURL: new("https://github.example.com/api/v3/"),
		},
		{
			name: "ghes_with_path",
			opts: ClientOptions{
				BaseURL: "https://github.example.com/some/path/",
				IsGHES:  true,
			},
			wantURL: new("https://github.example.com/some/path/api/v3/"),
		},
		{
			name: "invalid_url",
			opts: ClientOptions{
				BaseURL: "https://api.github.com/%%%",
			},
			wantErr: "unable to parse base url",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotURL, err := tt.opts.getRESTURL()
			if err != nil {
				if tt.wantErr == "" {
					t.Fatalf("unexpected error: %v", err)
				}
				if !regexp.MustCompile(regexp.QuoteMeta(tt.wantErr)).MatchString(err.Error()) {
					t.Fatalf("expected error to match %q, got %v", tt.wantErr, err)
				}
				return
			}

			if tt.wantErr != "" {
				t.Fatalf("expected error %q, got nil", tt.wantErr)
			}

			if gotURL == nil && tt.wantURL != nil {
				t.Fatalf("expected URL to be %q, got nil", *tt.wantURL)
			}

			if gotURL != nil && tt.wantURL == nil {
				t.Fatalf("expected URL to be nil, got %q", *gotURL)
			}

			if *gotURL != *tt.wantURL {
				t.Fatalf("expected URL to be %q, got %q", *tt.wantURL, *gotURL)
			}
		})
	}
}

func TestClientOptions_getGraphQLURL(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name    string
		opts    ClientOptions
		wantURL *string
		wantErr string
	}{
		{
			name:    "default",
			opts:    ClientOptions{},
			wantURL: new("https://api.github.com/graphql"),
		},
		{
			name: "dotcom",
			opts: ClientOptions{
				BaseURL: "https://api.github.com/",
			},
			wantURL: new("https://api.github.com/graphql"),
		},
		{
			name: "ghec",
			opts: ClientOptions{
				BaseURL: "https://api.my-enterprise.ghe.com/",
			},
			wantURL: new("https://api.my-enterprise.ghe.com/graphql"),
		},
		{
			name: "ghes",
			opts: ClientOptions{
				BaseURL: "https://github.example.com/",
				IsGHES:  true,
			},
			wantURL: new("https://github.example.com/api/graphql"),
		},
		{
			name: "ghes_with_path",
			opts: ClientOptions{
				BaseURL: "https://github.example.com/some/path/",
				IsGHES:  true,
			},
			wantURL: new("https://github.example.com/some/path/api/graphql"),
		},
		{
			name: "invalid_url",
			opts: ClientOptions{
				BaseURL: "https://api.github.com/%%%",
			},
			wantErr: "unable to parse base url",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotURL, err := tt.opts.getGraphQLURL()
			if err != nil {
				if tt.wantErr == "" {
					t.Fatalf("unexpected error: %v", err)
				}
				if !regexp.MustCompile(regexp.QuoteMeta(tt.wantErr)).MatchString(err.Error()) {
					t.Fatalf("expected error to match %q, got %v", tt.wantErr, err)
				}
				return
			}

			if tt.wantErr != "" {
				t.Fatalf("expected error %q, got nil", tt.wantErr)
			}

			if gotURL == nil && tt.wantURL != nil {
				t.Fatalf("expected URL to be %q, got nil", *tt.wantURL)
			}

			if gotURL != nil && tt.wantURL == nil {
				t.Fatalf("expected URL to be nil, got %q", *gotURL)
			}

			if *gotURL != *tt.wantURL {
				t.Fatalf("expected URL to be %q, got %q", *tt.wantURL, *gotURL)
			}
		})
	}
}
