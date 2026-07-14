package github

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestResourceGithubProjectItemCreate(t *testing.T) {
	t.Parallel()
	client, requests := newProjectV2TestClient(t, func(query string) string {
		switch {
		case strings.Contains(query, "addProjectV2ItemById"):
			return `{"data":{"addProjectV2ItemById":{"item":{"id":"PVTI_1"}}}}`
		case strings.Contains(query, "node(id:"):
			return `{"data":{"node":{"id":"PVTI_1","isArchived":false,"project":{"id":"PVT_1"},"content":{"id":"I_1"}}}}`
		default:
			t.Fatalf("unexpected GraphQL operation: %s", query)
			return ""
		}
	})
	resource := resourceGithubProjectItem()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{"project_id": "PVT_1", "content_id": "I_1"})
	diags := resourceGithubProjectItemCreate(t.Context(), d, &Owner{v4client: client})
	if diags.HasError() {
		t.Fatalf("creating project item returned diagnostics: %v\nrequests: %v", diags, *requests)
	}
	if d.Id() != "PVTI_1" || len(*requests) != 2 {
		t.Fatalf("unexpected item result: id=%q operations=%d", d.Id(), len(*requests))
	}
}
