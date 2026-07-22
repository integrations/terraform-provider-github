package github

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestResourceGithubProjectRepositoryCreate(t *testing.T) {
	t.Parallel()
	client, requests := newProjectV2TestClient(t, func(request projectV2GraphQLRequest) string {
		query := request.Query
		switch {
		case strings.Contains(query, "repository(owner:"):
			return `{"data":{"repository":{"id":"R_1","databaseId":201,"name":"planning","owner":{"login":"atls"}}}}`
		case strings.Contains(query, "linkProjectV2ToRepository"):
			return `{"data":{"linkProjectV2ToRepository":{"repository":{"id":"R_1","databaseId":201,"name":"planning","owner":{"login":"atls"}}}}}`
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
	if d.Id() != "PVT_1:R_1" || len(*requests) != 2 {
		t.Fatalf("unexpected repository link result: id=%q operations=%d", d.Id(), len(*requests))
	}
	if projectV2Get[int](d, "repository_id") != 201 {
		t.Fatalf("repository database ID was not stored: %v", d.Get("repository_id"))
	}
	assertProjectV2GraphQLInput(t, (*requests)[1], map[string]any{"projectId": "PVT_1", "repositoryId": "R_1"})
}

func TestResourceGithubProjectRepositoryReadUsesConnectionNode(t *testing.T) {
	t.Parallel()
	client, requests := newProjectV2TestClient(t, func(request projectV2GraphQLRequest) string {
		if !strings.Contains(request.Query, "repositories(first:") {
			t.Fatalf("unexpected GraphQL operation: %s", request.Query)
		}
		return `{"data":{"node":{"repositories":{"nodes":[{"id":"R_1","databaseId":201,"name":"planning-renamed","owner":{"login":"atls-renamed"}}],"pageInfo":{"hasNextPage":false,"endCursor":null}}}}}`
	})
	resource := resourceGithubProjectRepository()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{"project_id": "PVT_1", "repository_owner": "atls", "repository": "planning"})
	d.SetId("PVT_1:R_1")
	if diagnostics := resourceGithubProjectRepositoryRead(t.Context(), d, &Owner{v4client: client}); diagnostics.HasError() {
		t.Fatalf("reading project repository link returned diagnostics: %v", diagnostics)
	}
	if len(*requests) != 1 || projectV2Get[string](d, "repository") != "planning-renamed" || projectV2Get[int](d, "repository_id") != 201 {
		t.Fatalf("unexpected repository link read: operations=%d state=%#v", len(*requests), d.State())
	}
}

func TestResourceGithubProjectRepositoryUpdateRefreshesRenamedIdentity(t *testing.T) {
	t.Parallel()
	client, requests := newProjectV2TestClient(t, func(request projectV2GraphQLRequest) string {
		if !strings.Contains(request.Query, "repository(owner:") {
			t.Fatalf("unexpected GraphQL operation: %s", request.Query)
		}
		return `{"data":{"repository":{"id":"R_1","databaseId":201,"name":"planning-renamed","owner":{"login":"atls-renamed"}}}}`
	})
	resource := resourceGithubProjectRepository()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{"project_id": "PVT_1", "repository_owner": "atls-renamed", "repository": "planning-renamed"})
	d.SetId("PVT_1:R_1")
	if diagnostics := resourceGithubProjectRepositoryUpdate(t.Context(), d, &Owner{v4client: client}); diagnostics.HasError() {
		t.Fatalf("updating renamed project repository link returned diagnostics: %v", diagnostics)
	}
	if len(*requests) != 1 || projectV2Get[string](d, "repository") != "planning-renamed" || projectV2Get[int](d, "repository_id") != 201 {
		t.Fatalf("unexpected repository rename state: operations=%d state=%#v", len(*requests), d.State())
	}
}
