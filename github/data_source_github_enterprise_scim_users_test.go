package github

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	gh "github.com/google/go-github/v67/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestDataSourceGithubEnterpriseSCIMUsersRead_fetchAllPages_withFilter(t *testing.T) {
	filter := "userName eq \"test\""
	ts := githubApiMock([]*mockResponse{
		{
			ExpectedUri: "/scim/v2/enterprises/ent/Users?count=2&excludedAttributes=members&filter=userName+eq+%22test%22&startIndex=1",
			ExpectedHeaders: map[string]string{
				"Accept": enterpriseSCIMAcceptHeader,
			},
			StatusCode: 200,
			ResponseBody: `{
  "schemas": ["urn:ietf:params:scim:api:messages:2.0:ListResponse"],
  "totalResults": 3,
  "startIndex": 1,
  "itemsPerPage": 2,
  "Resources": [
    {"schemas": ["urn:ietf:params:scim:schemas:core:2.0:User"], "id": "u1", "externalId": "eu1", "userName": "test", "displayName": "Test User", "active": true},
    {"schemas": ["urn:ietf:params:scim:schemas:core:2.0:User"], "id": "u2", "externalId": "eu2", "userName": "test2", "displayName": "Test User 2", "active": false}
  ]
}`,
		},
		{
			ExpectedUri: "/scim/v2/enterprises/ent/Users?count=2&excludedAttributes=members&filter=userName+eq+%22test%22&startIndex=3",
			ExpectedHeaders: map[string]string{
				"Accept": enterpriseSCIMAcceptHeader,
			},
			StatusCode: 200,
			ResponseBody: `{
  "schemas": ["urn:ietf:params:scim:api:messages:2.0:ListResponse"],
  "totalResults": 3,
  "startIndex": 3,
  "itemsPerPage": 2,
  "Resources": [
    {"schemas": ["urn:ietf:params:scim:schemas:core:2.0:User"], "id": "u3", "externalId": "eu3", "userName": "test3", "displayName": "Test User 3", "active": true}
  ]
}`,
		},
	})
	defer ts.Close()

	httpClient := &http.Client{Transport: http.DefaultTransport}
	client := gh.NewClient(httpClient)
	baseURL, _ := url.Parse(ts.URL + "/")
	client.BaseURL = baseURL

	owner := &Owner{v3client: client}

	r := dataSourceGithubEnterpriseSCIMUsers()
	d := schema.TestResourceDataRaw(t, r.Schema, map[string]any{
		"enterprise":          "ent",
		"results_per_page":    2,
		"filter":              filter,
		"excluded_attributes": "members",
	})

	diags := dataSourceGithubEnterpriseSCIMUsersRead(context.Background(), d, owner)
	if len(diags) > 0 {
		t.Fatalf("unexpected diagnostics: %#v", diags)
	}

	resources := d.Get("resources").([]any)
	if len(resources) != 3 {
		t.Fatalf("expected 3 users, got %d", len(resources))
	}
	first := resources[0].(map[string]any)
	if first["user_name"].(string) != "test" {
		t.Fatalf("expected first user_name to be test, got %v", first["user_name"])
	}
}
