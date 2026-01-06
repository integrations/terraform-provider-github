package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubActionsHostedRunner(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates hosted runners without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
				name       = "tf-acc-test-group-%s"
				visibility = "all"
			}

			resource "github_actions_hosted_runner" "test" {
				name = "tf-acc-test-%s"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = github_actions_runner_group.test.id
			}
		`, randomID, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "name",
				fmt.Sprintf("tf-acc-test-%s", randomID),
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
			PreCheck:  func() { skipUnlessHasPaidOrgs(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("creates hosted runner with optional parameters", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
				name       = "tf-acc-test-group-%s"
				visibility = "all"
			}

			resource "github_actions_hosted_runner" "test" {
				name = "tf-acc-test-optional-%s"

				image {
					id     = "2306"
					source = "github"
				}

				size              = "2-core"
				runner_group_id   = github_actions_runner_group.test.id
				maximum_runners   = 5
				public_ip_enabled = true
			}
		`, randomID, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "name",
				fmt.Sprintf("tf-acc-test-optional-%s", randomID),
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
			PreCheck:  func() { skipUnlessHasPaidOrgs(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("updates hosted runner configuration", func(t *testing.T) {
		configBefore := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
				name       = "tf-acc-test-group-%s"
				visibility = "all"
			}

			resource "github_actions_hosted_runner" "test" {
				name = "tf-acc-test-update-%s"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = github_actions_runner_group.test.id
				maximum_runners = 3
			}
		`, randomID, randomID)

		configAfter := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
				name       = "tf-acc-test-group-%s"
				visibility = "all"
			}

			resource "github_actions_hosted_runner" "test" {
				name = "tf-acc-test-update-%s-updated"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = github_actions_runner_group.test.id
				maximum_runners = 5
			}
		`, randomID, randomID)

		checkBefore := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "name",
				fmt.Sprintf("tf-acc-test-update-%s", randomID),
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
				fmt.Sprintf("tf-acc-test-update-%s-updated", randomID),
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
			PreCheck:  func() { skipUnlessHasPaidOrgs(t) },
			Providers: testAccProviders,
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
		configBefore := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
				name       = "tf-acc-test-group-%s"
				visibility = "all"
			}

			resource "github_actions_hosted_runner" "test" {
				name = "tf-acc-test-size-%s"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = github_actions_runner_group.test.id
			}
		`, randomID, randomID)

		configAfter := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
				name       = "tf-acc-test-group-%s"
				visibility = "all"
			}

			resource "github_actions_hosted_runner" "test" {
				name = "tf-acc-test-size-%s"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "8-core"
				runner_group_id = github_actions_runner_group.test.id
			}
		`, randomID, randomID)

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
			PreCheck:  func() { skipUnlessHasPaidOrgs(t) },
			Providers: testAccProviders,
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
		config := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
				name       = "tf-acc-test-group-%s"
				visibility = "all"
			}

			resource "github_actions_hosted_runner" "test" {
				name = "tf-acc-test-import-%s"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = github_actions_runner_group.test.id
			}
		`, randomID, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"github_actions_hosted_runner.test", "id",
			),
			resource.TestCheckResourceAttr(
				"github_actions_hosted_runner.test", "name",
				fmt.Sprintf("tf-acc-test-import-%s", randomID),
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessHasPaidOrgs(t) },
			Providers: testAccProviders,
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
		config := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
				name       = "tf-acc-test-group-%s"
				visibility = "all"
			}

			resource "github_actions_hosted_runner" "test" {
				name = "tf-acc-test-delete-%s"

				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = github_actions_runner_group.test.id
			}
		`, randomID, randomID)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessHasPaidOrgs(t) },
			Providers: testAccProviders,
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
								name       = "tf-acc-test-group-%s"
								visibility = "all"
							}
						`, randomID),
				},
			},
		})
	})
}
