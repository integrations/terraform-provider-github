package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubOrganizationRoleTeamAssignment(t *testing.T) {
	// Using the predefined roles since custom roles are a strictly Enterprise feature ((https://github.blog/changelog/2024-07-10-pre-defined-organization-roles-that-grant-access-to-all-repositories/))
	githubPredefinedRoleMapping := make(map[string]string)
	githubPredefinedRoleMapping["all_repo_read"] = "8132"
	githubPredefinedRoleMapping["all_repo_triage"] = "8133"
	githubPredefinedRoleMapping["all_repo_write"] = "8134"
	githubPredefinedRoleMapping["all_repo_maintain"] = "8135"
	githubPredefinedRoleMapping["all_repo_admin"] = "8136"

	t.Run("creates repo assignment without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamSlug := fmt.Sprintf("tf-acc-test-team-repo-%s", randomID)

		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name        = "%s"
				description = "test"
			}
			resource github_organization_role_team_assignment "test" {
				team_slug = github_team.test.slug
				role_id = "%s"
			}
		`, teamSlug, githubPredefinedRoleMapping["all_repo_read"])

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_role_team_assignment.test", "id", fmt.Sprintf("%s:%s", teamSlug, githubPredefinedRoleMapping["all_repo_read"]),
			),
			resource.TestCheckResourceAttr(
				"github_organization_role_team_assignment.test", "team_slug", teamSlug,
			),
			resource.TestCheckResourceAttr(
				"github_organization_role_team_assignment.test", "role_id", githubPredefinedRoleMapping["all_repo_read"],
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessHasOrgs(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	// More tests can go here following the same format...
	t.Run("create and re-creates role assignment without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamSlug := fmt.Sprintf("tf-acc-test-team-repo-%s", randomID)

		configs := map[string]string{
			"before": fmt.Sprintf(`
				resource "github_team" "test" {
					name        = "%s"
					description = "test"
				}
				resource github_organization_role_team_assignment "test" {
					team_slug = github_team.test.slug
					role_id = "%s"
				}
		`, teamSlug, githubPredefinedRoleMapping["all_repo_read"]),
			"after": fmt.Sprintf(`
				resource "github_team" "test" {
					name        = "%s"
					description = "test"
				}
				resource github_organization_role_team_assignment "test" {
					team_slug = github_team.test.slug
					role_id = "%s"
				}
		`, teamSlug, githubPredefinedRoleMapping["all_repo_write"]),
		}

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_organization_role_team_assignment.test", "role_id", githubPredefinedRoleMapping["all_repo_read"],
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_organization_role_team_assignment.test", "role_id", githubPredefinedRoleMapping["all_repo_write"],
				),
			),
		}

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessHasOrgs(t) },
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
	})
}
