package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubOrganizationWebhooksDataSource(t *testing.T) {
	t.Run("manages organization webhooks", func(t *testing.T) {
		config := `
			resource "github_organization_webhook" "test" {
			  configuration {
			    url          = "https://google.de/webhook"
			    content_type = "json"
			    insecure_ssl = true
			  }

			  events = ["pull_request"]
			}
		 `

		config2 := config + `
			data "github_organization_webhooks" "test" {}
		`

		const resourceName = "data.github_organization_webhooks.test"
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
