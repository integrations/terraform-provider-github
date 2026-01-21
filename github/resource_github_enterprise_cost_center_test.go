package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccGithubEnterpriseCostCenter(t *testing.T) {
	t.Run("creates cost center without error", func(t *testing.T) {
		randomID := acctest.RandString(5)
		user := testAccConf.username

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(`
						data "github_enterprise" "enterprise" {
							slug = "%s"
						}

						resource "github_enterprise_cost_center" "test" {
							enterprise_slug = data.github_enterprise.enterprise.slug
							name            = "%s%s"

							users = [%q]
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID, user),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "enterprise_slug", testAccConf.enterpriseSlug),
						resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "name", testResourcePrefix+randomID),
						resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "state", "active"),
						resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "users.#", "1"),
						resource.TestCheckTypeSetElemAttr("github_enterprise_cost_center.test", "users.*", user),
					),
				},
			},
		})
	})

	t.Run("updates cost center without error", func(t *testing.T) {
		randomID := acctest.RandString(5)
		user := testAccConf.username

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(`
						data "github_enterprise" "enterprise" {
							slug = "%s"
						}

						resource "github_enterprise_cost_center" "test" {
							enterprise_slug = data.github_enterprise.enterprise.slug
							name            = "%s%s"

							users = [%q]
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID, user),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "name", testResourcePrefix+randomID),
						resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "users.#", "1"),
					),
				},
				{
					Config: fmt.Sprintf(`
						data "github_enterprise" "enterprise" {
							slug = "%s"
						}

						resource "github_enterprise_cost_center" "test" {
							enterprise_slug = data.github_enterprise.enterprise.slug
							name            = "%supdated-%s"

							users = [%q]
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID, user),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "name", testResourcePrefix+"updated-"+randomID),
						resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "users.#", "1"),
						resource.TestCheckTypeSetElemAttr("github_enterprise_cost_center.test", "users.*", user),
					),
				},
				{
					Config: fmt.Sprintf(`
						data "github_enterprise" "enterprise" {
							slug = "%s"
						}

						resource "github_enterprise_cost_center" "test" {
							enterprise_slug = data.github_enterprise.enterprise.slug
							name            = "%s%s"

							users = []
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "users.#", "0"),
					),
				},
			},
		})
	})

	t.Run("imports cost center without error", func(t *testing.T) {
		randomID := acctest.RandString(5)
		user := testAccConf.username

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(`
						data "github_enterprise" "enterprise" {
							slug = "%s"
						}

						resource "github_enterprise_cost_center" "test" {
							enterprise_slug = data.github_enterprise.enterprise.slug
							name            = "%s%s"

							users = [%q]
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID, user),
				},
				{
					ResourceName:      "github_enterprise_cost_center.test",
					ImportState:       true,
					ImportStateVerify: true,
					ImportStateIdFunc: func(s *terraform.State) (string, error) {
						rs, ok := s.RootModule().Resources["github_enterprise_cost_center.test"]
						if !ok {
							return "", fmt.Errorf("resource not found in state")
						}
						return buildTwoPartID(testAccConf.enterpriseSlug, rs.Primary.ID), nil
					},
				},
			},
		})
	})
}
