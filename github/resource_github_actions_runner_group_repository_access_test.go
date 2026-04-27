package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubActionsRunnerGroupRepositoryAccess(t *testing.T) {
	t.Run("set repo access directly and verify import", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-act-runner-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "%s"
			  vulnerability_alerts = false
			  auto_init = true
			}

			resource "github_actions_runner_group" "test" {
			  name       = github_repository.test.name
			  visibility = "selected"
			  allows_public_repositories = true
			}

			resource "github_actions_runner_group_repository_access" "test" {
				runner_group_id = github_actions_runner_group.test.id
				repository_id = github_repository.test.repo_id
			}
		`, repoName)
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_actions_runner_group_repository_access.test", "id"),
						resource.TestCheckResourceAttrSet("github_actions_runner_group_repository_access.test", "runner_group_id"),
						resource.TestCheckResourceAttrSet("github_actions_runner_group_repository_access.test", "repository_id"),
					),
				},
				{
					ResourceName:      "github_actions_runner_group_repository_access.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})
}
