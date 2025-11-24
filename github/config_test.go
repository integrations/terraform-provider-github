package github

import (
	"context"
	"net/url"
	"testing"

	"github.com/shurcooL/githubv4"
)

func TestGHECDataResidencyHostMatch(t *testing.T) {
	testCases := []struct {
		url         string
		matches     bool
		description string
	}{
		{
			url:         "https://customer.ghe.com/",
			matches:     true,
			description: "GHEC data residency URL with customer name",
		},
		{
			url:         "https://customer-name.ghe.com/",
			matches:     true,
			description: "GHEC data residency URL with hyphenated name",
		},
		{
			url:         "https://customer.ghe.com",
			matches:     true,
			description: "GHEC data residency URL without a trailing slash",
		},
		{
			url:         "https://ghe.com/",
			matches:     false,
			description: "GHEC domain without subdomain",
		},
		{
			url:         "https://github.com/",
			matches:     false,
			description: "GitHub.com URL",
		},
		{
			url:         "https://api.github.com/",
			matches:     false,
			description: "GitHub.com API URL",
		},
		{
			url:         "https://example.com/",
			matches:     false,
			description: "Generic URL",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			u, err := url.Parse(tc.url)
			if err != nil {
				t.Fatalf("failed to parse URL %q: %s", tc.url, err)
			}
			matches := GHECDataResidencyHostMatch.MatchString(u.Hostname())
			if matches != tc.matches {
				t.Errorf("URL %q: expected match=%v, got %v", tc.url, tc.matches, matches)
			}
		})
	}
}

func TestAccConfigMeta(t *testing.T) {
	t.Run("returns an anonymous client for the v3 REST API", func(t *testing.T) {
		config := Config{BaseURL: "https://api.github.com/"}
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
		skipUnlessMode(t, individual)

		config := Config{
			Token:   testAccConf.token,
			BaseURL: "https://api.github.com/",
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
		skipUnlessMode(t, individual)

		config := Config{
			Token:   testAccConf.token,
			BaseURL: "https://api.github.com/",
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
		skipUnlessMode(t, individual)

		config := Config{
			Token:   testAccConf.token,
			BaseURL: "https://api.github.com/",
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
		skipUnlessHasOrgs(t)

		config := Config{
			Token:   testAccConf.token,
			BaseURL: "https://api.github.com/",
			Owner:   testAccConf.owner,
		}
		meta, err := config.Meta()
		if err != nil {
			t.Fatalf("failed to return meta without error: %s", err.Error())
		}

		ctx := context.Background()
		client := meta.(*Owner).v3client
		_, _, err = client.Organizations.Get(ctx, testAccConf.owner)
		if err != nil {
			t.Fatalf("failed to validate returned client without error: %s", err.Error())
		}
	})

	t.Run("returns a v4 GraphQL API client to manage organization resources", func(t *testing.T) {
		skipUnlessHasOrgs(t)

		config := Config{
			Token:   testAccConf.token,
			BaseURL: "https://api.github.com/",
			Owner:   testAccConf.owner,
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
			"login": githubv4.String(testAccConf.owner),
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
