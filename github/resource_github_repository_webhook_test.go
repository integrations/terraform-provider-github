package github

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccGithubRepositoryWebhook(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	var webhookOwner string = os.Getenv("GITHUB_WEBHOOK_ORGANIZATION")

	if webhookOwner == "" {
		t.Fatal("err: webhook owner must be set for this test with GITHUB_WEBHOOK_ORGANIZATION")
	}

	githubProvider := Provider()
	githubProvider.Configure(&terraform.ResourceConfig{})
	var providerOwner string = githubProvider.(*schema.Provider).Meta().(*Owner).name
	if webhookOwner == providerOwner {
		t.Fatalf("err: webhook owner %s was the same as provider owner %s; they must be different for this test (hint: use GITHUB_WEBHOOK_ORGANIZATION to set the webhook owner)", webhookOwner, providerOwner)
	}

	t.Run("creates repository webhooks without error", func(t *testing.T) {
		configs := map[string]string{
			"default": fmt.Sprintf(`
				resource "github_repository" "test" {
					name         = "test-%[1]s"
					description  = "Terraform acceptance tests"
				}

				resource "github_repository_webhook" "test" {
					depends_on = ["github_repository.test"]
					repository = "test-%[1]s"

					configuration {
						url          = "https://google.de/webhook"
						content_type = "json"
						insecure_ssl = true
					}

					events = ["pull_request"]
				}
			`, randomID),
			"withOwner": fmt.Sprintf(`
				resource "github_repository" "test" {
					name         = "test-%[1]s"
					description  = "Terraform acceptance tests"
				}

				resource "github_repository_webhook" "test" {
					depends_on = ["github_repository.test"]
					repository = "test-%[1]s"
					owner      = %[2]s

					configuration {
						url          = "https://google.de/webhook"
						content_type = "json"
						insecure_ssl = true
					}

					

					events = ["pull_request"]
				}
			`, randomID, webhookOwner),
		}

		checks := map[string]resource.TestCheckFunc{
			"default": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository_webhook.test", "active", "true",
				),
				resource.TestCheckResourceAttr(
					"github_repository_webhook.test", "events.#", "1",
				),
			),
			"withOwner": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository_webhook.test", "active", "true",
				),
				resource.TestCheckResourceAttr(
					"github_repository_webhook.test", "events.#", "1",
				),
				resource.TestCheckResourceAttr(
					"github_repository_webhook.test", "owner", webhookOwner,
				),
			),
		}

		testCase := func(t *testing.T, mode string, caseName string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: configs[caseName],
						Check:  checks[caseName],
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual, "default")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization, "default")
		})

		t.Run("with a resource-specific owner", func(t *testing.T) {
			testCase(t, enterprise, "withOwner")
		})
	})

	t.Run("imports repository webhooks without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name         = "test-%[1]s"
				description  = "Terraform acceptance tests"
			}

			resource "github_repository_webhook" "test" {
				depends_on = ["github_repository.test"]
				repository = "test-%[1]s"
				configuration {
					url          = "https://google.de/webhook"
					content_type = "json"
					insecure_ssl = true
				}
				events = ["pull_request"]
			}
			`, randomID)

		check := resource.ComposeTestCheckFunc()

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
						ResourceName:        "github_repository_webhook.test",
						ImportState:         true,
						ImportStateVerify:   true,
						ImportStateIdPrefix: fmt.Sprintf("test-%s/", randomID),
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

	t.Run("updates repository webhooks without error", func(t *testing.T) {

		configs := map[string]string{
			"before": fmt.Sprintf(`
				resource "github_repository" "test" {
				  name         = "test-%[1]s"
				  description  = "Terraform acceptance tests"
				}

				resource "github_repository_webhook" "test" {
				  depends_on = ["github_repository.test"]
				  repository = "test-%[1]s"

				  configuration {
				    url          = "https://google.de/webhook"
				    content_type = "json"
				    insecure_ssl = true
				  }

				  events = ["pull_request"]
				}
			`, randomID),
			"after": fmt.Sprintf(`
				resource "github_repository" "test" {
				  name         = "test-%[1]s"
				  description  = "Terraform acceptance tests"
				}

				resource "github_repository_webhook" "test" {
				  depends_on = ["github_repository.test"]
				  repository = "test-%[1]s"

				  configuration {
				    secret       = "secret"
				    url          = "https://google.de/webhook"
				    content_type = "json"
				    insecure_ssl = true
				  }

				  events = ["pull_request"]
				}
			`, randomID),
		}

		checks := map[string]resource.TestCheckFunc{
			"before": resource.TestCheckResourceAttr(
				"github_repository_webhook.test", "events.#", "1",
			),
			"after": resource.TestCheckResourceAttr(
				"github_repository_webhook.test", "configuration.0.secret", "secret",
			),
		}

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: configs["before"],
						Check:  checks["before"],
					},
					{
						Config: configs["after"],
						Check:  checks["after"],
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
