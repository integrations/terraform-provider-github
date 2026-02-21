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
			ConfigStateChecks: []statecheck.StateCheck{
				statecheck.CompareValuePairs("data.github_enterprise_cost_center.test", tfjsonpath.New("cost_center_id"), "github_enterprise_cost_center.test", tfjsonpath.New("id"), compare.ValuesSame()),
				statecheck.CompareValuePairs("data.github_enterprise_cost_center.test", tfjsonpath.New("name"), "github_enterprise_cost_center.test", tfjsonpath.New("name"), compare.ValuesSame()),
				statecheck.ExpectKnownValue("data.github_enterprise_cost_center.test", tfjsonpath.New("state"), knownvalue.StringExact("active")),
			},
		}},
	})
}
