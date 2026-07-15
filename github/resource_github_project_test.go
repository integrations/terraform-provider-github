package github

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestResourceGithubProjectLifecycle(t *testing.T) {
	t.Parallel()
	client, requests := newProjectV2TestClient(t, func(request projectV2GraphQLRequest) string {
		query := request.Query
		switch {
		case strings.Contains(query, "organization(login:"):
			return `{"data":{"organization":{"id":"O_1"}}}`
		case strings.Contains(query, "createProjectV2"):
			return `{"data":{"createProjectV2":{"projectV2":{"id":"PVT_1","number":11,"title":"Planning","shortDescription":"","readme":"","public":false,"closed":false,"url":"https://github.com/orgs/atls/projects/11","owner":{"__typename":"Organization","databaseId":101,"login":"atls"}}}}}`
		case strings.Contains(query, "updateProjectV2"):
			return `{"data":{"updateProjectV2":{"projectV2":{"id":"PVT_1","number":11,"title":"Planning","shortDescription":"Operations","readme":"# Planning","public":false,"closed":false,"url":"https://github.com/orgs/atls/projects/11","owner":{"__typename":"Organization","databaseId":101,"login":"atls"}}}}}`
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
	if d.Id() != "PVT_1" || projectV2Get[int](d, "number") != 11 || projectV2Get[int](d, "owner_id") != 101 {
		t.Fatalf("unexpected project state: id=%q number=%v", d.Id(), d.Get("number"))
	}
	if len(*requests) != 3 {
		t.Fatalf("expected 3 GraphQL operations, got %d", len(*requests))
	}
	assertProjectV2GraphQLInput(t, (*requests)[1], map[string]any{"ownerId": "O_1", "title": "Planning"})
	assertProjectV2GraphQLInput(t, (*requests)[2], map[string]any{"projectId": "PVT_1", "shortDescription": "Operations", "readme": "# Planning"})
}

func TestResourceGithubProjectDefaultsOwnerKindFromProvider(t *testing.T) {
	t.Parallel()
	client, requests := newProjectV2TestClient(t, func(request projectV2GraphQLRequest) string {
		switch {
		case strings.Contains(request.Query, "user(login:"):
			return `{"data":{"user":{"id":"U_1"}}}`
		case strings.Contains(request.Query, "createProjectV2"):
			return `{"data":{"createProjectV2":{"projectV2":{"id":"PVT_1","number":1,"title":"Personal","shortDescription":"","readme":"","public":false,"closed":false,"url":"https://github.com/users/alice/projects/1","owner":{"__typename":"User","databaseId":101,"login":"alice"}}}}}`
		case strings.Contains(request.Query, "updateProjectV2"):
			return `{"data":{"updateProjectV2":{"projectV2":{"id":"PVT_1","number":1,"title":"Personal","shortDescription":"","readme":"","public":false,"closed":false,"url":"https://github.com/users/alice/projects/1","owner":{"__typename":"User","databaseId":101,"login":"alice"}}}}}`
		default:
			t.Fatalf("unexpected GraphQL operation: %s", request.Query)
			return ""
		}
	})
	resource := resourceGithubProject()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{"title": "Personal"})
	if diagnostics := resourceGithubProjectCreate(t.Context(), d, &Owner{name: "alice", IsOrganization: false, v4client: client}); diagnostics.HasError() {
		t.Fatalf("creating user project returned diagnostics: %v", diagnostics)
	}
	if len(*requests) != 3 || projectV2Get[string](d, "owner_type") != projectV2OwnerUser || projectV2Get[string](d, "owner") != "alice" {
		t.Fatalf("provider user owner was not preserved: operations=%d owner_type=%v owner=%v", len(*requests), d.Get("owner_type"), d.Get("owner"))
	}
}
