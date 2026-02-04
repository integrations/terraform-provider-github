package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("data.github_enterprise_cost_centers.test", "state", "active"),
				resource.TestCheckTypeSetElemAttrPair("data.github_enterprise_cost_centers.test", "cost_centers.*.id", "github_enterprise_cost_center.test", "id"),
			),
		}},
	})
}
