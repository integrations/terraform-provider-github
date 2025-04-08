package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubTeamOrganizationRoleAssignment(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	// Using the predefined roles since custom roles are a strictly Enterprise feature ((https://github.blog/changelog/2024-07-10-pre-defined-organization-roles-that-grant-access-to-all-repositories/))
	allRepoReadRoleName := "all_repo_read"
	allRepoWriteRoleName := "all_repo_write"

	t.Run("creates repo assignment without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name        = "tf-acc-test-team-repo-%s"
				description = "test"
			}
			resource "github_team_organization_role_assignment" "test" {
				team_id = github_team.test.id
				role_id = "%s"
			}
		`, randomID, allRepoReadRoleName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"github_team_organization_role_assignment.test", "id",
			),
			resource.TestCheckResourceAttrSet(
				"github_team_organization_role_assignment.test", "team_id",
			),
			resource.TestCheckResourceAttr(
				"github_team_organization_role_assignment.test", "role_id", allRepoReadRoleName,
			),
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
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})

	// More tests can go here following the same format...
	t.Run("create and re-creates role assignment without error", func(t *testing.T) {

		configs := map[string]string{
			"before": fmt.Sprintf(`
				resource "github_team" "test" {
					name        = "tf-acc-test-team-repo-%s"
					description = "test"
				}
				resource "github_team_organization_role_assignment" "test" {
					team_id = github_team.test.id
					role_id = "%s"
				}
		`, randomID, allRepoReadRoleName),
			"after": fmt.Sprintf(`
				resource "github_team" "test" {
					name        = "tf-acc-test-team-repo-%s"
					description = "test"
				}
				resource "github_team_organization_role_assignment" "test" {
					team_id = github_team.test.id
					role_id = "%s"
				}
		`, randomID, allRepoWriteRoleName),
		}

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_team_organization_role_assignment.test", "role_id", allRepoReadRoleName,
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_team_organization_role_assignment.test", "role_id", allRepoWriteRoleName,
				),
			),
		}

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: configs["before"],
						Check:  checks["before"],
					},
					{
						Config: configs["after"],
						Check:  checks["after"],
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}
