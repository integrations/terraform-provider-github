package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubOrganizationDataSource(t *testing.T) {

	t.Run("queries for an organization without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			data "github_organization" "test" {
				name = "%s"
			}
		`, testOrganization)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("data.github_organization.test", "login", testOrganization),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "name"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "orgname"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "node_id"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "description"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "plan"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "repositories.#"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "members.#"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "two_factor_requirement_enabled"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "default_repository_permission"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "members_can_create_repositories"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "members_allowed_repository_creation_type"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "members_can_create_public_repositories"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "members_can_create_private_repositories"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "members_can_create_internal_repositories"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "members_can_fork_private_repositories"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "web_commit_signoff_required"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "members_can_create_pages"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "members_can_create_public_pages"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "members_can_create_private_pages"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "advanced_security_enabled_for_new_repositories"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "dependabot_alerts_enabled_for_new_repositories"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "dependabot_security_updates_enabled_for_new_repositories"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "dependency_graph_enabled_for_new_repositories"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "secret_scanning_enabled_for_new_repositories"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "secret_scanning_push_protection_enabled_for_new_repositories"),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})
}
