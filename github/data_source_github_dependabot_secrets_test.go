package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubDependabotSecretsDataSource(t *testing.T) {
	t.Run("queries dependabot secrets from a repository", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-dep-secrets-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_dependabot_secret" "test" {
				secret_name 		= "dep_secret_1"
				repository  		= github_repository.test.name
				plaintext_value = "foo"
			}
		`, repoName)

		config2 := config + `
			data "github_dependabot_secrets" "test" {
				name = github_repository.test.name
			}
		`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("data.github_dependabot_secrets.test", "name", repoName),
			resource.TestCheckResourceAttr("data.github_dependabot_secrets.test", "secrets.#", "1"),
			resource.TestCheckResourceAttr("data.github_dependabot_secrets.test", "secrets.0.name", "DEP_SECRET_1"),
			resource.TestCheckResourceAttrSet("data.github_dependabot_secrets.test", "secrets.0.created_at"),
			resource.TestCheckResourceAttrSet("data.github_dependabot_secrets.test", "secrets.0.updated_at"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  resource.ComposeTestCheckFunc(),
				},
				{
					Config: config2,
					Check:  check,
				},
			},
		})
	})
}
