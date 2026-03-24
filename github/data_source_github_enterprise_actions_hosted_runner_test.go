package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubEnterpriseActionsHostedRunnerDataSource(t *testing.T) {
	t.Run("gets a specific enterprise hosted runner by ID", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		config := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
				slug = "%s"
			}

			resource "github_enterprise_actions_runner_group" "test" {
				enterprise_slug = data.github_enterprise.enterprise.slug
				name            = "%srunner-group-%s"
				visibility      = "all"
			}

			resource "github_enterprise_actions_hosted_runner" "test" {
				enterprise_slug = data.github_enterprise.enterprise.slug
				name            = "%srunner-datasource-%s"
				
				image {
					# GitHub-owned Ubuntu Latest 24.04 image ID
					# To list available images: GET /enterprises/{enterprise}/actions/hosted-runners/images/github-owned
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
		`, testAccConf.enterpriseSlug, testResourcePrefix, randomID, testResourcePrefix, randomID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_enterprise_actions_hosted_runner.test", tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%srunner-datasource-%s", testResourcePrefix, randomID))),
						statecheck.ExpectKnownValue("data.github_enterprise_actions_hosted_runner.test", tfjsonpath.New("runner_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_enterprise_actions_hosted_runner.test", tfjsonpath.New("status"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_enterprise_actions_hosted_runner.test", tfjsonpath.New("platform"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_enterprise_actions_hosted_runner.test", tfjsonpath.New("image_details").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("2306")),
						statecheck.ExpectKnownValue("data.github_enterprise_actions_hosted_runner.test", tfjsonpath.New("machine_size_details").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("4-core")),
					},
				},
			},
		})
	})
}
