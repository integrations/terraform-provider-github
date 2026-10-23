package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubActionsRepositoryForkPRContributorApproval(t *testing.T) {
	policies := []string{
		"first_time_contributors_new_to_github",
		"first_time_contributors",
		"all_external_contributors",
	}

	for _, policy := range policies {
		t.Run(fmt.Sprintf("test setting approval_policy to %s", policy), func(t *testing.T) {
			randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
			repoName := fmt.Sprintf("%srepo-fork-pr-approval-%s", testResourcePrefix, randomID)
			approvalPolicy := policy
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

			check := resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_actions_repository_fork_pr_contributor_approval.test", "approval_policy", approvalPolicy,
				),
				resource.TestCheckResourceAttr(
					"github_actions_repository_fork_pr_contributor_approval.test", "repository", repoName,
				),
			)

			resource.Test(t, resource.TestCase{
				PreCheck:          func() { skipUnlessMode(t, individual, organization) },
				ProviderFactories: providerFactories,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
					{
						ResourceName:      "github_actions_repository_fork_pr_contributor_approval.test",
						ImportState:       true,
						ImportStateVerify: true,
					},
				},
			})
		})
	}
}
