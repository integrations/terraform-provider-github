package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubAppDataSource(t *testing.T) {
	t.Parallel()

	t.Run("global", func(t *testing.T) {
		t.Parallel()

		config := `
			data "github_app" "test" {
				slug = "github-actions"
			}
		`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"data.github_app.test", "name",
			),
			resource.TestCheckResourceAttrSet(
				"data.github_app.test", "client_id",
			),
			resource.TestCheckResourceAttr(
				"data.github_app.test", "permissions.metadata", "read",
			),
			resource.TestCheckResourceAttrSet(
				"data.github_app.test", "events.#",
			),
		)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})
}
