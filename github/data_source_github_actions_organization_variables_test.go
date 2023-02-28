package github

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubActionsOrganizationVariablesDataSource(t *testing.T) {

	t.Run("queries actions variables from an organization", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
			resource "github_actions_organization_variable" "test" {
				variable_name 		= "org_variable_%s"
				value = "foo"
				visibility       = "all"
			}
		`, randomID)

		config2 := config + `
			data "github_actions_organization_variables" "test" {
			}
		`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("data.github_actions_organization_variables.test", "variables.#", "1"),
			resource.TestCheckResourceAttr("data.github_actions_organization_variables.test", "variables.0.name", strings.ToUpper(fmt.Sprintf("org_variable_%s", randomID))),
			resource.TestCheckResourceAttr("data.github_actions_organization_variables.test", "variables.0.value", "foo"),
			resource.TestCheckResourceAttr("data.github_actions_organization_variables.test", "variables.0.visibility", "all"),
			resource.TestCheckResourceAttrSet("data.github_actions_organization_variables.test", "variables.0.created_at"),
			resource.TestCheckResourceAttrSet("data.github_actions_organization_variables.test", "variables.0.updated_at"),
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
			testCase(t, organization)
		})
	})
}
