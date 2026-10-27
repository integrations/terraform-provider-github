package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubRepositoryRulesetsDataSource(t *testing.T) {
	t.Run("reads repository rulesets without error", func(t *testing.T) {
		repoName := fmt.Sprintf("%srepo-rulesets-%s", testResourcePrefix, acctest.RandString(5))
		rulesetName := fmt.Sprintf("test-ruleset-%s", acctest.RandString(5))

		config := fmt.Sprintf(`
		resource "github_repository" "test" {
			name      = "%s"
			auto_init = true
		}

		resource "github_repository_ruleset" "test" {
			name        = "%s"
			repository  = github_repository.test.name
			target      = "branch"
			enforcement = "active"

			conditions {
				ref_name {
					include = ["~ALL"]
					exclude = []
				}
			}

			rules {
				creation = true
			}
		}

		data "github_repository_rulesets" "test" {
			repository = github_repository.test.name

			depends_on = [github_repository_ruleset.test]
		}
		`, repoName, rulesetName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("data.github_repository_rulesets.test", "rulesets.#", "1"),
						resource.TestCheckResourceAttr("data.github_repository_rulesets.test", "rulesets.0.name", rulesetName),
						resource.TestCheckResourceAttr("data.github_repository_rulesets.test", "rulesets.0.target", "branch"),
						resource.TestCheckResourceAttr("data.github_repository_rulesets.test", "rulesets.0.enforcement", "active"),
						resource.TestCheckResourceAttrSet("data.github_repository_rulesets.test", "rulesets.0.ruleset_id"),
						resource.TestCheckResourceAttrSet("data.github_repository_rulesets.test", "rulesets.0.node_id"),
					),
				},
			},
		})
	})
}
