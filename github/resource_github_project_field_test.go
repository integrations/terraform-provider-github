package github

import (
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	fieldapplication "github.com/integrations/terraform-provider-github/v6/internal/application/projects/field"
)

func TestResourceGithubProjectFieldCreate(t *testing.T) {
	t.Parallel()
	client, requests := newProjectV2TestClient(t, func(request projectV2GraphQLRequest) string {
		query := request.Query
		switch {
		case strings.Contains(query, "createProjectV2Field"):
			return `{"data":{"createProjectV2Field":{"projectV2Field":{"__typename":"ProjectV2SingleSelectField","id":"PVTF_1","name":"Status","dataType":"SINGLE_SELECT","project":{"id":"PVT_1"},"options":[{"id":"opt-1","name":"To Do","description":"Ready","color":"GRAY"}]}}}}`
		default:
			t.Fatalf("unexpected GraphQL operation: %s", query)
			return ""
		}
	})
	resource := resourceGithubProjectField()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"project_id": "PVT_1", "name": "Status", "data_type": "SINGLE_SELECT",
		"single_select_option": []any{map[string]any{"name": "To Do", "description": "Ready", "color": "GRAY"}},
	})
	diags := resourceGithubProjectFieldCreate(t.Context(), d, &Owner{v4client: client})
	if diags.HasError() {
		t.Fatalf("creating project field returned diagnostics: %v\nrequests: %v", diags, *requests)
	}
	if d.Id() != "PVTF_1" || len(*requests) != 1 {
		t.Fatalf("unexpected field result: id=%q operations=%d", d.Id(), len(*requests))
	}
	if !strings.Contains((*requests)[0].Query, "... on ProjectV2SingleSelectField") {
		t.Fatalf("create mutation does not request single-select field data: %s", (*requests)[0].Query)
	}
	assertProjectV2GraphQLInput(t, (*requests)[0], map[string]any{"projectId": "PVT_1", "name": "Status", "dataType": "SINGLE_SELECT"})
}

func TestSetProjectV2FieldStatePreservesIterationStartDate(t *testing.T) {
	t.Parallel()
	resource := resourceGithubProjectField()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"project_id": "PVT_1", "name": "Sprint", "data_type": "ITERATION",
		"iteration_configuration": []any{map[string]any{"start_date": "2026-01-05", "duration": 14}},
	})
	field := fieldapplication.Result{
		ID: "PVTF_1", ProjectID: "PVT_1", Name: "Sprint", DataType: "ITERATION",
		IterationConfiguration: &fieldapplication.IterationConfiguration{Duration: 14, Iterations: []fieldapplication.Iteration{{
			ID: "iteration-1", Title: "Sprint 14", StartDate: mustParseProjectV2Date(t, "2026-07-06"), Duration: 14,
		}}},
	}
	if err := setProjectV2FieldState(d, field); err != nil {
		t.Fatalf("setting iteration field state: %v", err)
	}
	configuration := projectV2As[map[string]any](projectV2Get[[]any](d, "iteration_configuration")[0], "iteration_configuration")
	if got := projectV2MapGet[string](configuration, "start_date"); got != "2026-01-05" {
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
