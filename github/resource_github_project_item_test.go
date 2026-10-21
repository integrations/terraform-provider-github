package github

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestResourceGithubProjectItemCreate(t *testing.T) {
	t.Parallel()
	client, requests := newProjectV2TestClient(t, func(request projectV2GraphQLRequest) string {
		query := request.Query
		switch {
		case strings.Contains(query, "addProjectV2ItemById"):
			return `{"data":{"addProjectV2ItemById":{"item":{"id":"PVTI_1","isArchived":false,"project":{"id":"PVT_1"},"content":{"id":"I_1"}}}}}`
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
	if d.Id() != "PVTI_1" || len(*requests) != 1 {
		t.Fatalf("unexpected item result: id=%q operations=%d", d.Id(), len(*requests))
	}
	assertProjectV2GraphQLInput(t, (*requests)[0], map[string]any{"projectId": "PVT_1", "contentId": "I_1"})
}

func TestResourceGithubProjectItemReadRemovesDeletedItem(t *testing.T) {
	t.Parallel()
	client, _ := newProjectV2TestClient(t, func(projectV2GraphQLRequest) string {
		return `{"data":{"node":null}}`
	})
	d := schema.TestResourceDataRaw(t, resourceGithubProjectItem().Schema, map[string]any{
		"project_id": "PVT_1", "content_id": "I_1",
	})
	d.SetId("PVTI_deleted")
	if diagnostics := resourceGithubProjectItemRead(t.Context(), d, &Owner{v4client: client}); diagnostics.HasError() {
		t.Fatalf("reading deleted item returned diagnostics: %v", diagnostics)
	}
	if d.Id() != "" {
		t.Fatalf("deleted item remained in state with ID %q", d.Id())
	}
}
