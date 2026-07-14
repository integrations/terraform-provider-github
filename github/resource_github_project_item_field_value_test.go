package github

import (
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	valueapplication "github.com/integrations/terraform-provider-github/v6/internal/application/projects/item/field/value"
)

func TestResourceGithubProjectItemFieldValueCreate(t *testing.T) {
	t.Parallel()
	client, requests := newProjectV2TestClient(t, func(query string) string {
		switch {
		case strings.Contains(query, "updateProjectV2ItemFieldValue"):
			return `{"data":{"updateProjectV2ItemFieldValue":{"projectV2Item":{"id":"PVTI_1"}}}}`
		case strings.Contains(query, "fieldValueByName"):
			return `{"data":{"node":{"__typename":"ProjectV2Item","fieldValueByName":{"__typename":"ProjectV2ItemFieldNumberValue","number":0}}}}`
		case strings.Contains(query, "node(id:"):
			return `{"data":{"node":{"__typename":"ProjectV2Field","id":"PVTF_1","name":"Estimate","dataType":"NUMBER","project":{"id":"PVT_1"}}}}`
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
	if d.Id() != "PVT_1:PVTI_1:PVTF_1" || len(*requests) != 3 {
		t.Fatalf("unexpected field value result: id=%q operations=%d", d.Id(), len(*requests))
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
