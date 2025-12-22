package github

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	gh "github.com/google/go-github/v67/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestDataSourceGithubEnterpriseSCIMGroupsRead_fetchAllPages(t *testing.T) {
	ts := githubApiMock([]*mockResponse{
		{
			ExpectedUri: "/scim/v2/enterprises/ent/Groups?count=2&startIndex=1",
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
    {"schemas": ["urn:ietf:params:scim:schemas:core:2.0:Group"], "id": "g1", "externalId": "eg1", "displayName": "Group One", "meta": {"resourceType": "Group"}},
    {"schemas": ["urn:ietf:params:scim:schemas:core:2.0:Group"], "id": "g2", "externalId": "eg2", "displayName": "Group Two"}
  ]
}`,
		},
		{
			ExpectedUri: "/scim/v2/enterprises/ent/Groups?count=2&startIndex=3",
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
    {"schemas": ["urn:ietf:params:scim:schemas:core:2.0:Group"], "id": "g3", "externalId": "eg3", "displayName": "Group Three"}
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

	r := dataSourceGithubEnterpriseSCIMGroups()
	d := schema.TestResourceDataRaw(t, r.Schema, map[string]any{
		"enterprise":          "ent",
		"results_per_page":    2,
		"filter":              "",
		"excluded_attributes": "",
	})

	diags := dataSourceGithubEnterpriseSCIMGroupsRead(context.Background(), d, owner)
	if len(diags) > 0 {
		t.Fatalf("unexpected diagnostics: %#v", diags)
	}

	if d.Id() == "" {
		t.Fatalf("expected ID to be set")
	}

	resources := d.Get("resources").([]any)
	if len(resources) != 3 {
		t.Fatalf("expected 3 groups, got %d", len(resources))
	}
	first := resources[0].(map[string]any)
	if first["id"].(string) != "g1" {
		t.Fatalf("expected first id to be g1, got %v", first["id"])
	}
}
