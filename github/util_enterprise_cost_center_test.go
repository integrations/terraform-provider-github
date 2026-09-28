package github

import (
	"context"
	"fmt"
	"maps"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestGetEnterpriseCostCenterPaginatesResources(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/api/v3/enterprises/example/settings/billing/cost-centers/123" {
			t.Fatalf("unexpected request path %q", req.URL.Path)
		}
		if got := req.URL.Query().Get("per_page"); got != "2" {
			t.Fatalf("expected per_page 2, got %q", got)
		}

		switch page := req.URL.Query().Get("page"); page {
		case "1":
			fmt.Fprint(w, `{
				"id": "123",
				"name": "Engineering",
				"state": "active",
				"resources": [
					{"type": "User", "name": "octocat"},
					{"type": "Org", "name": "octo-org"}
				]
			}`)
		case "2":
			fmt.Fprint(w, `{
				"id": "123",
				"name": "Engineering",
				"state": "active",
				"resources": [
					{"type": "Repo", "name": "octo-org/example"}
				]
			}`)
		default:
			t.Fatalf("unexpected page %q", page)
		}
	}))
	t.Cleanup(server.Close)

	client, err := github.NewClient(github.WithEnterpriseURLs(server.URL+"/", server.URL+"/"))
	if err != nil {
		t.Fatalf("configuring GitHub client: %v", err)
	}

	costCenter, err := getEnterpriseCostCenter(context.Background(), client, "example", "123", 2)
	if err != nil {
		t.Fatalf("getting cost center: %v", err)
	}
	if costCenter.GetID() != "123" {
		t.Fatalf("expected cost center ID 123, got %q", costCenter.GetID())
	}
	if len(costCenter.Resources) != 3 {
		t.Fatalf("expected 3 resources, got %d", len(costCenter.Resources))
	}
	if costCenter.Resources[2].GetName() != "octo-org/example" {
		t.Fatalf("expected second-page resource, got %q", costCenter.Resources[2].GetName())
	}
}

func TestListEnterpriseCostCentersPaginates(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/api/v3/enterprises/example/settings/billing/cost-centers" {
			t.Fatalf("unexpected request path %q", req.URL.Path)
		}
		if got := req.URL.Query().Get("per_page"); got != "2" {
			t.Fatalf("expected per_page 2, got %q", got)
		}
		if got := req.URL.Query().Get("state"); got != "active" {
			t.Fatalf("expected state active, got %q", got)
		}

		switch page := req.URL.Query().Get("page"); page {
		case "1":
			fmt.Fprint(w, `{"costCenters":[
				{"id":"1","name":"Engineering","state":"active"},
				{"id":"2","name":"Product","state":"active"}
			]}`)
		case "2":
			fmt.Fprint(w, `{"costCenters":[
				{"id":"3","name":"Support","state":"active"}
			]}`)
		default:
			t.Fatalf("unexpected page %q", page)
		}
	}))
	t.Cleanup(server.Close)

	client, err := github.NewClient(github.WithEnterpriseURLs(server.URL+"/", server.URL+"/"))
	if err != nil {
		t.Fatalf("configuring GitHub client: %v", err)
	}

	state := "active"
	costCenters, err := listEnterpriseCostCenters(context.Background(), client, "example", &github.ListCostCenterOptions{State: &state}, 2)
	if err != nil {
		t.Fatalf("listing cost centers: %v", err)
	}
	if len(costCenters) != 3 {
		t.Fatalf("expected 3 cost centers, got %d", len(costCenters))
	}
	if costCenters[2].GetName() != "Support" {
		t.Fatalf("expected second-page cost center, got %q", costCenters[2].GetName())
	}
}

func TestCostCenterAssignmentReadsClearArchivedParent(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, `{"id":"123","name":"Archived","state":"deleted","resources":[]}`)
	}))
	t.Cleanup(server.Close)

	client, err := github.NewClient(github.WithEnterpriseURLs(server.URL+"/", server.URL+"/"))
	if err != nil {
		t.Fatalf("configuring GitHub client: %v", err)
	}
	owner := &Owner{v3client: client, maxPerPage: 100}

	tests := []struct {
		name        string
		resource    *schema.Resource
		read        schema.ReadContextFunc
		assignments map[string]any
	}{
		{
			name:        "users",
			resource:    resourceGithubEnterpriseCostCenterUsers(),
			read:        resourceGithubEnterpriseCostCenterUsersRead,
			assignments: map[string]any{"usernames": []any{"octocat"}},
		},
		{
			name:        "organizations",
			resource:    resourceGithubEnterpriseCostCenterOrganizations(),
			read:        resourceGithubEnterpriseCostCenterOrganizationsRead,
			assignments: map[string]any{"organization_logins": []any{"octo-org"}},
		},
		{
			name:        "repositories",
			resource:    resourceGithubEnterpriseCostCenterRepositories(),
			read:        resourceGithubEnterpriseCostCenterRepositoriesRead,
			assignments: map[string]any{"repository_names": []any{"octo-org/example"}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			raw := map[string]any{
				"enterprise_slug": "example",
				"cost_center_id":  "123",
			}
			maps.Copy(raw, test.assignments)

			data := schema.TestResourceDataRaw(t, test.resource.Schema, raw)
			data.SetId("123")

			diags := test.read(context.Background(), data, owner)
			if diags.HasError() {
				t.Fatalf("reading assignment resource: %s", diagnosticsSummary(diags))
			}
			if data.Id() != "" {
				t.Fatalf("expected archived cost center assignments to be removed from state, got ID %q", data.Id())
			}
		})
	}
}

func diagnosticsSummary(diags diag.Diagnostics) string {
	if len(diags) == 0 {
		return ""
	}
	return diags[0].Summary
}
