package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubActionsOrganizationPublicKeyDataSource(t *testing.T) {

	t.Run("queries an organization public key without error", func(t *testing.T) {

		config := `
			data "github_actions_organization_public_key" "test" {}
		`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"data.github_actions_organization_public_key.test", "key",
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
