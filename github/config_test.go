package github

import (
	"context"
	"testing"

	"github.com/shurcooL/githubv4"
)

func TestAccConfigMeta(t *testing.T) {

	// FIXME: Skip test runs during travis lint checking
	if testToken == "" {
		return
	}

	t.Run("returns an anonymous client for the v3 REST API", func(t *testing.T) {

		config := Config{BaseURL: "https://api.github.com/"}
		meta, err := config.Meta()
		if err != nil {
			t.Fatalf("failed to return meta without error: %s", err.Error())
		}

		ctx := context.Background()
		client := meta.(*Owner).v3client
		_, _, err = client.APIMeta(ctx)
		if err != nil {
			t.Fatalf("failed to validate returned client without error: %s", err.Error())
		}

	})

	t.Run("returns an anonymous client for the v4 GraphQL API", func(t *testing.T) {

		// https://developer.github.com/v4/guides/forming-calls/#authenticating-with-graphql
		t.Skip("anonymous client for the v4 GraphQL API is unsupported")

	})

	t.Run("returns a v3 REST API client to manage individual resources", func(t *testing.T) {

		config := Config{
			Token:   testToken,
			BaseURL: "https://api.github.com/",
		}
		meta, err := config.Meta()
		if err != nil {
			t.Fatalf("failed to return meta without error: %s", err.Error())
		}

		ctx := context.Background()
		client := meta.(*Owner).v3client
		_, _, err = client.APIMeta(ctx)
		if err != nil {
			t.Fatalf("failed to validate returned client without error: %s", err.Error())
		}

	})

	t.Run("returns a v4 GraphQL API client to manage individual resources", func(t *testing.T) {

		config := Config{
			Token:   testToken,
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

		config := Config{
			Token:   testToken,
			BaseURL: "https://api.github.com/",
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
			BaseURL: "https://api.github.com/",
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
		variables := map[string]interface{}{
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
