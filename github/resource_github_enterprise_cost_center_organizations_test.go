package github

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubEnterpriseCostCenterOrganizations(t *testing.T) {
	orgLogin := os.Getenv("ENTERPRISE_TEST_ORGANIZATION")
	if orgLogin == "" {
		t.Skip("ENTERPRISE_TEST_ORGANIZATION not set")
	}

	t.Run("manages organization assignments without error", func(t *testing.T) {
		randomID := acctest.RandString(5)

		config := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
				slug = "%s"
			}

			resource "github_enterprise_cost_center" "test" {
				enterprise_slug = data.github_enterprise.enterprise.slug
				name            = "%s%s"
			}

			resource "github_enterprise_cost_center_organizations" "test" {
				enterprise_slug     = data.github_enterprise.enterprise.slug
				cost_center_id      = github_enterprise_cost_center.test.id
				organization_logins = [%q]
			}
		`, testAccConf.enterpriseSlug, testResourcePrefix, randomID, orgLogin)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubEnterpriseCostCenterOrganizationsDestroy,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_cost_center_organizations.test", tfjsonpath.New("enterprise_slug"), knownvalue.StringExact(testAccConf.enterpriseSlug)),
						statecheck.ExpectKnownValue("github_enterprise_cost_center_organizations.test", tfjsonpath.New("organization_logins"), knownvalue.SetSizeExact(1)),
						statecheck.ExpectKnownValue("github_enterprise_cost_center_organizations.test", tfjsonpath.New("organization_logins"), knownvalue.SetPartial([]knownvalue.Check{knownvalue.StringExact(orgLogin)})),
					},
				},
				{
					ResourceName:        "github_enterprise_cost_center_organizations.test",
					ImportState:         true,
					ImportStateVerify:   true,
					ImportStateIdPrefix: testAccConf.enterpriseSlug + ":",
				},
			},
		})
	})
}

func testAccCheckGithubEnterpriseCostCenterOrganizationsDestroy(s *terraform.State) error {
	meta, err := getTestMeta()
	if err != nil {
		return err
	}
	client := meta.v3client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_enterprise_cost_center_organizations" {
			continue
		}

		enterpriseSlug := rs.Primary.Attributes["enterprise_slug"]
		costCenterID := rs.Primary.Attributes["cost_center_id"]

		cc, _, err := client.Enterprise.GetCostCenter(context.Background(), enterpriseSlug, costCenterID)
		if errIs404(err) {
			return nil
		}
		if err != nil {
			return err
		}

		// Check if organizations are still assigned
		for _, resource := range cc.Resources {
			if resource.Type == CostCenterResourceTypeOrg {
				return fmt.Errorf("cost center %s still has organization assignments", costCenterID)
			}
		}
	}

	return nil
}
