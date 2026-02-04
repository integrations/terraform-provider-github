package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubEnterpriseCostCenterDataSource(t *testing.T) {
	randomID := acctest.RandString(5)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { skipUnlessEnterprise(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{{
			Config: fmt.Sprintf(`
				data "github_enterprise" "enterprise" {
					slug = "%s"
				}

				resource "github_enterprise_cost_center" "test" {
					enterprise_slug = data.github_enterprise.enterprise.slug
					name            = "%s%s"
				}

				data "github_enterprise_cost_center" "test" {
					enterprise_slug = data.github_enterprise.enterprise.slug
					cost_center_id  = github_enterprise_cost_center.test.id
				}
			`, testAccConf.enterpriseSlug, testResourcePrefix, randomID),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrPair("data.github_enterprise_cost_center.test", "cost_center_id", "github_enterprise_cost_center.test", "id"),
				resource.TestCheckResourceAttrPair("data.github_enterprise_cost_center.test", "name", "github_enterprise_cost_center.test", "name"),
				resource.TestCheckResourceAttr("data.github_enterprise_cost_center.test", "state", "active"),
			),
		}},
	})
}
