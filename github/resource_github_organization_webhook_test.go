package github

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubOrganizationWebhook(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates and updates webhooks without error", func(t *testing.T) {
		config := fmt.Sprintf(`

			resource "github_repository" "test" {
			  name = "tf-acc-test-%s"
				auto_init = true
			}

			resource "github_organization_webhook" "test" {
			  configuration {
			    url = "https://google.de/webhook"
			    content_type = "json"
			    insecure_ssl = true
			  }

			  events = ["pull_request"]
			}

		`, randomID)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_organization_webhook.test", "configuration.0.url",
					"https://google.de/webhook",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_organization_webhook.test", "configuration.0.url",
					"https://google.de/updated",
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
							"https://google.de/webhook",
							"https://google.de/updated", 1),
						Check: checks["after"],
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})

	t.Run("imports webhooks without error", func(t *testing.T) {
		config := fmt.Sprintf(`

			resource "github_repository" "test" {
			  name = "tf-acc-test-%s"
				auto_init = true
			}

			resource "github_organization_webhook" "test" {
			  configuration {
			    url = "https://google.de/import"
			    content_type = "json"
			    insecure_ssl = true
			  }

			  events = ["issues"]
			}

		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_webhook.test", "events.#",
				"1",
			),
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
						ResourceName:      "github_organization_webhook.test",
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
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}
