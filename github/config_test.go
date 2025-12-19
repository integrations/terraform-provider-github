package github

import (
	"context"
	"testing"

	"github.com/shurcooL/githubv4"
)

func Test_getBaseURL(t *testing.T) {
	testCases := []struct {
		name        string
		url         string
		expectedURL string
		isGHES      bool
		errors      bool
	}{
		{
			name:        "dotcom",
			url:         "https://api.github.com/",
			expectedURL: "https://api.github.com/",
			isGHES:      false,
			errors:      false,
		},
		{
			name:        "dotcom no trailing slash",
			url:         "https://api.github.com",
			expectedURL: "https://api.github.com/",
			isGHES:      false,
			errors:      false,
		},
		{
			name:        "dotcom ui",
			url:         "https://github.com/",
			expectedURL: "https://api.github.com/",
			isGHES:      false,
			errors:      false,
		},
		{
			name:        "dotcom http errors",
			url:         "http://api.github.com/",
			expectedURL: "",
			isGHES:      false,
			errors:      true,
		},
		{
			name:        "dotcom with path errors",
			url:         "https://api.github.com/xxx/",
			expectedURL: "",
			isGHES:      false,
			errors:      true,
		},
		{
			name:        "ghec",
			url:         "https://api.customer.ghe.com/",
			expectedURL: "https://api.customer.ghe.com/",
			isGHES:      false,
			errors:      false,
		},
		{
			name:        "ghec no trailing slash",
			url:         "https://api.customer.ghe.com",
			expectedURL: "https://api.customer.ghe.com/",
			isGHES:      false,
			errors:      false,
		},
		{
			name:        "ghec ui",
			url:         "https://customer.ghe.com/",
			expectedURL: "https://api.customer.ghe.com/",
			isGHES:      false,
			errors:      false,
		},
		{
			name:        "ghec http errors",
			url:         "http://api.customer.ghe.com/",
			expectedURL: "",
			isGHES:      false,
			errors:      true,
		},
		{
			name:        "ghec with path errors",
			url:         "https://api.customer.ghe.com/xxx/",
			expectedURL: "",
			isGHES:      false,
			errors:      true,
		},
		{
			name:        "ghes",
			url:         "https://example.com/",
			expectedURL: "https://example.com/",
			isGHES:      true,
			errors:      false,
		},
		{
			name:        "ghes no trailing slash",
			url:         "https://example.com",
			expectedURL: "https://example.com/",
			isGHES:      true,
			errors:      false,
		},
		{
			name:        "ghes with path prefix",
			url:         "https://example.com/test/",
			expectedURL: "https://example.com/test/",
			isGHES:      true,
			errors:      false,
		},
		{
			name:        "empty url returns dotcom",
			url:         "",
			expectedURL: "https://api.github.com/",
			isGHES:      false,
			errors:      false,
		},
		{
			name:        "not absolute url errors",
			url:         "example.com/",
			expectedURL: "",
			isGHES:      false,
			errors:      true,
		},
		{
			name:        "invalid url errors",
			url:         "xxx",
			expectedURL: "",
			isGHES:      false,
			errors:      true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			baseURL, isGHES, err := getBaseURL(tc.url)
			if err != nil {
				if tc.errors {
					return
				}
				t.Fatalf("expected no error, got: %v", err)
			}

			if tc.errors {
				t.Fatalf("expected error, got none")
			}

			if baseURL.String() != tc.expectedURL {
				t.Errorf("expected base URL %q, got %q", tc.expectedURL, baseURL.String())
			}

			if isGHES != tc.isGHES {
				t.Errorf("expected isGHES to be %v, got %v", tc.isGHES, isGHES)
			}
		})
	}
}

func TestAccConfigMeta(t *testing.T) {
	// FIXME: Skip test runs during travis lint checking
	if testToken == "" {
		return
	}

	baseURL, _, err := getBaseURL(DotComAPIURL)
	if err != nil {
		t.Fatalf("failed to parse test base URL: %s", err.Error())
	}

	t.Run("returns an anonymous client for the v3 REST API", func(t *testing.T) {
		config := Config{BaseURL: baseURL}
		meta, err := config.Meta()
		if err != nil {
			t.Fatalf("failed to return meta without error: %s", err.Error())
		}

		ctx := context.Background()
		client := meta.(*Owner).v3client
		_, _, err = client.Meta.Get(ctx)
		if err != nil {
			t.Fatalf("failed to validate returned client without error: %s", err.Error())
		}
	})

	t.Run("returns a v3 REST API client to manage individual resources", func(t *testing.T) {
		config := Config{
			Token:   testToken,
			BaseURL: baseURL,
		}
		meta, err := config.Meta()
		if err != nil {
			t.Fatalf("failed to return meta without error: %s", err.Error())
		}

		ctx := context.Background()
		client := meta.(*Owner).v3client
		_, _, err = client.Meta.Get(ctx)
		if err != nil {
			t.Fatalf("failed to validate returned client without error: %s", err.Error())
		}
	})

	t.Run("returns a v3 REST API client with max retries", func(t *testing.T) {
		config := Config{
			Token:   testToken,
			BaseURL: baseURL,
			RetryableErrors: map[int]bool{
				500: true,
				502: true,
			},
			MaxRetries: 3,
		}
		meta, err := config.Meta()
		if err != nil {
			t.Fatalf("failed to return meta without error: %s", err.Error())
		}

		ctx := context.Background()
		client := meta.(*Owner).v3client
		_, _, err = client.Meta.Get(ctx)
		if err != nil {
			t.Fatalf("failed to validate returned client without error: %s", err.Error())
		}
	})

	t.Run("returns a v4 GraphQL API client to manage individual resources", func(t *testing.T) {
		config := Config{
			Token:   testToken,
			BaseURL: baseURL,
		}
		meta, err := config.Meta()
		if err != nil {
			t.Fatalf("failed to return meta without error: %s", err.Error())
		}

		client := meta.(*Owner).v4client
		var query struct {
			Meta struct {
				GitHubServicesSha githubv4.String
			}
		}
		err = client.Query(context.Background(), &query, nil)
		if err != nil {
			t.Fatalf("failed to validate returned client without error: %s", err.Error())
		}
	})

	t.Run("returns a v3 REST API client to manage organization resources", func(t *testing.T) {
		config := Config{
			Token:   testToken,
			BaseURL: baseURL,
			Owner:   testOrganization,
		}
		meta, err := config.Meta()
		if err != nil {
			t.Fatalf("failed to return meta without error: %s", err.Error())
		}

		ctx := context.Background()
		client := meta.(*Owner).v3client
		_, _, err = client.Organizations.Get(ctx, testOrganization)
		if err != nil {
			t.Fatalf("failed to validate returned client without error: %s", err.Error())
		}
	})

	t.Run("returns a v4 GraphQL API client to manage organization resources", func(t *testing.T) {
		config := Config{
			Token:   testToken,
			BaseURL: baseURL,
			Owner:   testOrganization,
		}
		meta, err := config.Meta()
		if err != nil {
			t.Fatalf("failed to return meta without error: %s", err.Error())
		}

		client := meta.(*Owner).v4client

		var query struct {
			Organization struct {
				ViewerCanAdminister githubv4.Boolean
			} `graphql:"organization(login: $login)"`
		}
		variables := map[string]any{
			"login": githubv4.String(testOrganization),
		}
		err = client.Query(context.Background(), &query, variables)
		if err != nil {
			t.Fatalf("failed to validate returned client without error: %s", err.Error())
		}

		if query.Organization.ViewerCanAdminister != true {
			t.Fatalf("unexpected response when validating client")
		}
	})
}
