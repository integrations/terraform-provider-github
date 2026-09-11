package github

import (
	"testing"

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
			for _, networkID := range []string{"network-1", "network-2"} {
				t.Run(networkID, func(t *testing.T) {
					attributes := map[string]string{
						"name":                     "before",
						"visibility":               "all",
						"network_configuration_id": "network-1",
						"etag":                     "previous-etag",
					}
					config := map[string]any{
						"name":                     "after",
						"visibility":               "all",
						"network_configuration_id": networkID,
					}
					if tc.name == "enterprise" {
						attributes["enterprise_slug"] = "test-enterprise"
						config["enterprise_slug"] = "test-enterprise"
					}

					diff, err := tc.resource.Diff(t.Context(), &terraform.InstanceState{
						ID:         "42",
						Attributes: attributes,
					}, terraform.NewResourceConfigRaw(config), nil)
					if err != nil {
						t.Fatal(err)
					}
					if diff == nil || diff.Empty() {
						t.Fatal("expected an in-place update")
					}
					if diff.RequiresNew() {
						t.Error("setting or retaining a network configuration must not replace the runner group")
					}
					if _, ok := diff.Attributes["etag"]; ok {
						t.Error("ETag diff must remain suppressed when networking diff logic is installed")
					}
				})
			}
		})
	}
}
