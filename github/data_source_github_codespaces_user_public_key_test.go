package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubCodespacesUserPublicKeyDataSource(t *testing.T) {
	t.Run("queries an user public key without error", func(t *testing.T) {
		config := `
			data "github_codespaces_user_public_key" "test" {}
		`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"data.github_codespaces_user_public_key.test", "key",
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

		t.Run("with an organization account", func(t *testing.T) {
			t.Skip("organization account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})
	})
}
