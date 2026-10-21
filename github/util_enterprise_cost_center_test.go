package github

import (
	"context"
	"fmt"
	"maps"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
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
			w.Header().Set("Link", `<https://example.invalid/api/v3/enterprises/example/settings/billing/cost-centers/123?page=2&per_page=2>; rel="next"`)
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
			w.Header().Set("Link", `<https://example.invalid/api/v3/enterprises/example/settings/billing/cost-centers?page=2&per_page=2&state=active>; rel="next"`)
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

func TestCostCenterAssignmentSchemasNormalizeCase(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		resource   *schema.Resource
		field      string
		configured string
		canonical  string
	}{
		{
			name:       "users",
			resource:   resourceGithubEnterpriseCostCenterUsers(),
			field:      "usernames",
			configured: "MONALISA",
			canonical:  "Monalisa",
		},
		{
			name:       "organizations",
			resource:   resourceGithubEnterpriseCostCenterOrganizations(),
			field:      "organization_logins",
			configured: "OCTO-ORG",
			canonical:  "Octo-Org",
		},
		{
			name:       "repositories",
			resource:   resourceGithubEnterpriseCostCenterRepositories(),
			field:      "repository_names",
			configured: "OCTO-ORG/EXAMPLE",
			canonical:  "Octo-Org/Example",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			fieldSchema := test.resource.Schema[test.field]
			if fieldSchema.Set(test.configured) != fieldSchema.Set(test.canonical) {
				t.Fatalf("expected %s set hashing to be case-insensitive", test.field)
			}
			elementSchema, ok := fieldSchema.Elem.(*schema.Schema)
			if !ok {
				t.Fatalf("expected %s elements to have a schema", test.field)
			}
			if got := elementSchema.StateFunc(test.configured); got != strings.ToLower(test.configured) {
				t.Fatalf("expected normalized %s state %q, got %q", test.field, strings.ToLower(test.configured), got)
			}

			data := schema.TestResourceDataRaw(t, test.resource.Schema, map[string]any{
				"enterprise_slug": "example",
				"cost_center_id":  "123",
				test.field:        []any{test.configured, test.canonical},
			})
			values, ok := data.Get(test.field).(*schema.Set)
			if !ok {
				t.Fatalf("expected %s to be a set", test.field)
			}
			if values.Len() != 1 {
				t.Fatalf("expected case variants in %s to collapse to one value, got %d", test.field, values.Len())
			}
		})
	}
}

func TestCostCenterAssignmentUpdatesIgnoreCaseOnlyDifferences(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		resource     *schema.Resource
		update       schema.UpdateContextFunc
		read         schema.ReadContextFunc
		field        string
		resourceType string
		configured   string
		canonical    string
	}{
		{
			name:         "users",
			resource:     resourceGithubEnterpriseCostCenterUsers(),
			update:       resourceGithubEnterpriseCostCenterUsersUpdate,
			read:         resourceGithubEnterpriseCostCenterUsersRead,
			field:        "usernames",
			resourceType: CostCenterResourceTypeUser,
			configured:   "MONALISA",
			canonical:    "Monalisa",
		},
		{
			name:         "organizations",
			resource:     resourceGithubEnterpriseCostCenterOrganizations(),
			update:       resourceGithubEnterpriseCostCenterOrganizationsUpdate,
			read:         resourceGithubEnterpriseCostCenterOrganizationsRead,
			field:        "organization_logins",
			resourceType: CostCenterResourceTypeOrg,
			configured:   "OCTO-ORG",
			canonical:    "Octo-Org",
		},
		{
			name:         "repositories",
			resource:     resourceGithubEnterpriseCostCenterRepositories(),
			update:       resourceGithubEnterpriseCostCenterRepositoriesUpdate,
			read:         resourceGithubEnterpriseCostCenterRepositoriesRead,
			field:        "repository_names",
			resourceType: CostCenterResourceTypeRepo,
			configured:   "OCTO-ORG/EXAMPLE",
			canonical:    "Octo-Org/Example",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var mutations atomic.Int32
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				if req.Method != http.MethodGet {
					mutations.Add(1)
					w.WriteHeader(http.StatusNoContent)
					return
				}

				fmt.Fprintf(w, `{
					"id": "123",
					"name": "Engineering",
					"state": "active",
					"resources": [{"type": %q, "name": %q}]
				}`, test.resourceType, test.canonical)
			}))
			t.Cleanup(server.Close)

			client, err := github.NewClient(github.WithEnterpriseURLs(server.URL+"/", server.URL+"/"))
			if err != nil {
				t.Fatalf("configuring GitHub client: %v", err)
			}

			data := schema.TestResourceDataRaw(t, test.resource.Schema, map[string]any{
				"enterprise_slug": "example",
				"cost_center_id":  "123",
				test.field:        []any{test.configured},
			})
			data.SetId("123")

			diags := test.update(t.Context(), data, &Owner{v3client: client, maxPerPage: 100})
			if diags.HasError() {
				t.Fatalf("updating assignment resource: %s", diagnosticsSummary(diags))
			}
			if got := mutations.Load(); got != 0 {
				t.Fatalf("expected no mutation for a case-only difference, got %d", got)
			}

			diags = test.read(t.Context(), data, &Owner{v3client: client, maxPerPage: 100})
			if diags.HasError() {
				t.Fatalf("reading assignment resource: %s", diagnosticsSummary(diags))
			}
			assertNormalizedAssignmentSet(t, data, test.field, strings.ToLower(test.canonical))
		})
	}
}

func assertNormalizedAssignmentSet(t *testing.T, data *schema.ResourceData, field, expected string) {
	t.Helper()

	values, ok := data.Get(field).(*schema.Set)
	if !ok {
		t.Fatalf("expected %s to be a set", field)
	}
	if values.Len() != 1 {
		t.Fatalf("expected case variants in %s to collapse to one value, got %d", field, values.Len())
	}

	value, ok := values.List()[0].(string)
	if !ok {
		t.Fatalf("expected %s to contain strings", field)
	}
	if value != expected {
		t.Fatalf("expected normalized %s value %q, got %q", field, expected, value)
	}
}

func diagnosticsSummary(diags diag.Diagnostics) string {
	if len(diags) == 0 {
		return ""
	}
	return diags[0].Summary
}
