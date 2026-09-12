package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubActionsOrganizationForkPRContributorApproval(t *testing.T) {
	policies := []string{
		"first_time_contributors_new_to_github",
		"first_time_contributors",
		"all_external_contributors",
	}

	for _, policy := range policies {
		t.Run(fmt.Sprintf("test setting org approval_policy to %s", policy), func(t *testing.T) {
			approvalPolicy := policy
			config := fmt.Sprintf(`
				resource "github_actions_organization_fork_pr_contributor_approval" "test" {
					approval_policy = "%s"
				}
			`, approvalPolicy)

			check := resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_actions_organization_fork_pr_contributor_approval.test", "approval_policy", approvalPolicy,
				),
			)

			resource.Test(t, resource.TestCase{
				PreCheck:          func() { skipUnlessHasOrgs(t) },
				ProviderFactories: providerFactories,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
					{
						ResourceName:      "github_actions_organization_fork_pr_contributor_approval.test",
						ImportState:       true,
						ImportStateVerify: true,
					},
				},
			})
		})
	}
}
