package github

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccGithubEnterpriseCostCenterUsers(t *testing.T) {
	t.Run("manages user assignments without error", func(t *testing.T) {
		randomID := acctest.RandString(5)
		user := testAccConf.username

		config := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
				slug = "%s"
			}

			resource "github_enterprise_cost_center" "test" {
				enterprise_slug = data.github_enterprise.enterprise.slug
				name            = "%s%s"
			}

			resource "github_enterprise_cost_center_users" "test" {
				enterprise_slug = data.github_enterprise.enterprise.slug
				cost_center_id  = github_enterprise_cost_center.test.id
				usernames       = [%q]
			}
		`, testAccConf.enterpriseSlug, testResourcePrefix, randomID, user)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubEnterpriseCostCenterUsersDestroy,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_enterprise_cost_center_users.test", "enterprise_slug", testAccConf.enterpriseSlug),
						resource.TestCheckResourceAttr("github_enterprise_cost_center_users.test", "usernames.#", "1"),
						resource.TestCheckTypeSetElemAttr("github_enterprise_cost_center_users.test", "usernames.*", user),
					),
				},
				{
					ResourceName:      "github_enterprise_cost_center_users.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})
}

func testAccCheckGithubEnterpriseCostCenterUsersDestroy(s *terraform.State) error {
	meta, err := getTestMeta()
	if err != nil {
		return err
	}
	client := meta.v3client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_enterprise_cost_center_users" {
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

		// Check if users are still assigned
		for _, resource := range cc.Resources {
			if resource.Type == CostCenterResourceTypeUser {
				return fmt.Errorf("cost center %s still has user assignments", costCenterID)
			}
		}
	}

	return nil
}
