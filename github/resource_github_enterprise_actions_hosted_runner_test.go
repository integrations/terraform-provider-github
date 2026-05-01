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

func testAccGithubEnterpriseActionsHostedRunnerConfig(enterpriseSlug, randomID, name, size string, extraArgs string) string {
	return fmt.Sprintf(`
		data "github_enterprise" "enterprise" {
			slug = "%s"
		}

		resource "github_enterprise_actions_runner_group" "test" {
			enterprise_slug = data.github_enterprise.enterprise.slug
			name            = "tf-acc-test-group-%s"
			visibility      = "all"
		}

		resource "github_enterprise_actions_hosted_runner" "test" {
			enterprise_slug = data.github_enterprise.enterprise.slug
			name            = "%s"

			# Image ID "2306" is the GitHub-owned Ubuntu Latest 24.04 image
			# To list available images: GET /enterprises/{enterprise}/actions/hosted-runners/images/github-owned
			image {
				id     = "2306"
				source = "github"
			}

			size            = "%s"
			runner_group_id = github_enterprise_actions_runner_group.test.id
			%s
		}
	`, enterpriseSlug, randomID, name, size, extraArgs)
}

func TestAccGithubEnterpriseActionsHostedRunner(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates enterprise hosted runners without error", func(t *testing.T) {
		config := testAccGithubEnterpriseActionsHostedRunnerConfig(
			testAccConf.enterpriseSlug, randomID,
			fmt.Sprintf("tf-acc-test-%s", randomID),
			"4-core", "",
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_actions_hosted_runner.test", tfjsonpath.New("enterprise_slug"), knownvalue.StringExact(testAccConf.enterpriseSlug)),
						statecheck.ExpectKnownValue("github_enterprise_actions_hosted_runner.test", tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-acc-test-%s", randomID))),
						statecheck.ExpectKnownValue("github_enterprise_actions_hosted_runner.test", tfjsonpath.New("size"), knownvalue.StringExact("4-core")),
						statecheck.ExpectKnownValue("github_enterprise_actions_hosted_runner.test", tfjsonpath.New("image").AtSliceIndex(0).AtMapKey("id"), knownvalue.StringExact("2306")),
						statecheck.ExpectKnownValue("github_enterprise_actions_hosted_runner.test", tfjsonpath.New("image").AtSliceIndex(0).AtMapKey("source"), knownvalue.StringExact("github")),
						statecheck.ExpectKnownValue("github_enterprise_actions_hosted_runner.test", tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_enterprise_actions_hosted_runner.test", tfjsonpath.New("status"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_enterprise_actions_hosted_runner.test", tfjsonpath.New("platform"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_enterprise_actions_hosted_runner.test", tfjsonpath.New("image").AtSliceIndex(0).AtMapKey("size_gb"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_enterprise_actions_hosted_runner.test", tfjsonpath.New("machine_size_details").AtSliceIndex(0).AtMapKey("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_enterprise_actions_hosted_runner.test", tfjsonpath.New("machine_size_details").AtSliceIndex(0).AtMapKey("cpu_cores"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_enterprise_actions_hosted_runner.test", tfjsonpath.New("machine_size_details").AtSliceIndex(0).AtMapKey("memory_gb"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_enterprise_actions_hosted_runner.test", tfjsonpath.New("machine_size_details").AtSliceIndex(0).AtMapKey("storage_gb"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("updates enterprise hosted runners without error", func(t *testing.T) {
		config := testAccGithubEnterpriseActionsHostedRunnerConfig(
			testAccConf.enterpriseSlug, randomID,
			fmt.Sprintf("tf-acc-test-%s", randomID),
			"4-core", "maximum_runners = 5\npublic_ip_enabled = false",
		)

		configUpdated := testAccGithubEnterpriseActionsHostedRunnerConfig(
			testAccConf.enterpriseSlug, randomID,
			fmt.Sprintf("tf-acc-test-updated-%s", randomID),
			"8-core", "maximum_runners = 10\npublic_ip_enabled = true",
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_actions_hosted_runner.test", tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-acc-test-%s", randomID))),
						statecheck.ExpectKnownValue("github_enterprise_actions_hosted_runner.test", tfjsonpath.New("size"), knownvalue.StringExact("4-core")),
						statecheck.ExpectKnownValue("github_enterprise_actions_hosted_runner.test", tfjsonpath.New("maximum_runners"), knownvalue.Int64Exact(5)),
						statecheck.ExpectKnownValue("github_enterprise_actions_hosted_runner.test", tfjsonpath.New("public_ip_enabled"), knownvalue.Bool(false)),
					},
				},
				{
					Config: configUpdated,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_actions_hosted_runner.test", tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-acc-test-updated-%s", randomID))),
						statecheck.ExpectKnownValue("github_enterprise_actions_hosted_runner.test", tfjsonpath.New("size"), knownvalue.StringExact("8-core")),
						statecheck.ExpectKnownValue("github_enterprise_actions_hosted_runner.test", tfjsonpath.New("maximum_runners"), knownvalue.Int64Exact(10)),
						statecheck.ExpectKnownValue("github_enterprise_actions_hosted_runner.test", tfjsonpath.New("public_ip_enabled"), knownvalue.Bool(true)),
					},
				},
			},
		})
	})

	t.Run("imports enterprise hosted runners without error", func(t *testing.T) {
		config := testAccGithubEnterpriseActionsHostedRunnerConfig(
			testAccConf.enterpriseSlug, randomID,
			fmt.Sprintf("tf-acc-test-%s", randomID),
			"4-core", "",
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_actions_hosted_runner.test", tfjsonpath.New("enterprise_slug"), knownvalue.StringExact(testAccConf.enterpriseSlug)),
						statecheck.ExpectKnownValue("github_enterprise_actions_hosted_runner.test", tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-acc-test-%s", randomID))),
						statecheck.ExpectKnownValue("github_enterprise_actions_hosted_runner.test", tfjsonpath.New("size"), knownvalue.StringExact("4-core")),
					},
				},
				{
					ResourceName:      "github_enterprise_actions_hosted_runner.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})
}
