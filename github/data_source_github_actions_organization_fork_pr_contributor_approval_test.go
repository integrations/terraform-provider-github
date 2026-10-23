package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubActionsOrganizationForkPRContributorApprovalDataSource(t *testing.T) {
	t.Run("read the organization fork PR contributor approval policy", func(t *testing.T) {
		approvalPolicy := "all_external_contributors"

		config := `
			resource "github_actions_organization_fork_pr_contributor_approval" "test" {
				approval_policy = "all_external_contributors"
			}
		`

		config2 := config + `
			data "github_actions_organization_fork_pr_contributor_approval" "test" {}
		`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"data.github_actions_organization_fork_pr_contributor_approval.test", "approval_policy", approvalPolicy,
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
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
