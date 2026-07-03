package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubActionsHostedRunner(t *testing.T) {
	t.Parallel()

	t.Run("creates hosted runners without error", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandString(5)
		runnerGroupName := fmt.Sprintf("%sgroup-%s", testResourcePrefix, randomID)
		hostedRunnerName := fmt.Sprintf("%srunner-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
				name       = "%s"
				visibility = "all"
			}

			resource "github_actions_hosted_runner" "test" {
				name = "%s"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = github_actions_runner_group.test.id
			}
		`, runnerGroupName, hostedRunnerName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "name",
				hostedRunnerName,
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "size",
				"4-core",
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "image.0.id",
				"2306",
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "image.0.source",
				"github",
			),
			resource.TestCheckResourceAttrSet(
				"github_actions_hosted_runner.test", "id",
			),
			resource.TestCheckResourceAttrSet(
				"github_actions_hosted_runner.test", "status",
			),
			resource.TestCheckResourceAttrSet(
				"github_actions_hosted_runner.test", "platform",
			),
			resource.TestCheckResourceAttrSet(
				"github_actions_hosted_runner.test", "image.0.size_gb",
			),
			resource.TestCheckResourceAttrSet(
				"github_actions_hosted_runner.test", "machine_size_details.0.id",
			),
			resource.TestCheckResourceAttrSet(
				"github_actions_hosted_runner.test", "machine_size_details.0.cpu_cores",
			),
			resource.TestCheckResourceAttrSet(
				"github_actions_hosted_runner.test", "machine_size_details.0.memory_gb",
			),
			resource.TestCheckResourceAttrSet(
				"github_actions_hosted_runner.test", "machine_size_details.0.storage_gb",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("creates hosted runner with optional parameters", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandString(5)
		runnerGroupName := fmt.Sprintf("%sgroup-%s", testResourcePrefix, randomID)
		hostedRunnerName := fmt.Sprintf("%srunner-optional-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
				name       = "%s"
				visibility = "all"
			}

			resource "github_actions_hosted_runner" "test" {
				name = "%s"

				image {
					id     = "2306"
					source = "github"
				}

				size              = "2-core"
				runner_group_id   = github_actions_runner_group.test.id
				maximum_runners   = 5
				public_ip_enabled = true
			}
		`, runnerGroupName, hostedRunnerName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "name",
				hostedRunnerName,
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "size",
				"2-core",
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "maximum_runners",
				"5",
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "public_ip_enabled",
				"true",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("updates hosted runner configuration", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandString(5)
		runnerGroupName := fmt.Sprintf("%sgroup-%s", testResourcePrefix, randomID)
		hostedRunnerName := fmt.Sprintf("%srunner-update-%s", testResourcePrefix, randomID)

		configBefore := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
				name       = "%s"
				visibility = "all"
			}

			resource "github_actions_hosted_runner" "test" {
				name = "%s"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = github_actions_runner_group.test.id
				maximum_runners = 3
			}
		`, runnerGroupName, hostedRunnerName)

		configAfter := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
				name       = "%s"
				visibility = "all"
			}

			resource "github_actions_hosted_runner" "test" {
				name = "%s-updated"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = github_actions_runner_group.test.id
				maximum_runners = 5
			}
		`, runnerGroupName, hostedRunnerName)

		checkBefore := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "name",
				hostedRunnerName,
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "size",
				"4-core",
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "maximum_runners",
				"3",
			),
		)

		checkAfter := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "name",
				fmt.Sprintf("%s-updated", hostedRunnerName),
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "size",
				"4-core",
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "maximum_runners",
				"5",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configBefore,
					Check:  checkBefore,
				},
				{
					Config: configAfter,
					Check:  checkAfter,
				},
			},
		})
	})

	t.Run("updates size field", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandString(5)
		runnerGroupName := fmt.Sprintf("%sgroup-%s", testResourcePrefix, randomID)
		hostedRunnerName := fmt.Sprintf("%srunner-size-%s", testResourcePrefix, randomID)

		configBefore := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
				name       = "%s"
				visibility = "all"
			}

			resource "github_actions_hosted_runner" "test" {
				name = "%s"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = github_actions_runner_group.test.id
			}
		`, runnerGroupName, hostedRunnerName)

		configAfter := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
				name       = "%s"
				visibility = "all"
			}

			resource "github_actions_hosted_runner" "test" {
				name = "%s"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "8-core"
				runner_group_id = github_actions_runner_group.test.id
			}
		`, runnerGroupName, hostedRunnerName)

		checkBefore := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "size",
				"4-core",
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "machine_size_details.0.cpu_cores",
				"4",
			),
		)

		checkAfter := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "size",
				"8-core",
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "machine_size_details.0.cpu_cores",
				"8",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configBefore,
					Check:  checkBefore,
				},
				{
					Config: configAfter,
					Check:  checkAfter,
				},
			},
		})
	})

	t.Run("imports hosted runner", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandString(5)
		runnerGroupName := fmt.Sprintf("%sgroup-%s", testResourcePrefix, randomID)
		hostedRunnerName := fmt.Sprintf("%srunner-import-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
				name       = "%s"
				visibility = "all"
			}

			resource "github_actions_hosted_runner" "test" {
				name = "%s"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = github_actions_runner_group.test.id
			}
		`, runnerGroupName, hostedRunnerName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"github_actions_hosted_runner.test", "id",
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "name",
				hostedRunnerName,
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
				{
					ResourceName:            "github_actions_hosted_runner.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"image", "image_gen"},
				},
			},
		})
	})

	t.Run("deletes hosted runner", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandString(5)
		runnerGroupName := fmt.Sprintf("%sgroup-%s", testResourcePrefix, randomID)
		hostedRunnerName := fmt.Sprintf("%srunner-delete-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
				name       = "%s"
				visibility = "all"
			}

			resource "github_actions_hosted_runner" "test" {
				name = "%s"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = github_actions_runner_group.test.id
			}
		`, runnerGroupName, hostedRunnerName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet(
							"github_actions_hosted_runner.test", "id",
						),
					),
				},
				// This step should successfully delete the runner
				{
					Config: fmt.Sprintf(`
							resource "github_actions_runner_group" "test" {
								name       = "%s"
								visibility = "all"
							}
						`, runnerGroupName),
				},
			},
		})
	})
}
