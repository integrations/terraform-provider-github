package github

import (
	"fmt"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubUserGpgKey(t *testing.T) {
	t.Run("creates a GPG key without error", func(t *testing.T) {
		config := fmt.Sprintf(`
				resource "github_user_gpg_key" "test" {
					armored_public_key = "${file("%s")}"
				}
			`, filepath.Join("test-fixtures", "gpg-pubkey.asc"))

		check := resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(
				"github_user_gpg_key.test",
				"armored_public_key",
				regexp.MustCompile("^-----BEGIN PGP PUBLIC KEY BLOCK-----"),
			),
			resource.TestCheckResourceAttr(
				"github_user_gpg_key.test",
				"key_id",
				"AC541D2D1709CD33",
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
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}
