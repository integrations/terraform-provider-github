package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceGithubOrganizationRoleUsers(t *testing.T) {
	t.Run("get the organization role users without error", func(t *testing.T) {
		if testAccConf.testOrgUser == "" {
			t.Skip("Skipping test because no organization user has been configured")
		}

		roleId := 8134
		config := fmt.Sprintf(`
			resource "github_organization_role_user" "test" {
				role_id = %d
				login   = "%s"
			}

			data "github_organization_role_users" "test" {
				role_id = %[1]d

				depends_on = [
					github_organization_role_user.test
				]
			}
		`, roleId, testAccConf.testOrgUser)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("data.github_organization_role_users.test", "users.#"),
						resource.TestCheckResourceAttr("data.github_organization_role_users.test", "users.#", "1"),
						resource.TestCheckResourceAttrSet("data.github_organization_role_users.test", "users.0.user_id"),
						resource.TestCheckResourceAttr("data.github_organization_role_users.test", "users.0.login", testAccConf.testOrgUser),
					),
				},
			},
		})
	})

	t.Run("get indirect organization role users without error", func(t *testing.T) {
		if testAccConf.testOrgUser == "" {
			t.Skip("Skipping test because no organization user has been configured")
		}

		randomID := acctest.RandString(5)
		teamName := fmt.Sprintf("%steam-%s", testResourcePrefix, randomID)
		roleId := 8134
		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name    = "%s"
				privacy = "closed"
			}

			resource "github_team_membership" "test" {
				team_id  = github_team.test.id
				username = "%s"
			}

			resource "github_organization_role_team" "test" {
				role_id   = %d
				team_slug = github_team.test.slug
			}

			data "github_organization_role_users" "test" {
				role_id = %[3]d

				depends_on = [
					github_organization_role_team.test
				]
			}
		`, teamName, testAccConf.testOrgUser, roleId)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("data.github_organization_role_users.test", "users.#"),
						resource.TestCheckResourceAttr("data.github_organization_role_users.test", "users.#", "1"),
						resource.TestCheckResourceAttrSet("data.github_organization_role_users.test", "users.0.user_id"),
						resource.TestCheckResourceAttr("data.github_organization_role_users.test", "users.0.login", testAccConf.testOrgUser),
					),
				},
			},
		})
	})
}
