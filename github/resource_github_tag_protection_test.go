package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGitHubTagProtection(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates tag protection without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name         = "test-%[1]s"
			  description  = "Terraform acceptance tests"
			}

			resource "github_tag_protection" "test" {
			  depends_on = ["github_repository.test"]
			  repository = "test-%[1]s"
			  pattern    = "v*"
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_tag_protection.test", "pattern", "v*",
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
		t.Run("with a user account", func(t *testing.T) {
			t.Skip("user account not supported for this operation")
		})
		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, "organization")
		})

	})

}
