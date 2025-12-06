package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubOrganizationAppInstallations(t *testing.T) {
	t.Run("queries without error", func(t *testing.T) {
		config := `data "github_organization_app_installations" "test" {}`

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_organization_app_installations.test", "installations.0.id"),
			resource.TestCheckResourceAttrSet("data.github_organization_app_installations.test", "installations.0.slug"),
			resource.TestCheckResourceAttrSet("data.github_organization_app_installations.test", "installations.0.app_id"),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
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
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}
