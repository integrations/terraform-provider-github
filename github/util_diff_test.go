package github

import (
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestRunnerGroupNetworkingDiff(t *testing.T) {
	for _, tc := range []struct {
		name     string
		resource *schema.Resource
	}{
		{"organization", resourceGithubActionsRunnerGroup()},
		{"enterprise", resourceGithubActionsEnterpriseRunnerGroup()},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for _, change := range []struct {
				name    string
				oldID   string
				newID   cty.Value
				omitted bool
				wantNew *string
			}{
				{name: "retained", oldID: "network-1", newID: cty.StringVal("network-1")},
				{name: "changed", oldID: "network-1", newID: cty.StringVal("network-2"), wantNew: new("network-2")},
				{name: "assigned", newID: cty.StringVal("network-1"), wantNew: new("network-1")},
				{name: "omitted", oldID: "network-1", omitted: true},
				{name: "null", oldID: "network-1", newID: cty.NullVal(cty.String)},
				{name: "unknown", oldID: "network-1", newID: cty.UnknownVal(cty.String)},
				{name: "unassigned", omitted: true},
				{name: "cleared", oldID: "network-1", newID: cty.StringVal(""), wantNew: new("")},
				{name: "already cleared", newID: cty.StringVal("")},
			} {
				t.Run(change.name, func(t *testing.T) {
					attributes := map[string]string{
						"name":                     "before",
						"visibility":               "all",
						"network_configuration_id": change.oldID,
						"etag":                     "previous-etag",
					}
					config := map[string]cty.Value{
						"name":       cty.StringVal("after"),
						"visibility": cty.StringVal("all"),
					}
					if !change.omitted {
						config["network_configuration_id"] = change.newID
					}
					if tc.name == "enterprise" {
						attributes["enterprise_slug"] = "test-enterprise"
						config["enterprise_slug"] = cty.StringVal("test-enterprise")
					}

					rawConfig, err := tc.resource.CoreConfigSchema().CoerceValue(cty.ObjectVal(config))
					if err != nil {
						t.Fatal(err)
					}
					diff, err := tc.resource.Diff(t.Context(), &terraform.InstanceState{
						ID:         "42",
						Attributes: attributes,
						RawConfig:  rawConfig,
					}, terraform.NewResourceConfigShimmed(rawConfig, tc.resource.CoreConfigSchema()), nil)
					if err != nil {
						t.Fatal(err)
					}
					if diff == nil || diff.Empty() {
						t.Fatal("expected a diff")
					}
					if diff.RequiresNew() {
						t.Error("network assignment changes must not replace the runner group")
					}
					if change.wantNew != nil {
						networkDiff := diff.Attributes["network_configuration_id"]
						if networkDiff == nil {
							t.Fatalf("expected network assignment diff to %q", *change.wantNew)
						}
						if networkDiff.New != *change.wantNew || networkDiff.NewComputed != (*change.wantNew == "") || networkDiff.NewRemoved {
							t.Errorf("network assignment diff = %+v, want value %q", networkDiff, *change.wantNew)
						}
					}
					if change.omitted || change.newID.IsNull() {
						if networkDiff, ok := diff.Attributes["network_configuration_id"]; ok {
							t.Errorf("unconfigured network assignment must be preserved, got diff %+v", networkDiff)
						}
					}
					if _, ok := diff.Attributes["etag"]; ok {
						t.Error("ETag diff must remain suppressed when networking diff logic is installed")
					}
				})
			}
		})
	}
}
