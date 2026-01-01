package github

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	gh "github.com/google/go-github/v67/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestDataSourceGithubEnterpriseSCIMUserRead(t *testing.T) {
	ts := githubApiMock([]*mockResponse{
		{
			ExpectedUri: "/scim/v2/enterprises/ent/Users/u1",
			ExpectedHeaders: map[string]string{
				"Accept": enterpriseSCIMAcceptHeader,
			},
			StatusCode: 200,
			ResponseBody: `{
  "schemas": ["urn:ietf:params:scim:schemas:core:2.0:User"],
  "id": "u1",
  "externalId": "eu1",
  "userName": "test",
  "displayName": "Test User",
  "active": true,
  "name": {"formatted": "Test User", "familyName": "User", "givenName": "Test"},
  "emails": [{"value": "test@example.com", "type": "work", "primary": true}],
  "roles": [{"value": "member", "display": "Member", "type": "direct", "primary": true}],
  "meta": {"resourceType": "User"}
}`,
		},
	})
	defer ts.Close()

	httpClient := &http.Client{Transport: http.DefaultTransport}
	client := gh.NewClient(httpClient)
	baseURL, _ := url.Parse(ts.URL + "/")
	client.BaseURL = baseURL

	owner := &Owner{v3client: client}

	r := dataSourceGithubEnterpriseSCIMUser()
	d := schema.TestResourceDataRaw(t, r.Schema, map[string]any{
		"enterprise":          "ent",
		"scim_user_id":        "u1",
		"excluded_attributes": "",
	})

	diags := dataSourceGithubEnterpriseSCIMUserRead(context.Background(), d, owner)
	if len(diags) > 0 {
		t.Fatalf("unexpected diagnostics: %#v", diags)
	}

	if got := d.Get("user_name").(string); got != "test" {
		t.Fatalf("expected user_name test, got %q", got)
	}
	name := d.Get("name").([]any)
	if len(name) != 1 {
		t.Fatalf("expected name block, got %d", len(name))
	}
	n0 := name[0].(map[string]any)
	if n0["given_name"].(string) != "Test" {
		t.Fatalf("expected given_name Test, got %v", n0["given_name"])
	}

	emails := d.Get("emails").([]any)
	if len(emails) != 1 {
		t.Fatalf("expected 1 email, got %d", len(emails))
	}
}
