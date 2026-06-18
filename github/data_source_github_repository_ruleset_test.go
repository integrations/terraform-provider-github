package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubRepositoryRulesetDataSource(t *testing.T) {
	t.Run("fetches a repository ruleset", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-ruleset-ds-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_repository_ruleset" "test" {
				name        = "test-ds"
				repository  = github_repository.test.name
				target      = "branch"
				enforcement = "active"

				conditions {
					ref_name {
						include = ["refs/heads/main"]
						exclude = []
					}
				}

				rules {
					creation = true
				}
			}

			data "github_repository_ruleset" "test" {
				repository = github_repository.test.name
				ruleset_id = github_repository_ruleset.test.ruleset_id
			}
		`, repoName)

		const resourceName = "data.github_repository_ruleset.test"

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "name", "test-ds"),
			resource.TestCheckResourceAttr(resourceName, "target", "branch"),
			resource.TestCheckResourceAttr(resourceName, "enforcement", "active"),
			resource.TestCheckResourceAttr(resourceName, "rules.0.creation", "true"),
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
