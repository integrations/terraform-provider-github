package github

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestResourceGithubProjectRepositoryCreate(t *testing.T) {
	t.Parallel()
	client, requests := newProjectV2TestClient(t, func(query string) string {
		switch {
		case strings.Contains(query, "repository(owner:"):
			return `{"data":{"repository":{"id":"R_1","name":"planning","owner":{"login":"atls"}}}}`
		case strings.Contains(query, "linkProjectV2ToRepository"):
			return `{"data":{"linkProjectV2ToRepository":{"repository":{"id":"R_1"}}}}`
		case strings.Contains(query, "repositories(first:"):
			return `{"data":{"node":{"repositories":{"nodes":[{"id":"R_1"}],"pageInfo":{"hasNextPage":false,"endCursor":null}}}}}`
		case strings.Contains(query, "... on Repository"):
			return `{"data":{"node":{"id":"R_1","name":"planning","owner":{"login":"atls"}}}}`
		default:
			t.Fatalf("unexpected GraphQL operation: %s", query)
			return ""
		}
	})
	resource := resourceGithubProjectRepository()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{"project_id": "PVT_1", "repository_owner": "atls", "repository": "planning"})
	diags := resourceGithubProjectRepositoryCreate(t.Context(), d, &Owner{v4client: client})
	if diags.HasError() {
		t.Fatalf("creating project repository link returned diagnostics: %v\nrequests: %v", diags, *requests)
	}
	if d.Id() != "PVT_1:R_1" || len(*requests) != 4 {
		t.Fatalf("unexpected repository link result: id=%q operations=%d", d.Id(), len(*requests))
	}
}
