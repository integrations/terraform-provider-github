package github

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubDependabotOrganizationSecretsDataSource(t *testing.T) {
	t.Parallel()
	skipUnlessHasOrgs(t)

	t.Run("queries_organization_dependabot_secrets_from_a_repository", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
			resource "github_dependabot_organization_secret" "test" {
				secret_name = "org_dep_secret_1_%s"
				value       = "foo"
				visibility  = "private"
			}
		`, randomID)

		config2 := config + `
			data "github_dependabot_organization_secrets" "test" {
			}
		`

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					Config: config2,
					Check: resource.TestCheckTypeSetElemNestedAttrs(
						"data.github_dependabot_organization_secrets.test",
						"secrets.#",
						map[string]string{
							"name":       strings.ToUpper(fmt.Sprintf("ORG_DEP_SECRET_1_%s", randomID)),
							"visibility": "private",
						},
					),
				},
			},
		})
	})
}
