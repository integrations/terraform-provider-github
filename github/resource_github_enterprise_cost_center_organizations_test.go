package github

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubEnterpriseCostCenterOrganizations(t *testing.T) {
	initialOrganization := os.Getenv("ENTERPRISE_TEST_ORGANIZATION")
	updatedOrganization := os.Getenv("ENTERPRISE_TEST_ORGANIZATION_UPDATED")
	if initialOrganization == "" || updatedOrganization == "" {
		t.Skip("ENTERPRISE_TEST_ORGANIZATION and ENTERPRISE_TEST_ORGANIZATION_UPDATED must be set")
	}
	if initialOrganization == updatedOrganization {
		t.Skip("ENTERPRISE_TEST_ORGANIZATION and ENTERPRISE_TEST_ORGANIZATION_UPDATED must identify different organizations")
	}

	t.Run("manages organization assignments without error", func(t *testing.T) {
		randomID := acctest.RandString(5)

		config := func(organizationLogin string) string {
			return fmt.Sprintf(`
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
		`, testAccConf.enterpriseSlug, testResourcePrefix, randomID, organizationLogin)
		}

		idValuesSame := statecheck.CompareValue(compare.ValuesSame())
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubEnterpriseCostCenterOrganizationsDestroy,
			Steps: []resource.TestStep{
				{
					Config: config(initialOrganization),
					ConfigStateChecks: []statecheck.StateCheck{
						idValuesSame.AddStateValue("github_enterprise_cost_center_organizations.test", tfjsonpath.New("id")),
						statecheck.ExpectKnownValue("github_enterprise_cost_center_organizations.test", tfjsonpath.New("enterprise_slug"), knownvalue.StringExact(testAccConf.enterpriseSlug)),
						statecheck.ExpectKnownValue("github_enterprise_cost_center_organizations.test", tfjsonpath.New("organization_logins"), knownvalue.SetSizeExact(1)),
						statecheck.ExpectKnownValue("github_enterprise_cost_center_organizations.test", tfjsonpath.New("organization_logins"), knownvalue.SetPartial([]knownvalue.Check{knownvalue.StringExact(initialOrganization)})),
					},
				},
				{
					Config: config(updatedOrganization),
					ConfigStateChecks: []statecheck.StateCheck{
						idValuesSame.AddStateValue("github_enterprise_cost_center_organizations.test", tfjsonpath.New("id")),
						statecheck.ExpectKnownValue("github_enterprise_cost_center_organizations.test", tfjsonpath.New("organization_logins"), knownvalue.SetSizeExact(1)),
						statecheck.ExpectKnownValue("github_enterprise_cost_center_organizations.test", tfjsonpath.New("organization_logins"), knownvalue.SetPartial([]knownvalue.Check{knownvalue.StringExact(updatedOrganization)})),
					},
				},
				{
					ResourceName:        "github_enterprise_cost_center_organizations.test",
					ImportState:         true,
					ImportStateVerify:   true,
					ImportStateIdPrefix: testAccConf.enterpriseSlug + ":",
				},
				{
					Config: fmt.Sprintf(`
						data "github_enterprise" "enterprise" {
							slug = "%s"
						}

						resource "github_enterprise_cost_center" "test" {
							enterprise_slug = data.github_enterprise.enterprise.slug
							name            = "%s%s"
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID),
					Check: testAccCheckGithubEnterpriseCostCenterOrganizationsAssignmentRemoved,
				},
			},
		})
	})
}

// testAccCheckGithubEnterpriseCostCenterOrganizationsAssignmentRemoved verifies, via the
// API, that organization assignments were removed by the resource's DeleteContext while the
// parent cost center still exists. This ensures the destroy check in
// testAccCheckGithubEnterpriseCostCenterOrganizationsDestroy isn't trivially satisfied by the
// parent cost center having been archived.
func testAccCheckGithubEnterpriseCostCenterOrganizationsAssignmentRemoved(s *terraform.State) error {
	meta, err := getTestMeta(testAccConf)
	if err != nil {
		return err
	}

	rs, ok := s.RootModule().Resources["github_enterprise_cost_center.test"]
	if !ok {
		return fmt.Errorf("github_enterprise_cost_center.test not found in state")
	}

	enterpriseSlug := rs.Primary.Attributes["enterprise_slug"]
	costCenterID := rs.Primary.ID

	cc, err := getEnterpriseCostCenter(context.Background(), meta.v3client, enterpriseSlug, costCenterID, meta.maxPerPage)
	if err != nil {
		return fmt.Errorf("verifying organization assignments were removed: %w", err)
	}
	if cc.GetState() == "deleted" {
		return fmt.Errorf("expected cost center %s to still exist after removing organization assignments", costCenterID)
	}
	for _, resource := range cc.Resources {
		if resource.Type == CostCenterResourceTypeOrg {
			return fmt.Errorf("cost center %s still has organization assignments after resource deletion", costCenterID)
		}
	}

	return nil
}

func testAccCheckGithubEnterpriseCostCenterOrganizationsDestroy(s *terraform.State) error {
	meta, err := getTestMeta(testAccConf)
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

		cc, err := getEnterpriseCostCenter(context.Background(), client, enterpriseSlug, costCenterID, meta.maxPerPage)
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
