package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubOrganizationDataSource(t *testing.T) {

	t.Run("queries for an organization without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			data "github_organization" "test" {
				name = "%s"
			}
		`, testOrganization)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("data.github_organization.test", "login", testOrganization),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "name"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "description"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "plan"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "repositories.#"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "members.#"),
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
