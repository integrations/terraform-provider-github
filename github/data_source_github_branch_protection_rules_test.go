package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubBranchProtectionRulesDataSource(t *testing.T) {
	t.Parallel()

	t.Run("queries_branch_protection_rules_without_error", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-bp-rules-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "%[1]s"
				auto_init = true
			}

			data "github_branch_protection_rules" "all" {
				repository = github_repository.test.name
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  resource.TestCheckResourceAttr("data.github_branch_protection_rules.all", "rules.#", "0"),
				},
			},
		})
	})

	t.Run("queries_branch_protection", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-bp-rules2-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "%[1]s"
				auto_init = true
			}

			resource "github_branch_protection" "protection" {
				repository_id = github_repository.test.name
			 	pattern = "main*"
			}
		`, repoName)

		config2 := config + `
			data "github_branch_protection_rules" "all" {
				repository = github_repository.test.name
			}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					Config: config2,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("data.github_branch_protection_rules.all", "rules.#", "1"),
						resource.TestCheckResourceAttr("data.github_branch_protection_rules.all", "rules.0.pattern", "main*"),
					),
				},
			},
		})
	})
}
