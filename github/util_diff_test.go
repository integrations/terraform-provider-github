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
				name        string
				oldID       string
				newID       cty.Value
				omitted     bool
				wantReplace bool
			}{
				{name: "retained", oldID: "network-1", newID: cty.StringVal("network-1")},
				{name: "changed", oldID: "network-1", newID: cty.StringVal("network-2")},
				{name: "assigned", newID: cty.StringVal("network-1")},
				{name: "omitted", oldID: "network-1", omitted: true},
				{name: "null", oldID: "network-1", newID: cty.NullVal(cty.String)},
				{name: "unknown", oldID: "network-1", newID: cty.UnknownVal(cty.String)},
				{name: "unassigned", omitted: true},
				{name: "cleared", oldID: "network-1", newID: cty.StringVal(""), wantReplace: true},
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
					if diff.RequiresNew() != change.wantReplace {
						t.Errorf("replacement = %t, want %t", diff.RequiresNew(), change.wantReplace)
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
