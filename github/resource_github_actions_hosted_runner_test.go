package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubActionsHostedRunner(t *testing.T) {
	t.Parallel()

	t.Run("creates_hosted_runners_without_error", func(t *testing.T) {
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

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("name"), knownvalue.StringExact(hostedRunnerName)),
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("runner_group_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("size"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("status"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("platform"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("image").AtSliceIndex(0).AtMapKey("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("image").AtSliceIndex(0).AtMapKey("source"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("image").AtSliceIndex(0).AtMapKey("size_gb"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("machine_size_details").AtSliceIndex(0).AtMapKey("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("machine_size_details").AtSliceIndex(0).AtMapKey("cpu_cores"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("machine_size_details").AtSliceIndex(0).AtMapKey("memory_gb"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("machine_size_details").AtSliceIndex(0).AtMapKey("storage_gb"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("creates_hosted_runner_with_optional_parameters", func(t *testing.T) {
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

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("name"), knownvalue.StringExact(hostedRunnerName)),
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("size"), knownvalue.StringExact("2-core")),
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("maximum_runners"), knownvalue.Int64Exact(5)),
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("public_ip_enabled"), knownvalue.Bool(true)),
					},
				},
			},
		})
	})

	t.Run("updates_hosted_runner_configuration", func(t *testing.T) {
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

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configBefore,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("name"), knownvalue.StringExact(hostedRunnerName)),
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("size"), knownvalue.StringExact("4-core")),
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("maximum_runners"), knownvalue.Int64Exact(3)),
					},
				},
				{
					Config: configAfter,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_actions_hosted_runner.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("%s-updated", hostedRunnerName))),
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("size"), knownvalue.StringExact("4-core")),
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("maximum_runners"), knownvalue.Int64Exact(5)),
					},
				},
			},
		})
	})

	t.Run("updates_size_field", func(t *testing.T) {
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

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configBefore,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("size"), knownvalue.StringExact("4-core")),
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("machine_size_details").AtSliceIndex(0).AtMapKey("cpu_cores"), knownvalue.Int64Exact(4)),
					},
				},
				{
					Config: configAfter,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_actions_hosted_runner.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("size"), knownvalue.StringExact("8-core")),
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("machine_size_details").AtSliceIndex(0).AtMapKey("cpu_cores"), knownvalue.Int64Exact(8)),
					},
				},
			},
		})
	})

	t.Run("imports_hosted_runner", func(t *testing.T) {
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

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("name"), knownvalue.StringExact(hostedRunnerName)),
					},
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

	t.Run("deletes_hosted_runner", func(t *testing.T) {
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
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("id"), knownvalue.NotNull()),
					},
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
