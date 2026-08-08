package github

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestResourceGithubTeamProjectCreate(t *testing.T) {
	t.Parallel()
	client, requests := newProjectV2TestClient(t, func(request projectV2GraphQLRequest) string {
		query := request.Query
		switch {
		case strings.Contains(query, "organization(login:"):
			return `{"data":{"organization":{"team":{"id":"T_1","databaseId":301,"slug":"platform","organization":{"login":"atls"}}}}}`
		case strings.Contains(query, "linkProjectV2ToTeam"):
			return `{"data":{"linkProjectV2ToTeam":{"team":{"id":"T_1","databaseId":301,"slug":"platform","organization":{"login":"atls"}}}}}`
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
	if d.Id() != "PVT_1:T_1" || len(*requests) != 2 {
		t.Fatalf("unexpected team link result: id=%q operations=%d", d.Id(), len(*requests))
	}
	if projectV2Get[int](d, "team_id") != 301 {
		t.Fatalf("team database ID was not stored: %v", d.Get("team_id"))
	}
	assertProjectV2GraphQLInput(t, (*requests)[1], map[string]any{"projectId": "PVT_1", "teamId": "T_1"})
}

func TestResourceGithubTeamProjectReadUsesConnectionNode(t *testing.T) {
	t.Parallel()
	client, requests := newProjectV2TestClient(t, func(request projectV2GraphQLRequest) string {
		if !strings.Contains(request.Query, "teams(first:") {
			t.Fatalf("unexpected GraphQL operation: %s", request.Query)
		}
		return `{"data":{"node":{"teams":{"nodes":[{"id":"T_1","databaseId":301,"slug":"platform-renamed","organization":{"login":"atls-renamed"}}],"pageInfo":{"hasNextPage":false,"endCursor":null}}}}}`
	})
	resource := resourceGithubTeamProject()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{"project_id": "PVT_1", "organization": "atls", "team_slug": "platform"})
	d.SetId("PVT_1:T_1")
	if diagnostics := resourceGithubTeamProjectRead(t.Context(), d, &Owner{v4client: client}); diagnostics.HasError() {
		t.Fatalf("reading team project link returned diagnostics: %v", diagnostics)
	}
	if len(*requests) != 1 || projectV2Get[string](d, "team_slug") != "platform-renamed" || projectV2Get[int](d, "team_id") != 301 {
		t.Fatalf("unexpected team project read: operations=%d state=%#v", len(*requests), d.State())
	}
}

func TestResourceGithubTeamProjectUpdateRefreshesRenamedIdentity(t *testing.T) {
	t.Parallel()
	client, requests := newProjectV2TestClient(t, func(request projectV2GraphQLRequest) string {
		if !strings.Contains(request.Query, "organization(login:") {
			t.Fatalf("unexpected GraphQL operation: %s", request.Query)
		}
		return `{"data":{"organization":{"team":{"id":"T_1","databaseId":301,"slug":"platform-renamed","organization":{"login":"atls-renamed"}}}}}`
	})
	resource := resourceGithubTeamProject()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{"project_id": "PVT_1", "organization": "atls-renamed", "team_slug": "platform-renamed"})
	d.SetId("PVT_1:T_1")
	if diagnostics := resourceGithubTeamProjectUpdate(t.Context(), d, &Owner{v4client: client}); diagnostics.HasError() {
		t.Fatalf("updating renamed team project link returned diagnostics: %v", diagnostics)
	}
	if len(*requests) != 1 || projectV2Get[string](d, "team_slug") != "platform-renamed" || projectV2Get[int](d, "team_id") != 301 {
		t.Fatalf("unexpected team rename state: operations=%d state=%#v", len(*requests), d.State())
	}
}
