package github

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestResourceGithubProjectLifecycle(t *testing.T) {
	t.Parallel()
	client, requests := newProjectV2TestClient(t, func(query string) string {
		switch {
		case strings.Contains(query, "organization(login:"):
			return `{"data":{"organization":{"id":"O_1"}}}`
		case strings.Contains(query, "createProjectV2"):
			return `{"data":{"createProjectV2":{"projectV2":{"id":"PVT_1"}}}}`
		case strings.Contains(query, "updateProjectV2"):
			return `{"data":{"updateProjectV2":{"projectV2":{"id":"PVT_1"}}}}`
		case strings.Contains(query, "node(id:"):
			return `{"data":{"node":{"id":"PVT_1","number":11,"title":"Planning","shortDescription":"Operations","readme":"# Planning","public":false,"closed":false,"url":"https://github.com/orgs/atls/projects/11","owner":{"__typename":"Organization","login":"atls"}}}}`
		default:
			t.Fatalf("unexpected GraphQL operation: %s", query)
			return ""
		}
	})
	resource := resourceGithubProject()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"owner_type": projectV2OwnerOrganization, "owner": "atls", "title": "Planning", "short_description": "Operations", "readme": "# Planning", "public": false, "closed": false,
	})
	diags := resourceGithubProjectCreate(t.Context(), d, &Owner{name: "atls", v4client: client})
	if diags.HasError() {
		t.Fatalf("creating project returned diagnostics: %v", diags)
	}
	if d.Id() != "PVT_1" || projectV2Get[int](d, "number") != 11 {
		t.Fatalf("unexpected project state: id=%q number=%v", d.Id(), d.Get("number"))
	}
	if len(*requests) != 4 {
		t.Fatalf("expected 4 GraphQL operations, got %d", len(*requests))
	}
}
