package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubUserEmailAddress(t *testing.T) {

	randomEmailAddress := "example@example.com"

	t.Run("creates and destroys a user email address without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_user_email_address" "test" {
				email_address = "tf-acc-test-%s"
			}
		`, randomEmailAddress)

		check := resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(
				"github_user_email_address.test", "email_address",
				regexp.MustCompile(randomEmailAddress),
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

	})

}
