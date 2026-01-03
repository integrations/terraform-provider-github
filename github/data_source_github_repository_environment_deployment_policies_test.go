package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubRepositoryEnvironmentDeploymentPolicies(t *testing.T) {
	t.Run("queries environment deployment policies", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-env-deploy-pol-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_repository_environment" "env" {
				repository  = github_repository.test.name
				environment = "my_env"
				deployment_branch_policy {
					protected_branches     = false
					custom_branch_policies = true
				}
			}

			resource "github_repository_environment_deployment_policy" "branch" {
				repository       = github_repository.test.name
				environment      = github_repository_environment.env.environment
				branch_pattern   = "foo"
			}

			resource "github_repository_environment_deployment_policy" "tag" {
				repository       = github_repository.test.name
				environment      = github_repository_environment.env.environment
				tag_pattern      = "bar"
			}

			data "github_repository_environment_deployment_policies" "test" {
				repository  = github_repository.test.name
				environment = github_repository_environment.env.environment

				depends_on = [github_repository_environment_deployment_policy.branch, github_repository_environment_deployment_policy.tag]
			}
	`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnauthenticated(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("data.github_repository_environment_deployment_policies.test", "policies.#", "2"),
						resource.TestCheckResourceAttr("data.github_repository_environment_deployment_policies.test", "policies.0.type", "branch"),
						resource.TestCheckResourceAttr("data.github_repository_environment_deployment_policies.test", "policies.0.pattern", "foo"),
						resource.TestCheckResourceAttr("data.github_repository_environment_deployment_policies.test", "policies.1.type", "tag"),
						resource.TestCheckResourceAttr("data.github_repository_environment_deployment_policies.test", "policies.1.pattern", "bar"),
					),
				},
			},
		})
	})
}
