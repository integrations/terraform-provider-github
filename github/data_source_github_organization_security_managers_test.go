package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceGithubOrganizationSecurityManagers(t *testing.T) {
	t.Run("get the organization security managers without error", func(t *testing.T) {
		randomID := acctest.RandString(5)
		teamName := fmt.Sprintf("%steam-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name = "%s"
			}

			resource "github_organization_security_manager" "test" {
				team_slug = github_team.test.slug
			}

			data "github_organization_security_managers" "test" {
				depends_on = [
					github_organization_security_manager.test
				]
			}
		`, teamName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("data.github_organization_security_managers.test", "teams.#"),
						resource.TestCheckResourceAttr("data.github_organization_security_managers.test", "teams.#", "1"),
						resource.TestCheckResourceAttr("data.github_organization_security_managers.test", "teams.0.name", teamName),
					),
				},
			},
		})
	})
}
