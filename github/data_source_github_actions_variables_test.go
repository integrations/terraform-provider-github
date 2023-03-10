package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubActionsVariablesDataSource(t *testing.T) {

	t.Run("queries actions variables from a repository", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "tf-acc-test-%s"
				auto_init = true
			}

			resource "github_actions_variable" "test" {
				variable_name 		= "variable_1"
				repository  		= github_repository.test.name
				value = "foo"
			}
		`, randomID)

		config2 := config + `
			data "github_actions_variables" "test" {
				name = github_repository.test.name
			}
		`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("data.github_actions_variables.test", "name", fmt.Sprintf("tf-acc-test-%s", randomID)),
			resource.TestCheckResourceAttr("data.github_actions_variables.test", "variables.#", "1"),
			resource.TestCheckResourceAttr("data.github_actions_variables.test", "variables.0.name", "VARIABLE_1"),
			resource.TestCheckResourceAttr("data.github_actions_variables.test", "variables.0.value", "foo"),
			resource.TestCheckResourceAttrSet("data.github_actions_variables.test", "variables.0.created_at"),
			resource.TestCheckResourceAttrSet("data.github_actions_variables.test", "variables.0.updated_at"),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
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
		}

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, "organization")
		})
	})
}
