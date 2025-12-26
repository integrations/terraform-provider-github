package github

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccGithubEnterpriseCostCentersDataSource(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	if isEnterprise != "true" {
		t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
	}
	if testEnterprise == "" {
		t.Skip("Skipping because `ENTERPRISE_SLUG` is not set")
	}

	config := fmt.Sprintf(`
		data "github_enterprise" "enterprise" {
			slug = "%s"
		}

		resource "github_enterprise_cost_center" "test" {
			enterprise_slug = data.github_enterprise.enterprise.slug
			name            = "tf-acc-test-%s"
		}

		data "github_enterprise_cost_centers" "test" {
			enterprise_slug = data.github_enterprise.enterprise.slug
			state           = "active"
			depends_on      = [github_enterprise_cost_center.test]
		}
	`, testEnterprise, randomID)

	check := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr("data.github_enterprise_cost_centers.test", "state", "active"),
		testAccCheckEnterpriseCostCentersListContains("github_enterprise_cost_center.test", "data.github_enterprise_cost_centers.test"),
	)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipUnlessMode(t, enterprise) },
		Providers: testAccProviders,
		Steps:     []resource.TestStep{{Config: config, Check: check}},
	})
}

func testAccCheckEnterpriseCostCentersListContains(costCenterResourceName, dataSourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		cc, ok := s.RootModule().Resources[costCenterResourceName]
		if !ok {
			return fmt.Errorf("resource %q not found in state", costCenterResourceName)
		}
		ccID := cc.Primary.ID
		if ccID == "" {
			return fmt.Errorf("resource %q has empty ID", costCenterResourceName)
		}

		ds, ok := s.RootModule().Resources[dataSourceName]
		if !ok {
			return fmt.Errorf("data source %q not found in state", dataSourceName)
		}

		for k, v := range ds.Primary.Attributes {
			if strings.HasPrefix(k, "cost_centers.") && strings.HasSuffix(k, ".id") {
				if v == ccID {
					return nil
				}
			}
		}

		return fmt.Errorf("expected cost center id %q to be present in %q", ccID, dataSourceName)
	}
}
