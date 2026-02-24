package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubEnterpriseCostCentersDataSource(t *testing.T) {
	randomID := acctest.RandString(5)

	config := fmt.Sprintf(`
		data "github_enterprise" "enterprise" {
			slug = "%s"
		}

		resource "github_enterprise_cost_center" "test" {
			enterprise_slug = data.github_enterprise.enterprise.slug
			name            = "%s%s"
		}

		data "github_enterprise_cost_centers" "test" {
			enterprise_slug = data.github_enterprise.enterprise.slug
			state           = "active"
			depends_on      = [github_enterprise_cost_center.test]
		}
	`, testAccConf.enterpriseSlug, testResourcePrefix, randomID)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { skipUnlessEnterprise(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{{
			Config: config,
			ConfigStateChecks: []statecheck.StateCheck{
				statecheck.ExpectKnownValue("data.github_enterprise_cost_centers.test", tfjsonpath.New("state"), knownvalue.StringExact("active")),
				statecheck.CompareValueCollection("data.github_enterprise_cost_centers.test", []tfjsonpath.Path{tfjsonpath.New("cost_centers"), tfjsonpath.New("id")}, "github_enterprise_cost_center.test", tfjsonpath.New("id"), compare.ValuesSame()),
			},
		}},
	})
}
