package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubActionsRepositoryForkPRContributorApprovalDataSource(t *testing.T) {
	t.Run("read the repository fork PR contributor approval policy", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-fork-pr-approval-ds-%s", testResourcePrefix, randomID)
		approvalPolicy := "all_external_contributors"

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "%[1]s"
				description = "Terraform acceptance tests %[1]s"
				topics		= ["terraform", "testing"]
				visibility  = "public"
			}

			resource "github_actions_repository_fork_pr_contributor_approval" "test" {
				approval_policy = "%[2]s"
				repository      = github_repository.test.name
			}
		`, repoName, approvalPolicy)

		config2 := config + `
			data "github_actions_repository_fork_pr_contributor_approval" "test" {
				repository = github_actions_repository_fork_pr_contributor_approval.test.repository
			}
		`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"data.github_actions_repository_fork_pr_contributor_approval.test", "approval_policy", approvalPolicy,
			),
			resource.TestCheckResourceAttr(
				"data.github_actions_repository_fork_pr_contributor_approval.test", "repository", repoName,
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, individual, organization) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  resource.ComposeTestCheckFunc(),
				},
				{
					Config: config2,
					Check:  check,
				},
			},
		})
	})
}
