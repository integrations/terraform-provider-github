package github

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubAutomatedSecurityFixes(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("enables automated security fixes without error", func(t *testing.T) {
		enabled := "enabled = false"
		updatedEnabled := "enabled = true"
		config := fmt.Sprintf(`

			resource "github_repository" "test" {
				name = "tf-acc-test-%s"
				visibility = "private"
			  	auto_init = true
				vulnerability_alerts   = true
			}


			resource "github_repository_dependabot_security_updates" "test" {
			  repository  = github_repository.test.id
			  %s
			}
		`, randomID, enabled)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository_dependabot_security_updates.test", "enabled",
					"false",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository_dependabot_security_updates.test", "enabled",
					"true",
				),
			),
		}

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  checks["before"],
					},
					{
						Config: strings.Replace(config,
							enabled,
							updatedEnabled, 1),
						Check: checks["after"],
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})

	t.Run("disables automated security fixes without error", func(t *testing.T) {
		enabled := "enabled = true"
		updatedEnabled := "enabled = false"

		config := fmt.Sprintf(`

			resource "github_repository" "test" {
				name = "tf-acc-test-%s"
				visibility = "private"
			  	auto_init = true
				vulnerability_alerts   = true
			}


			resource "github_repository_dependabot_security_updates" "test" {
			  repository  = github_repository.test.id
			  %s
			}
		`, randomID, enabled)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository_dependabot_security_updates.test", "enabled",
					"true",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository_dependabot_security_updates.test", "enabled",
					"false",
				),
			),
		}

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  checks["before"],
					},
					{
						Config: strings.Replace(config,
							enabled,
							updatedEnabled, 1),
						Check: checks["after"],
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})

	t.Run("imports automated security fixes without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "tf-acc-test-%s"
			  vulnerability_alerts   = true
			}

			resource "github_repository_dependabot_security_updates" "test" {
			  repository  = github_repository.test.id
			  enabled = false
			}
    `, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("github_repository_dependabot_security_updates.test", "repository"),
			resource.TestCheckResourceAttrSet("github_repository_dependabot_security_updates.test", "enabled"),
			resource.TestCheckResourceAttr("github_repository_dependabot_security_updates.test", "enabled", "false"),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
					{
						ResourceName:      "github_repository_dependabot_security_updates.test",
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
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}
