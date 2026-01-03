package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubActionsRegistrationTokenDataSource(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	repoName := fmt.Sprintf("%srepo-actions-regtoken-%s", testResourcePrefix, randomID)

	t.Run("get a repository registration token without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "%[1]s"
				auto_init = true
			}

			data "github_actions_registration_token" "test" {
				repository = github_repository.test.id
			}
		`, repoName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("data.github_actions_registration_token.test", "repository", repoName),
			resource.TestCheckResourceAttrSet("data.github_actions_registration_token.test", "token"),
			resource.TestCheckResourceAttrSet("data.github_actions_registration_token.test", "expires_at"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})
}
