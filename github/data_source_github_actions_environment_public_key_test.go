package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubActionsEnvironmentPublicKeyDataSource(t *testing.T) {
	t.Run("queries a repository environment public key without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name = "tf-acc-test-%[1]s"
				auto_init = true
			}

			resource "github_repository_environment" "test" {
				repository = github_repository.test.name
				environment = "tf-acc-test-%[1]s"
			}

			data "github_actions_environment_public_key" "test" {
				repository = github_repository.test.name
				environment = github_repository_environment.test.environment
			}`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"data.github_actions_environment_public_key.test", "key",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnauthenticated(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})
}
