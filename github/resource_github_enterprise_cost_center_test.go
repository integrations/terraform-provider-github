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
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "enterprise_slug", testAccConf.enterpriseSlug),
						resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "name", testResourcePrefix+randomID),
						resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "state", "active"),
					),
				},
			},
		})
	})

	t.Run("updates cost center name without error", func(t *testing.T) {
		randomID := acctest.RandString(5)

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
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "name", testResourcePrefix+randomID),
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
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_enterprise_cost_center.test", "name", testResourcePrefix+"updated-"+randomID),
					),
				},
			},
		})
	})

	t.Run("imports cost center without error", func(t *testing.T) {
		randomID := acctest.RandString(5)

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
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID),
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
						id, err := buildID(testAccConf.enterpriseSlug, rs.Primary.ID)
						if err != nil {
							return "", err
						}
						return id, nil
					},
				},
			},
		})
	})
}
