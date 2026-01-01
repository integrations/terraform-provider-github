package github

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	gh "github.com/google/go-github/v67/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestDataSourceGithubEnterpriseSCIMGroupRead(t *testing.T) {
	ts := githubApiMock([]*mockResponse{
		{
			ExpectedUri: "/scim/v2/enterprises/ent/Groups/g1",
			ExpectedHeaders: map[string]string{
				"Accept": enterpriseSCIMAcceptHeader,
			},
			StatusCode: 200,
			ResponseBody: `{
  "schemas": ["urn:ietf:params:scim:schemas:core:2.0:Group"],
  "id": "g1",
  "externalId": "eg1",
  "displayName": "Group One",
  "members": [{"value": "u1", "$ref": "https://example.test/u1", "display": "user1"}],
  "meta": {"resourceType": "Group", "created": "2020-01-01T00:00:00Z"}
}`,
		},
	})
	defer ts.Close()

	httpClient := &http.Client{Transport: http.DefaultTransport}
	client := gh.NewClient(httpClient)
	baseURL, _ := url.Parse(ts.URL + "/")
	client.BaseURL = baseURL

	owner := &Owner{v3client: client}

	r := dataSourceGithubEnterpriseSCIMGroup()
	d := schema.TestResourceDataRaw(t, r.Schema, map[string]any{
		"enterprise":          "ent",
		"scim_group_id":       "g1",
		"excluded_attributes": "",
	})

	diags := dataSourceGithubEnterpriseSCIMGroupRead(context.Background(), d, owner)
	if len(diags) > 0 {
		t.Fatalf("unexpected diagnostics: %#v", diags)
	}

	if got := d.Get("id").(string); got != "g1" {
		t.Fatalf("expected id g1, got %q", got)
	}
	members := d.Get("members").([]any)
	if len(members) != 1 {
		t.Fatalf("expected 1 member, got %d", len(members))
	}
	m0 := members[0].(map[string]any)
	if m0["ref"].(string) != "https://example.test/u1" {
		t.Fatalf("expected ref to be set, got %v", m0["ref"])
	}
}
