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

func TestAccGithubEnterpriseCostCenterRepositories(t *testing.T) {
	repoName := os.Getenv("ENTERPRISE_TEST_REPOSITORY")
	if repoName == "" {
		t.Skip("ENTERPRISE_TEST_REPOSITORY not set")
	}

	t.Run("manages repository assignments without error", func(t *testing.T) {
		randomID := acctest.RandString(5)

		config := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
				slug = "%s"
			}

			resource "github_enterprise_cost_center" "test" {
				enterprise_slug = data.github_enterprise.enterprise.slug
				name            = "%s%s"
			}

			resource "github_enterprise_cost_center_repositories" "test" {
				enterprise_slug  = data.github_enterprise.enterprise.slug
				cost_center_id   = github_enterprise_cost_center.test.id
				repository_names = [%q]
			}
		`, testAccConf.enterpriseSlug, testResourcePrefix, randomID, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubEnterpriseCostCenterRepositoriesDestroy,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_cost_center_repositories.test", tfjsonpath.New("enterprise_slug"), knownvalue.StringExact(testAccConf.enterpriseSlug)),
						statecheck.ExpectKnownValue("github_enterprise_cost_center_repositories.test", tfjsonpath.New("repository_names"), knownvalue.SetSizeExact(1)),
						statecheck.ExpectKnownValue("github_enterprise_cost_center_repositories.test", tfjsonpath.New("repository_names"), knownvalue.SetPartial([]knownvalue.Check{knownvalue.StringExact(repoName)})),
					},
				},
				{
					ResourceName:      "github_enterprise_cost_center_repositories.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})
}

func testAccCheckGithubEnterpriseCostCenterRepositoriesDestroy(s *terraform.State) error {
	meta, err := getTestMeta()
	if err != nil {
		return err
	}
	client := meta.v3client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_enterprise_cost_center_repositories" {
			continue
		}

		enterpriseSlug, costCenterID, err := parseID2(rs.Primary.ID)
		if err != nil {
			return err
		}

		cc, _, err := client.Enterprise.GetCostCenter(context.Background(), enterpriseSlug, costCenterID)
		if errIs404(err) {
			return nil
		}
		if err != nil {
			return err
		}

		// Check if repositories are still assigned
		for _, resource := range cc.Resources {
			if resource.Type == CostCenterResourceTypeRepo {
				return fmt.Errorf("cost center %s still has repository assignments", costCenterID)
			}
		}
	}

	return nil
}
