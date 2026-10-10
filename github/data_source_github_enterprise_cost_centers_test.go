package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubEnterpriseCostCentersDataSource(t *testing.T) {
	skipUnlessEnterprise(t)

	cc := mustCreateTestEnterpriseCostCenter(t)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { skipUnlessEnterprise(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{{
			Config: fmt.Sprintf(`
				data "github_enterprise_cost_centers" "test" {
					enterprise_slug = "%s"
					state           = "active"
				}
			`, testAccConf.enterpriseSlug),
			ConfigStateChecks: []statecheck.StateCheck{
				statecheck.ExpectKnownValue("data.github_enterprise_cost_centers.test", tfjsonpath.New("id"), knownvalue.StringExact(testAccConf.enterpriseSlug+idSeparator+"active")),
				statecheck.ExpectKnownValue("data.github_enterprise_cost_centers.test", tfjsonpath.New("enterprise_slug"), knownvalue.StringExact(testAccConf.enterpriseSlug)),
				statecheck.ExpectKnownValue("data.github_enterprise_cost_centers.test", tfjsonpath.New("state"), knownvalue.StringExact("active")),
				statecheck.ExpectKnownValue("data.github_enterprise_cost_centers.test", tfjsonpath.New("cost_centers"), knownvalue.SetPartial([]knownvalue.Check{
					knownvalue.ObjectExact(map[string]knownvalue.Check{
						"id":                 knownvalue.StringExact(cc.ID),
						"name":               knownvalue.StringExact(cc.Name),
						"state":              knownvalue.StringExact("active"),
						"azure_subscription": knownvalue.StringExact(cc.GetAzureSubscription()),
					}),
				})),
			},
		}},
	})
}
