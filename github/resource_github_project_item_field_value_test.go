package github

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	valueapplication "github.com/integrations/terraform-provider-github/v6/internal/application/projects/item/field/value"
)

func TestResourceGithubProjectItemFieldValueCreate(t *testing.T) {
	t.Parallel()
	client, requests := newProjectV2TestClient(t, func(request projectV2GraphQLRequest) string {
		query := request.Query
		switch {
		case strings.Contains(query, "updateProjectV2ItemFieldValue"):
			return `{"data":{"updateProjectV2ItemFieldValue":{"projectV2Item":{"id":"PVTI_1"}}}}`
		default:
			t.Fatalf("unexpected GraphQL operation: %s", query)
			return ""
		}
	})
	resource := resourceGithubProjectItemFieldValue()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{"project_id": "PVT_1", "item_id": "PVTI_1", "field_id": "PVTF_1", "number": 0.0})
	diags := resourceGithubProjectItemFieldValueCreateOrUpdate(t.Context(), d, &Owner{v4client: client})
	if diags.HasError() {
		t.Fatalf("creating project item field value returned diagnostics: %v\nrequests: %v", diags, *requests)
	}
	if d.Id() != "PVT_1:PVTI_1:PVTF_1" || len(*requests) != 1 {
		t.Fatalf("unexpected field value result: id=%q operations=%d", d.Id(), len(*requests))
	}
	assertProjectV2GraphQLInput(t, (*requests)[0], map[string]any{"projectId": "PVT_1", "itemId": "PVTI_1", "fieldId": "PVTF_1"})
}

func TestResourceGithubProjectItemFieldValueReadUsesFieldID(t *testing.T) {
	t.Parallel()
	client, requests := newProjectV2TestClient(t, func(request projectV2GraphQLRequest) string {
		if !strings.Contains(request.Query, "fieldValues(first:") {
			t.Fatalf("unexpected GraphQL operation: %s", request.Query)
		}
		return `{"data":{"node":{"fieldValues":{"nodes":[{"__typename":"ProjectV2ItemFieldNumberValue","field":{"__typename":"ProjectV2Field","id":"PVTF_OTHER"},"number":9},{"__typename":"ProjectV2ItemFieldNumberValue","field":{"__typename":"ProjectV2Field","id":"PVTF_1"},"number":0}],"pageInfo":{"hasNextPage":false,"endCursor":null}}}}}`
	})
	resource := resourceGithubProjectItemFieldValue()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{"project_id": "PVT_1", "item_id": "PVTI_1", "field_id": "PVTF_1", "number": 5.0})
	d.SetId("PVT_1:PVTI_1:PVTF_1")
	if diagnostics := resourceGithubProjectItemFieldValueRead(t.Context(), d, &Owner{v4client: client}); diagnostics.HasError() {
		t.Fatalf("reading project item field value returned diagnostics: %v", diagnostics)
	}
	if len(*requests) != 1 || projectV2Get[float64](d, "number") != 0 {
		t.Fatalf("unexpected field value read: operations=%d number=%v", len(*requests), d.Get("number"))
	}
}

func TestExpandProjectV2ItemFieldValueSupportsZero(t *testing.T) {
	t.Parallel()
	resource := resourceGithubProjectItemFieldValue()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{"project_id": "PVT_1", "item_id": "PVTI_1", "field_id": "PVTF_1", "number": 0.0})
	value, err := expandProjectV2ItemFieldValue(d)
	if err != nil {
		t.Fatalf("expanding field value: %v", err)
	}
	if value.Kind != valueapplication.KindNumber || value.Number != 0 {
		t.Fatalf("zero number was not preserved: %#v", value)
	}
}

func TestSetProjectV2ItemFieldValueStateUsesType(t *testing.T) {
	t.Parallel()
	resource := resourceGithubProjectItemFieldValue()
	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{"project_id": "PVT_1", "item_id": "PVTI_1", "field_id": "PVTF_1", "text": "stale"})
	value := valueapplication.Result{Kind: valueapplication.KindNumber, Number: 0}
	if err := setProjectV2ItemFieldValueState(d, value); err != nil {
		t.Fatalf("setting field value state: %v", err)
	}
	if got := projectV2Get[float64](d, "number"); got != 0 {
		t.Fatalf("unexpected number %v", got)
	}
	if got := projectV2Get[string](d, "text"); got != "" {
		t.Fatalf("stale text value was preserved: %q", got)
	}
}
