package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/compare"
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

		runnerGroup := mustCreateTestOrganizationActionsRunnerGroup(t)
		randomID := acctest.RandString(5)
		hostedRunnerName := fmt.Sprintf("%srunner-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_actions_hosted_runner" "test" {
				name = "%s"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = "%d"
			}
		`, hostedRunnerName, runnerGroup.GetID())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("status"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("platform"), knownvalue.NotNull()),
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

		runnerGroup := mustCreateTestOrganizationActionsRunnerGroup(t)
		randomID := acctest.RandString(5)
		hostedRunnerName := fmt.Sprintf("%srunner-optional-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_actions_hosted_runner" "test" {
				name = "%s"

				image {
					id     = "2306"
					source = "github"
				}

				size              = "2-core"
				runner_group_id = "%d"
				maximum_runners   = 2
				public_ip_enabled = true
			}
		`, hostedRunnerName, runnerGroup.GetID())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
			},
		})
	})

	t.Run("updates_hosted_runner_configuration", func(t *testing.T) {
		t.Parallel()

		runnerGroup := mustCreateTestOrganizationActionsRunnerGroup(t)
		randomID := acctest.RandString(5)
		hostedRunnerName := fmt.Sprintf("%srunner-update-%s", testResourcePrefix, randomID)

		configBefore := fmt.Sprintf(`
			resource "github_actions_hosted_runner" "test" {
				name = "%s"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = "%d"
				maximum_runners = 2
			}
		`, hostedRunnerName, runnerGroup.GetID())

		configAfter := fmt.Sprintf(`
			resource "github_actions_hosted_runner" "test" {
				name = "%s-updated"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = "%d"
				maximum_runners = 3
			}
		`, hostedRunnerName, runnerGroup.GetID())

		compareMaxRunnersUpdated := statecheck.CompareValue(compare.ValuesDiffer())
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configBefore,
					ConfigStateChecks: []statecheck.StateCheck{
						compareMaxRunnersUpdated.AddStateValue("github_actions_hosted_runner.test", tfjsonpath.New("maximum_runners")),
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
						compareMaxRunnersUpdated.AddStateValue("github_actions_hosted_runner.test", tfjsonpath.New("maximum_runners")),
					},
				},
			},
		})
	})

	t.Run("updates_size_field", func(t *testing.T) {
		t.Parallel()

		runnerGroup := mustCreateTestOrganizationActionsRunnerGroup(t)
		randomID := acctest.RandString(5)
		hostedRunnerName := fmt.Sprintf("%srunner-size-%s", testResourcePrefix, randomID)

		configBefore := fmt.Sprintf(`
			resource "github_actions_hosted_runner" "test" {
				name = "%s"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = "%d"
			}
		`, hostedRunnerName, runnerGroup.GetID())

		configAfter := fmt.Sprintf(`
			resource "github_actions_hosted_runner" "test" {
				name = "%s"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "8-core"
				runner_group_id = "%d"
			}
		`, hostedRunnerName, runnerGroup.GetID())

		compareSizeUpdated := statecheck.CompareValue(compare.ValuesDiffer())
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configBefore,
					ConfigStateChecks: []statecheck.StateCheck{
						compareSizeUpdated.AddStateValue("github_actions_hosted_runner.test", tfjsonpath.New("size")),
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
						compareSizeUpdated.AddStateValue("github_actions_hosted_runner.test", tfjsonpath.New("size")),
						statecheck.ExpectKnownValue("github_actions_hosted_runner.test", tfjsonpath.New("machine_size_details").AtSliceIndex(0).AtMapKey("cpu_cores"), knownvalue.Int64Exact(8)),
					},
				},
			},
		})
	})

	t.Run("imports_hosted_runner", func(t *testing.T) {
		t.Parallel()

		runnerGroup := mustCreateTestOrganizationActionsRunnerGroup(t)
		randomID := acctest.RandString(5)
		hostedRunnerName := fmt.Sprintf("%srunner-import-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_actions_hosted_runner" "test" {
				name = "%s"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = "%d"
			}
		`, hostedRunnerName, runnerGroup.GetID())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					ResourceName:            "github_actions_hosted_runner.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"status", "image.0.size_gb", "image_gen"},
				},
			},
		})
	})

	t.Run("deletes_hosted_runner", func(t *testing.T) {
		t.Parallel()

		runnerGroup := mustCreateTestOrganizationActionsRunnerGroup(t)
		randomID := acctest.RandString(5)
		hostedRunnerName := fmt.Sprintf("%srunner-delete-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_actions_hosted_runner" "test" {
				name = "%s"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = "%d"
			}
		`, hostedRunnerName, runnerGroup.GetID())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					Config:  config,
					Destroy: true,
				},
			},
		})
	})

	t.Run("validates_image_version_only_allowed_custom_image", func(t *testing.T) {
		t.Parallel()

		runnerGroup := mustCreateTestOrganizationActionsRunnerGroup(t)
		randomID := acctest.RandString(5)
		hostedRunnerName := fmt.Sprintf("%simage-version-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_actions_hosted_runner" "test" {
	name = "%s"

	image {
		id     = "2306"
		source = "github"
	}
	image_version = "1.0.0"
	size            = "4-core"
				runner_group_id = "%d"
}`, hostedRunnerName, runnerGroup.GetID())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile("`image_version` can only be set when `image\\[0\\].source` is 'custom'"),
				},
			},
		})
	})
}
