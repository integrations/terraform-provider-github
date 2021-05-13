package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubIpRangesDataSource(t *testing.T) {

	t.Run("reads IP ranges without error", func(t *testing.T) {

		config := `data "github_ip_ranges" "test" {}`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_ip_ranges.test", "hooks.#"),
			resource.TestCheckResourceAttrSet("data.github_ip_ranges.test", "git.#"),
			resource.TestCheckResourceAttrSet("data.github_ip_ranges.test", "pages.#"),
			resource.TestCheckResourceAttrSet("data.github_ip_ranges.test", "importer.#"),
			resource.TestCheckResourceAttrSet("data.github_ip_ranges.test", "actions.#"),
			resource.TestCheckResourceAttrSet("data.github_ip_ranges.test", "dependabot.#"),
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
