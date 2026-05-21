package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubActionsOrganizationSelfHostedRunners(t *testing.T) {
	t.Run("test setting of basic self-hosted runners policy", func(t *testing.T) {
		enabledRepositories := "all"

		config := fmt.Sprintf(`
			resource "github_actions_organization_self_hosted_runners" "test" {
				enabled_repositories = "%s"
			}
		`, enabledRepositories)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_organization_self_hosted_runners.test", "enabled_repositories", enabledRepositories,
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
			},
		})
	})

	t.Run("test setting selected repositories with import", func(t *testing.T) {
		enabledRepositories := "selected"
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-selfhosted-runners-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "%[1]s"
				description = "Terraform acceptance tests %[1]s"
				topics      = ["terraform", "testing"]
			}

			resource "github_actions_organization_self_hosted_runners" "test" {
				enabled_repositories = "%s"
				enabled_repositories_config {
					repository_ids = [github_repository.test.repo_id]
				}
			}
		`, repoName, enabledRepositories)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_organization_self_hosted_runners.test", "enabled_repositories", enabledRepositories,
			),
			resource.TestCheckResourceAttr(
				"github_actions_organization_self_hosted_runners.test", "enabled_repositories_config.#", "1",
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
					ResourceName:      "github_actions_organization_self_hosted_runners.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("test updating from all to selected repositories", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-selfhosted-runners-%s", testResourcePrefix, randomID)

		configAll := `
			resource "github_actions_organization_self_hosted_runners" "test" {
				enabled_repositories = "all"
			}
		`

		configSelected := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "%[1]s"
				description = "Terraform acceptance tests %[1]s"
				topics      = ["terraform", "testing"]
			}

			resource "github_actions_organization_self_hosted_runners" "test" {
				enabled_repositories = "selected"
				enabled_repositories_config {
					repository_ids = [github_repository.test.repo_id]
				}
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configAll,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(
							"github_actions_organization_self_hosted_runners.test", "enabled_repositories", "all",
						),
					),
				},
				{
					Config: configSelected,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(
							"github_actions_organization_self_hosted_runners.test", "enabled_repositories", "selected",
						),
						resource.TestCheckResourceAttr(
							"github_actions_organization_self_hosted_runners.test", "enabled_repositories_config.#", "1",
						),
					),
				},
			},
		})
	})
}
