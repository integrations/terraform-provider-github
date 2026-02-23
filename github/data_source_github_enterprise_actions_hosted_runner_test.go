package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubEnterpriseActionsHostedRunnerDataSource(t *testing.T) {
	t.Run("gets a specific enterprise hosted runner by ID", func(t *testing.T) {
		config := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
				slug = "%s"
			}

			resource "github_enterprise_actions_runner_group" "test" {
				enterprise_slug = data.github_enterprise.enterprise.slug
				name            = "test-runner-group"
				visibility      = "all"
			}

			resource "github_enterprise_actions_hosted_runner" "test" {
				enterprise_slug = data.github_enterprise.enterprise.slug
				name            = "test-runner-for-datasource"
				
				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = github_enterprise_actions_runner_group.test.id
			}

			data "github_enterprise_actions_hosted_runner" "test" {
				enterprise_slug = data.github_enterprise.enterprise.slug
				runner_id       = github_enterprise_actions_hosted_runner.test.runner_id
			}
		`, testAccConf.enterpriseSlug)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("data.github_enterprise_actions_hosted_runner.test", "name", "test-runner-for-datasource"),
						resource.TestCheckResourceAttrSet("data.github_enterprise_actions_hosted_runner.test", "runner_id"),
						resource.TestCheckResourceAttrSet("data.github_enterprise_actions_hosted_runner.test", "status"),
						resource.TestCheckResourceAttrSet("data.github_enterprise_actions_hosted_runner.test", "platform"),
						resource.TestCheckResourceAttr("data.github_enterprise_actions_hosted_runner.test", "image_details.0.id", "2306"),
						resource.TestCheckResourceAttr("data.github_enterprise_actions_hosted_runner.test", "machine_size_details.0.id", "4-core"),
					),
				},
			},
		})
	})
}
