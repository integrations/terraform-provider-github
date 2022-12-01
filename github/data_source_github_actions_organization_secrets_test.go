package github

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubActionsOrganizationSecretsDataSource(t *testing.T) {

	t.Run("queries organization actions secrets from a repository", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
			resource "github_actions_organization_secret" "test" {
				secret_name 		= "org_secret_1_%s"
				plaintext_value = "foo"
				visibility      = "private"
			}
	`, randomID)

		config2 := config + `
			data "github_actions_organization_secrets" "test" {
			}
		`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("data.github_actions_organization_secrets.test", "secrets.#", "1"),
			resource.TestCheckResourceAttr("data.github_actions_organization_secrets.test", "secrets.0.name", strings.ToUpper(fmt.Sprintf("ORG_SECRET_1_%s", randomID))),
			resource.TestCheckResourceAttr("data.github_actions_organization_secrets.test", "secrets.0.visibility", "private"),
			resource.TestCheckResourceAttrSet("data.github_actions_organization_secrets.test", "secrets.0.created_at"),
			resource.TestCheckResourceAttrSet("data.github_actions_organization_secrets.test", "secrets.0.updated_at"),
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
