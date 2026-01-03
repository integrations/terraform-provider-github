package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubActionsVariablesDataSource(t *testing.T) {
	t.Run("queries actions variables from a repository", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-actions-vars-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_actions_variable" "test" {
				variable_name 		= "variable_1"
				repository  		= github_repository.test.name
				value = "foo"
			}
		`, repoName)

		config2 := config + `
			data "github_actions_variables" "test" {
				name = github_repository.test.name
			}
		`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("data.github_actions_variables.test", "name", repoName),
			resource.TestCheckResourceAttr("data.github_actions_variables.test", "variables.#", "1"),
			resource.TestCheckResourceAttr("data.github_actions_variables.test", "variables.0.name", "VARIABLE_1"),
			resource.TestCheckResourceAttr("data.github_actions_variables.test", "variables.0.value", "foo"),
			resource.TestCheckResourceAttrSet("data.github_actions_variables.test", "variables.0.created_at"),
			resource.TestCheckResourceAttrSet("data.github_actions_variables.test", "variables.0.updated_at"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
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
