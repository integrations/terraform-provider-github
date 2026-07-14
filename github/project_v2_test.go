package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"
)

func TestProjectV2ResourcesValidate(t *testing.T) {
	t.Parallel()

	resources := map[string]*schema.Resource{
		"github_project":                  resourceGithubProject(),
		"github_project_field":            resourceGithubProjectField(),
		"github_project_item":             resourceGithubProjectItem(),
		"github_project_item_field_value": resourceGithubProjectItemFieldValue(),
		"github_project_repository":       resourceGithubProjectRepository(),
		"github_team_project":             resourceGithubTeamProject(),
	}
	for name, resource := range resources {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if len(resource.Schema) == 0 {
				t.Fatal("resource has no schema")
			}
			if NewProvider("test", "none")().ResourcesMap[name] == nil {
				t.Fatalf("resource is not registered in the provider")
			}
		})
	}
}

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
		"owner_type":        projectV2OwnerOrganization,
		"owner":             "atls",
		"title":             "Planning",
		"short_description": "Operations",
		"readme":            "# Planning",
		"public":            false,
		"closed":            false,
	})
	diags := resourceGithubProjectCreate(t.Context(), d, &Owner{name: "atls", v4client: client})
	if diags.HasError() {
		t.Fatalf("creating project returned diagnostics: %v", diags)
	}
	if d.Id() != "PVT_1" {
		t.Fatalf("unexpected project ID %q", d.Id())
	}
	if got := d.Get("number").(int); got != 11 {
		t.Fatalf("unexpected project number %d", got)
	}
	if len(*requests) != 4 {
		t.Fatalf("expected 4 GraphQL operations, got %d", len(*requests))
	}
}

func TestResourceGithubProjectFieldCreate(t *testing.T) {
	t.Parallel()

	client, requests := newProjectV2TestClient(t, func(query string) string {
		switch {
		case strings.Contains(query, "createProjectV2Field"):
			return `{"data":{"createProjectV2Field":{"projectV2Field":{"__typename":"ProjectV2SingleSelectField","id":"PVTF_1","name":"Status","dataType":"SINGLE_SELECT","project":{"id":"PVT_1"},"options":[{"id":"opt-1","name":"To Do","description":"Ready","color":"GRAY"}]}}}}`
		case strings.Contains(query, "node(id:"):
			return `{"data":{"node":{"__typename":"ProjectV2SingleSelectField","id":"PVTF_1","name":"Status","dataType":"SINGLE_SELECT","project":{"id":"PVT_1"},"options":[{"id":"opt-1","name":"To Do","description":"Ready","color":"GRAY"}]}}}`
		default:
			t.Fatalf("unexpected GraphQL operation: %s", query)
			return ""
		}
	})

	resource := resourceGithubProjectField()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"project_id": "PVT_1",
		"name":       "Status",
		"data_type":  "SINGLE_SELECT",
		"single_select_option": []any{map[string]any{
			"name": "To Do", "description": "Ready", "color": "GRAY",
		}},
	})
	diags := resourceGithubProjectFieldCreate(t.Context(), d, &Owner{v4client: client})
	if diags.HasError() {
		t.Fatalf("creating project field returned diagnostics: %v\nrequests: %v", diags, *requests)
	}
	if d.Id() != "PVTF_1" {
		t.Fatalf("unexpected field ID %q", d.Id())
	}
	if len(*requests) != 2 {
		t.Fatalf("expected 2 GraphQL operations, got %d", len(*requests))
	}
	if !strings.Contains((*requests)[0], "... on ProjectV2SingleSelectField") {
		t.Fatalf("create mutation does not request single-select field data: %s", (*requests)[0])
	}
}

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
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"project_id": "PVT_1",
		"content_id": "I_1",
	})
	diags := resourceGithubProjectItemCreate(t.Context(), d, &Owner{v4client: client})
	if diags.HasError() {
		t.Fatalf("creating project item returned diagnostics: %v\nrequests: %v", diags, *requests)
	}
	if d.Id() != "PVTI_1" {
		t.Fatalf("unexpected item ID %q", d.Id())
	}
	if len(*requests) != 2 {
		t.Fatalf("expected 2 GraphQL operations, got %d", len(*requests))
	}
}

func TestResourceGithubProjectItemFieldValueCreate(t *testing.T) {
	t.Parallel()

	client, requests := newProjectV2TestClient(t, func(query string) string {
		switch {
		case strings.Contains(query, "updateProjectV2ItemFieldValue"):
			return `{"data":{"updateProjectV2ItemFieldValue":{"projectV2Item":{"id":"PVTI_1"}}}}`
		case strings.Contains(query, "fieldValueByName"):
			return `{"data":{"node":{"__typename":"ProjectV2Item","fieldValueByName":{"__typename":"ProjectV2ItemFieldNumberValue","id":"PVTFV_1","number":0}}}}`
		case strings.Contains(query, "node(id:"):
			return `{"data":{"node":{"__typename":"ProjectV2Field","id":"PVTF_1","name":"Estimate","dataType":"NUMBER","project":{"id":"PVT_1"}}}}`
		default:
			t.Fatalf("unexpected GraphQL operation: %s", query)
			return ""
		}
	})

	resource := resourceGithubProjectItemFieldValue()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"project_id": "PVT_1",
		"item_id":    "PVTI_1",
		"field_id":   "PVTF_1",
		"number":     0.0,
	})
	diags := resourceGithubProjectItemFieldValueCreateOrUpdate(t.Context(), d, &Owner{v4client: client})
	if diags.HasError() {
		t.Fatalf("creating project item field value returned diagnostics: %v\nrequests: %v", diags, *requests)
	}
	if d.Id() != "PVT_1:PVTI_1:PVTF_1" {
		t.Fatalf("unexpected field value ID %q", d.Id())
	}
	if len(*requests) != 3 {
		t.Fatalf("expected 3 GraphQL operations, got %d", len(*requests))
	}
}

func TestResourceGithubProjectRepositoryCreate(t *testing.T) {
	t.Parallel()

	client, requests := newProjectV2TestClient(t, func(query string) string {
		switch {
		case strings.Contains(query, "repository(owner:"):
			return `{"data":{"repository":{"id":"R_1","name":"planning","nameWithOwner":"atls/planning","owner":{"login":"atls"}}}}`
		case strings.Contains(query, "linkProjectV2ToRepository"):
			return `{"data":{"linkProjectV2ToRepository":{"repository":{"id":"R_1"}}}}`
		case strings.Contains(query, "repositories(first:"):
			return `{"data":{"node":{"repositories":{"nodes":[{"id":"R_1"}],"pageInfo":{"hasNextPage":false,"endCursor":null}}}}}`
		case strings.Contains(query, "... on Repository"):
			return `{"data":{"node":{"id":"R_1","name":"planning","nameWithOwner":"atls/planning","owner":{"login":"atls"}}}}`
		default:
			t.Fatalf("unexpected GraphQL operation: %s", query)
			return ""
		}
	})

	resource := resourceGithubProjectRepository()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"project_id":       "PVT_1",
		"repository_owner": "atls",
		"repository":       "planning",
	})
	diags := resourceGithubProjectRepositoryCreate(t.Context(), d, &Owner{v4client: client})
	if diags.HasError() {
		t.Fatalf("creating project repository link returned diagnostics: %v\nrequests: %v", diags, *requests)
	}
	if d.Id() != "PVT_1:R_1" {
		t.Fatalf("unexpected repository link ID %q", d.Id())
	}
	if len(*requests) != 4 {
		t.Fatalf("expected 4 GraphQL operations, got %d", len(*requests))
	}
}

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
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"project_id":   "PVT_1",
		"organization": "atls",
		"team_slug":    "platform",
	})
	diags := resourceGithubTeamProjectCreate(t.Context(), d, &Owner{v4client: client})
	if diags.HasError() {
		t.Fatalf("creating team project link returned diagnostics: %v\nrequests: %v", diags, *requests)
	}
	if d.Id() != "PVT_1:T_1" {
		t.Fatalf("unexpected team link ID %q", d.Id())
	}
	if len(*requests) != 4 {
		t.Fatalf("expected 4 GraphQL operations, got %d", len(*requests))
	}
}

func TestExpandProjectV2ItemFieldValueSupportsZero(t *testing.T) {
	t.Parallel()

	resource := resourceGithubProjectItemFieldValue()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"project_id": "PVT_1",
		"item_id":    "PVTI_1",
		"field_id":   "PVTF_1",
		"number":     0.0,
	})
	value, err := expandProjectV2ItemFieldValue(d)
	if err != nil {
		t.Fatalf("expanding field value: %v", err)
	}
	if value.Number == nil || float64(*value.Number) != 0 {
		t.Fatalf("zero number was not preserved: %#v", value.Number)
	}
}

func TestSetProjectV2ItemFieldValueStateUsesTypename(t *testing.T) {
	t.Parallel()

	resource := resourceGithubProjectItemFieldValue()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"project_id": "PVT_1",
		"item_id":    "PVTI_1",
		"field_id":   "PVTF_1",
		"text":       "stale",
	})
	value := projectV2ItemFieldValueNode{Typename: "ProjectV2ItemFieldNumberValue"}
	value.Text.ID = "PVTFV_1"
	value.Number.ID = "PVTFV_1"
	value.Number.Number = 0

	if err := setProjectV2ItemFieldValueState(d, value); err != nil {
		t.Fatalf("setting field value state: %v", err)
	}
	if got := d.Get("number").(float64); got != 0 {
		t.Fatalf("unexpected number %v", got)
	}
	if got := d.Get("text").(string); got != "" {
		t.Fatalf("stale text value was preserved: %q", got)
	}
}

func TestValidateProjectV2Date(t *testing.T) {
	t.Parallel()

	if _, errors := validateProjectV2Date("2026-07-14", "date"); len(errors) != 0 {
		t.Fatalf("valid date was rejected: %v", errors)
	}
	if _, errors := validateProjectV2Date("2026-07-14T00:00:00Z", "date"); len(errors) == 0 {
		t.Fatal("RFC3339 timestamp was accepted as a Projects V2 date")
	}
}

func TestSetProjectV2FieldStatePreservesIterationStartDate(t *testing.T) {
	t.Parallel()

	resource := resourceGithubProjectField()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"project_id": "PVT_1",
		"name":       "Sprint",
		"data_type":  "ITERATION",
		"iteration_configuration": []any{map[string]any{
			"start_date": "2026-01-05",
			"duration":   14,
		}},
	})
	field := projectV2FieldNode{Typename: "ProjectV2IterationField"}
	field.Iteration.ID = "PVTF_1"
	field.Iteration.Name = "Sprint"
	field.Iteration.DataType = "ITERATION"
	field.Iteration.Project.ID = "PVT_1"
	field.Iteration.Configuration.Duration = 14
	field.Iteration.Configuration.Iterations = []projectV2IterationFragment{{
		ID: "iteration-1", Title: "Sprint 14", StartDate: githubv4.Date{Time: mustParseProjectV2Date(t, "2026-07-06")}, Duration: 14,
	}}

	if err := setProjectV2FieldState(d, field); err != nil {
		t.Fatalf("setting iteration field state: %v", err)
	}
	configuration := d.Get("iteration_configuration").([]any)[0].(map[string]any)
	if got := configuration["start_date"].(string); got != "2026-01-05" {
		t.Fatalf("iteration start date changed to %q", got)
	}
}

func mustParseProjectV2Date(t *testing.T, value string) time.Time {
	t.Helper()
	parsed, err := time.Parse(time.DateOnly, value)
	if err != nil {
		t.Fatalf("parsing date: %v", err)
	}
	return parsed
}

func newProjectV2TestClient(t *testing.T, response func(query string) string) (*githubv4.Client, *[]string) {
	t.Helper()
	requests := make([]string, 0)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var payload struct {
			Query string `json:"query"`
		}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			t.Errorf("decoding GraphQL request: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		requests = append(requests, payload.Query)
		body := response(payload.Query)
		w.Header().Set("Content-Type", "application/json")
		if _, err := fmt.Fprint(w, body); err != nil {
			t.Errorf("writing GraphQL response: %v", err)
		}
	}))
	t.Cleanup(server.Close)
	return githubv4.NewEnterpriseClient(server.URL, server.Client()), &requests
}
