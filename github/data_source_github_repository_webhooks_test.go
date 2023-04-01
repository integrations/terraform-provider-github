package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubRepositoryWebhooksDataSource(t *testing.T) {
	t.Run("manages repository webhooks", func(t *testing.T) {
		repoName := fmt.Sprintf("tf-acc-test-webhooks-%s", acctest.RandString(5))

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_repository_webhook" "test" {
			  repository = github_repository.test.name

			  configuration {
			    url          = "https://google.de/webhook"
			    content_type = "json"
			    insecure_ssl = true
			  }

			  events = ["pull_request"]
			}
		 `, repoName)

		config2 := config + `
			data "github_repository_webhooks" "test" {
				repository = github_repository.test.name
			}
		`

		const resourceName = "data.github_repository_webhooks.test"
		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "webhooks.#", "1"),
			resource.TestCheckResourceAttr(resourceName, "webhooks.0.name", "web"),
			resource.TestCheckResourceAttr(resourceName, "webhooks.0.url", "https://google.de/webhook"),
			resource.TestCheckResourceAttr(resourceName, "webhooks.0.active", "true"),
			resource.TestCheckResourceAttrSet(resourceName, "webhooks.0.id"),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  resource.ComposeTestCheckFunc(),
					},
					{
						Config: config2,
						Check:  check,
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
