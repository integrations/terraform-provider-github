package github

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubActionsOrganizationSecretsDataSource(t *testing.T) {
	t.Run("queries organization actions secrets from a repository", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
			resource "github_actions_organization_secret" "test" {
				secret_name 		= "org_secret_1_%s"
				plaintext_value = "foo"
				visibility      = "all" # going with all as it does not require a paid subscrption
			}
	`, randomID)

		config2 := config + `
			data "github_actions_organization_secrets" "test" {
			}
		`

		check := resource.ComposeTestCheckFunc(
			// resource.TestCheckResourceAttr("data.github_actions_organization_secrets.test", "secrets.#", "1"), // There is no feasible way to know how many secrets exist in the Org during test runs. And I couldn't find a "greater than" operator
			resource.TestCheckTypeSetElemAttr("data.github_actions_organization_secrets.test", "secrets.*.*", strings.ToUpper(fmt.Sprintf("ORG_SECRET_1_%s", randomID))),
			resource.TestCheckTypeSetElemAttr("data.github_actions_organization_secrets.test", "secrets.*.*", "all"),
			resource.TestCheckResourceAttrSet("data.github_actions_organization_secrets.test", "secrets.0.created_at"),
			resource.TestCheckResourceAttrSet("data.github_actions_organization_secrets.test", "secrets.0.updated_at"),
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
