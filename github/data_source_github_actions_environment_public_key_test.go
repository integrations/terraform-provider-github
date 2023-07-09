package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubActionsEnvironmentPublicKeyDataSource(t *testing.T) {

	t.Run("queries an environment public key without error", func(t *testing.T) {

		config := `
		    resource "github_repository" "test" {
		        name      = "tf-acc-test-%s"
		        auto_init = true
		    }

			resource "github_repository_environment" "test" {
				repository       = github_repository.test.name
				environment      = "test_environment_name"
			}

			data "github_actions_environment_public_key" "test" {
				repository       = github_repository.test.name
				environment      = "test_environment_name"
			}
		`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"data.github_actions_environment_public_key.test", "key",
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