package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubEnterpriseActionsHostedRunner(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates enterprise hosted runners without error", func(t *testing.T) {
		config := fmt.Sprintf(`
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
				name            = "tf-acc-test-%s"
				
				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = github_enterprise_actions_runner_group.test.id
			}
		`, testAccConf.enterpriseSlug, randomID, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_hosted_runner.test", "enterprise_slug",
				testAccConf.enterpriseSlug,
			),
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_hosted_runner.test", "name",
				fmt.Sprintf("tf-acc-test-%s", randomID),
			),
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_hosted_runner.test", "size",
				"4-core",
			),
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_hosted_runner.test", "image.0.id",
				"2306",
			),
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_hosted_runner.test", "image.0.source",
				"github",
			),
			resource.TestCheckResourceAttrSet(
				"github_enterprise_actions_hosted_runner.test", "id",
			),
			resource.TestCheckResourceAttrSet(
				"github_enterprise_actions_hosted_runner.test", "status",
			),
			resource.TestCheckResourceAttrSet(
				"github_enterprise_actions_hosted_runner.test", "platform",
			),
			resource.TestCheckResourceAttrSet(
				"github_enterprise_actions_hosted_runner.test", "image.0.size_gb",
			),
			resource.TestCheckResourceAttrSet(
				"github_enterprise_actions_hosted_runner.test", "machine_size_details.0.id",
			),
			resource.TestCheckResourceAttrSet(
				"github_enterprise_actions_hosted_runner.test", "machine_size_details.0.cpu_cores",
			),
			resource.TestCheckResourceAttrSet(
				"github_enterprise_actions_hosted_runner.test", "machine_size_details.0.memory_gb",
			),
			resource.TestCheckResourceAttrSet(
				"github_enterprise_actions_hosted_runner.test", "machine_size_details.0.storage_gb",
			),
		)

		testCase := func(t *testing.T, mode testMode) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			t.Skip("individual account not supported for enterprise hosted runners")
		})

		t.Run("with an organization account", func(t *testing.T) {
			t.Skip("organization account not supported for enterprise operations")
		})

		t.Run("with an enterprise account", func(t *testing.T) {
			testCase(t, enterprise)
		})
	})

	t.Run("updates enterprise hosted runners without error", func(t *testing.T) {
		config := fmt.Sprintf(`
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
				name            = "tf-acc-test-%s"
				
				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = github_enterprise_actions_runner_group.test.id
				maximum_runners = 5
				public_ip_enabled = false
			}
		`, testAccConf.enterpriseSlug, randomID, randomID)

		configUpdated := fmt.Sprintf(`
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
				name            = "tf-acc-test-updated-%s"
				
				image {
					id     = "2306"
					source = "github"
				}

				size            = "8-core"
				runner_group_id = github_enterprise_actions_runner_group.test.id
				maximum_runners = 10
				public_ip_enabled = true
			}
		`, testAccConf.enterpriseSlug, randomID, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_hosted_runner.test", "name",
				fmt.Sprintf("tf-acc-test-%s", randomID),
			),
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_hosted_runner.test", "size",
				"4-core",
			),
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_hosted_runner.test", "maximum_runners",
				"5",
			),
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_hosted_runner.test", "public_ip_enabled",
				"false",
			),
		)

		checkUpdated := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_hosted_runner.test", "name",
				fmt.Sprintf("tf-acc-test-updated-%s", randomID),
			),
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_hosted_runner.test", "size",
				"8-core",
			),
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_hosted_runner.test", "maximum_runners",
				"10",
			),
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_hosted_runner.test", "public_ip_enabled",
				"true",
			),
		)

		testCase := func(t *testing.T, mode testMode) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
					{
						Config: configUpdated,
						Check:  checkUpdated,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			t.Skip("individual account not supported for enterprise hosted runners")
		})

		t.Run("with an organization account", func(t *testing.T) {
			t.Skip("organization account not supported for enterprise operations")
		})

		t.Run("with an enterprise account", func(t *testing.T) {
			testCase(t, enterprise)
		})
	})

	t.Run("imports enterprise hosted runners without error", func(t *testing.T) {
		config := fmt.Sprintf(`
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
				name            = "tf-acc-test-%s"
				
				image {
					id     = "2306"
					source = "github"
				}

				size            = "4-core"
				runner_group_id = github_enterprise_actions_runner_group.test.id
			}
		`, testAccConf.enterpriseSlug, randomID, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_hosted_runner.test", "enterprise_slug",
				testAccConf.enterpriseSlug,
			),
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_hosted_runner.test", "name",
				fmt.Sprintf("tf-acc-test-%s", randomID),
			),
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_hosted_runner.test", "size",
				"4-core",
			),
		)

		testCase := func(t *testing.T, mode testMode) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
					{
						ResourceName:      "github_enterprise_actions_hosted_runner.test",
						ImportState:       true,
						ImportStateVerify: true,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			t.Skip("individual account not supported for enterprise hosted runners")
		})

		t.Run("with an organization account", func(t *testing.T) {
			t.Skip("organization account not supported for enterprise operations")
		})

		t.Run("with an enterprise account", func(t *testing.T) {
			testCase(t, enterprise)
		})
	})
}
