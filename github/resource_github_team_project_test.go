package github

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestResourceGithubTeamProjectCreate(t *testing.T) {
	t.Parallel()
	client, requests := newProjectV2TestClient(t, func(query string) string {
		switch {
		case strings.Contains(query, "organization(login:"):
			return `{"data":{"organization":{"team":{"id":"T_1","slug":"platform","organization":{"login":"atls"}}}}}`
		case strings.Contains(query, "linkProjectV2ToTeam"):
			return `{"data":{"linkProjectV2ToTeam":{"team":{"id":"T_1"}}}}`
		case strings.Contains(query, "teams(first:"):
			return `{"data":{"node":{"teams":{"nodes":[{"id":"T_1"}],"pageInfo":{"hasNextPage":false,"endCursor":null}}}}}`
		case strings.Contains(query, "... on Team"):
			return `{"data":{"node":{"id":"T_1","slug":"platform","organization":{"login":"atls"}}}}`
		default:
			t.Fatalf("unexpected GraphQL operation: %s", query)
			return ""
		}
	})
	resource := resourceGithubTeamProject()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{"project_id": "PVT_1", "organization": "atls", "team_slug": "platform"})
	diags := resourceGithubTeamProjectCreate(t.Context(), d, &Owner{v4client: client})
	if diags.HasError() {
		t.Fatalf("creating team project link returned diagnostics: %v\nrequests: %v", diags, *requests)
	}
	if d.Id() != "PVT_1:T_1" || len(*requests) != 4 {
		t.Fatalf("unexpected team link result: id=%q operations=%d", d.Id(), len(*requests))
	}
}
